package serverHttp

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/liangdas/mqant/log"
	mqrpc "github.com/liangdas/mqant/rpc"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"server/pb/pb_enum"
	"server/pb/pb_http"
	"server/pb/pb_rpc"
	"strings"
	"time"
)

var key = "天王盖地虎,宝塔镇河妖"

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (this *ServerHttp) entry(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	req := new(pb_http.ReqEntry)
	if err := proto.Unmarshal(buf, req); err != nil {
		return
	}
	log.Info("[entry], req.Secret=%s\n", req.Secret)
	//如果客户端的包里不带密钥或是密钥错误，将无法获取真实的游戏服务器地址
	resp := new(pb_http.RespEntry)
	if strings.Contains(req.Secret, "天王盖地虎") && strings.Contains(req.Secret, "宝塔镇河妖") {
		resp.ErrCode = pb_enum.ErrorCode_OK
		resp.LoginUrl = this.loginUrl
		resp.RegisterUrl = this.registerUrl
		resp.WebSocketUrl = this.websocketUrl
		resp.TcpUrl = this.tcpUrl
	} else {
		resp.ErrCode = pb_enum.ErrorCode_EntryError
	}
	bytes, err := proto.Marshal(resp)
	log.Info("[entry] result=%v\n", resp)
	if err != nil {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
}

func (this *ServerHttp) register(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	req := new(pb_http.ReqRegister)
	if err := proto.Unmarshal(buf, req); err != nil {
		return
	}
	log.Info("onRegister account=%s, password=%s\n", req.Account, req.Password)
	resp := new(pb_http.RespRegister)

	//rpc从db中加载account
	exists, rpcErr := mqrpc.Bool(this.Call(context.Background(), "SV_DB", "rpcIsAccountExist", mqrpc.Param(req.Account)))
	log.Info("RpcCall exists=%v , err %v\n",exists,rpcErr)

	if exists==true {
		//账号已存在
		resp.ErrCode = pb_enum.ErrorCode_RegisterAccountExit
	} else {
		isOk, rpcErr2 :=  mqrpc.Bool(this.Call(context.Background(),"SV_DB", "rpcCreateAccount", mqrpc.Param( req.Account, req.Password)))
		if rpcErr2 != nil {
			return
		}
		if isOk==true {
			resp.ErrCode = pb_enum.ErrorCode_OK
		}
	}

	bytes, err3 := proto.Marshal(resp)
	log.Info("[entry] result=%v\n", resp)
	if err3 != nil {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
}

//login 用户通过http登录
func (this *ServerHttp) login(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	req := new(pb_http.ReqLogin)
	if err := proto.Unmarshal(buf, req); err != nil {
		return
	}
	log.Info("on login account=%s, password=%s\n", req.Account, req.Password)
	resp := new(pb_http.RespLogin)

	isExists, _ :=mqrpc.Bool(this.Call(context.Background(),"SV_DB", "rpcIsAccountExist", mqrpc.Param(req.Account)))
	if isExists == false {
		//账号不存在
		resp.ErrCode = pb_enum.ErrorCode_LoginAccountUnExixtent
	} else {
		//rpc从db中加载account
		ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
		a := new(pb_rpc.DbAccount)
		err := mqrpc.Proto(a, func() (reply interface{}, errStr interface{}) {
			return this.Call(ctx, "SV_DB", "rpcLoadAccount", mqrpc.Param(req.Account), )
		})
		log.Info("RpcCall ,account=%+v ,err= %v", a, err)

		if a.Password == req.Password {
			resp.ErrCode = pb_enum.ErrorCode_OK
			a.LastLoginTime = time.Now().Unix()
			a.Token = genToken(req.Account, req.Password)
			isOk, err3 := mqrpc.Bool(this.Call(context.Background(),"SV_DB", "rpcSaveAccount", mqrpc.Param( req.Account, a)))
			if err3 != nil {
				log.Error("save account err=%+v\n",err3.Error())
				return
			}
			if isOk==true {
				resp.Token = a.Token
			}
		} else {
			resp.ErrCode = pb_enum.ErrorCode_LoginPasswordError
		}
	}

	bytes, err3 := proto.Marshal(resp)
	log.Info("[entry] result=%v\n", resp)
	if err3 != nil {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
}

func genToken(account string, password string) string {
	h := md5.New()
	key := fmt.Sprintf("account=%s,password=%s,time=%d", account, password, time.Now().Nanosecond())
	log.Info("token key=", key)
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}
