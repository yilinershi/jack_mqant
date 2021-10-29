package login

import (
	"github.com/gomodule/redigo/redis"
	"github.com/liangdas/mqant/server"
)

type loginComponent struct {
	server                 server.Server
	redisConn              redis.Conn
	RedisKeyAccount        string
	RedisKeyAutoIdAccount  string
	RedisAccountStartIndex int64
	RedisKeyUser           string
}

func NewLogin(server server.Server, conn redis.Conn) *loginComponent {
	l:= &loginComponent{
		server:server,
		redisConn: conn,
		RedisKeyAccount: "account",
		RedisKeyAutoIdAccount: "auto_id_account",
		RedisAccountStartIndex: 10000,
		RedisKeyUser: "user",
	}
	return l
}

func (this *loginComponent) Register()  {
	this.server.Register("rpcIsAccountExist", this.rpcIsAccountExist)
	this.server.Register("rpcCreateAccount", this.rpcCreateAccount)
	this.server.Register("rpcLoadAccount", this.rpcLoadAccount)
	this.server.Register("rpcSaveAccount", this.rpcSaveAccount)
	this.server.Register("rpcLoadUser", this.rpcLoadUser)
	this.server.Register("rpcSaveUser", this.rpcSaveUser)
}
