package sv_db

import (
	"github.com/gomodule/redigo/redis"
	"github.com/liangdas/mqant-modules/tools"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
)

func NewServerDB() module.Module {
	return new(ServerDB)
}

type ServerDB struct {
	basemodule.BaseModule
	redisFactory *tools.RedisFactory
	redisConn    redis.Conn
}

func (this *ServerDB) GetType() string {
	return "SV_DB"
}

func (this *ServerDB) Version() string {
	return "1.0.0"
}

func (this *ServerDB) OnInit(app module.App, settings *conf.ModuleSettings) {
	this.BaseModule.OnInit(this, app, settings)

	var redisUri = this.GetModuleSettings().Settings["RedisUri"].(string)
	this.redisFactory = tools.GetRedisFactory()
	var redisPool = this.redisFactory.GetPool(redisUri)
	redisConn, err := redisPool.Dial()
	if err != nil {
		return
	}
	this.redisConn = redisConn

	this.GetServer().Register("rpcIsAccountExist", this.rpcIsAccountExist)
	this.GetServer().Register("rpcCreateAccount", this.rpcCreateAccount)
	this.GetServer().Register("rpcLoadAccount", this.rpcLoadAccount)
	this.GetServer().Register("rpcSaveAccount", this.rpcSaveAccount)
	this.GetServer().Register("rpcLoadUser", this.rpcLoadUser)
	this.GetServer().Register("rpcSaveUser", this.rpcSaveUser)
}

func (this *ServerDB) Run(closeSig chan bool) {

}

func (this *ServerDB) OnDestroy() {
	if err := this.GetServer().OnDestroy(); err != nil {
		log.Warning("Module server destroy with err: %v", err)
	}
}
