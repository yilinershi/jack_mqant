package serverDB

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"server/pb/pb_enum"
	"server/pb/pb_rpc"
	"strconv"
	"time"
)

const (
	RedisKeyAccount        = "account"
	RedisKeyAutoIdAccount  = "auto_id_account"
	RedisAccountStartIndex = 10000
	RedisKeyUser           = "user"
)

/**
1.普通格式的rpc，返回系统变量及string,即 return (bool,string)
2.protobuf格式的rpc,返回自定义的protobuf及自定义消息，即 return (* pb_xxx, error), 其中pb_xxx为自定义的protobuf结构
*/

//rpcIsAccountExist 通过rpc查询redis中账号是否已存在,
func (this *ServerDB) rpcIsAccountExist(account string) (bool, string) {
	isExists, err := redis.Bool(this.redisConn.Do("HEXISTS", RedisKeyAccount, account))
	if err != nil {
		return false, err.Error()
	}
	return isExists, ""
}

//rpcCreatePlayerData 通过rpc创建账号
func (this *ServerDB) rpcCreateAccount(account string, password string) (bool, string) {
	//step1:创建玩家账户
	id, err := redis.Int(this.redisConn.Do("INCR", RedisKeyAutoIdAccount))
	if err != nil {
		return false, "create user uid fail"
	}
	a := &pb_rpc.DbAccount{}
	a.UID = RedisAccountStartIndex + int64(id)
	a.Account = account
	a.Password = password
	a.CreateTime = time.Now().Unix()
	if _, err2 := this.saveAccount(account, a); err2 != nil {
		return false, err2.Error()
	}

	//step2:创建账户的同时，创建玩家数据
	u := &pb_rpc.DbUser{
		UID:      a.UID,
		NickName: "玩家" + strconv.Itoa(int(a.UID)),
		Sex:      pb_enum.Sex_Unknow,
		Gold:     1000,
		Diamond:  0,
		Icon:     "default_icon",
	}
	if _, err2 := this.saveUser(a.UID, u); err2 != nil {
		return false, err2.Error()
	}
	return true, ""
}

//rpcLoadAccount 通过rpc加载玩家账户数据
func (this *ServerDB) rpcLoadAccount(account string) (*pb_rpc.DbAccount, error) {
	res, err := redis.String(this.redisConn.Do("HGET", RedisKeyAccount, account))
	if err != nil {
		return nil, err
	}
	var a = &pb_rpc.DbAccount{}
	if jsonErr := json.Unmarshal([]byte(res), a); jsonErr != nil {
		return nil, jsonErr
	}
	return a, nil
}

//rpcSaveAccount 保存玩家账户数据到数据库
func (this *ServerDB) rpcSaveAccount(account string, a *pb_rpc.DbAccount) (bool, string) {
	_, err := this.saveAccount(account, a)
	if err != nil {
		return false, "save account error"
	}
	return true, ""
}

//rpcLoadUser 加载用户数据
func (this *ServerDB) rpcLoadUser(uid int64) (*pb_rpc.DbUser, error) {
	res, err := redis.String(this.redisConn.Do("HGET", RedisKeyUser, uid))
	if err != nil {
		return nil, err
	}
	var a = &pb_rpc.DbUser{}
	if jsonErr := json.Unmarshal([]byte(res), a); jsonErr != nil {
		return nil, jsonErr
	}
	return a, nil
}

//rpcSaveUser 保存玩家数据到数据库
func (this *ServerDB) rpcSaveUser(uid int64, a *pb_rpc.DbUser) (string, error) {
	isOk, err := this.saveUser(uid, a)
	if err != nil {
		return "false", err
	}
	if isOk {

		return "true", nil
	}
	return "false", nil
}

func (this *ServerDB) saveAccount(account string, a *pb_rpc.DbAccount) (bool, error) {
	buf, err1 := json.Marshal(a)
	if err1 != nil {
		return false, err1
	}
	buffer := string(buf)
	if _, err2 := this.redisConn.Do("HSET", RedisKeyAccount, account, buffer); err2 != nil {
		return false, err2
	}
	return true, nil
}

//SaveUser 存储用户数据
func (this *ServerDB) saveUser(uid int64, u *pb_rpc.DbUser) (bool, error) {
	buf, err1 := json.Marshal(u)
	if err1 != nil {
		return false, err1
	}
	buffer := string(buf)
	if _, err2 := this.redisConn.Do("HSET", RedisKeyUser, uid, buffer); err2 != nil {
		return false, err2
	}
	return true, nil
}
