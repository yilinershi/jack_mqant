package sv_db

import (
	"github.com/liangdas/mqant-modules/tools"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"server/servers/sv_db/login"
	"server/servers/sv_db/record_bjl"
)

/*
	数据服务器：
	应用场景1：当游戏逻辑数据发生改变时，只能通过该服务器保存，所有的其它服务器只能通过rpc保存数据，这样当多个服务器改变同一数据时，数据不会发生竞争（因为rpc是基于nats消息队列的）
	应用场景2：当客户端订阅了数据信息，当数据发生改变时，由该服务器将数据push给客户端
*/

func NewServerDB() module.Module {
	return new(SV_DB)
}

type SV_DB struct {
	basemodule.BaseModule
}

func (this *SV_DB) GetType() string {
	return "SV_DB"
}

func (this *SV_DB) Version() string {
	return "1.0.0"
}

func (this *SV_DB) OnInit(app module.App, settings *conf.ModuleSettings) {
	this.BaseModule.OnInit(this, app, settings)
	var redisUri = this.GetModuleSettings().Settings["RedisUri"].(string)
	var redisFactory = tools.GetRedisFactory()
	var redisPool = redisFactory.GetPool(redisUri)
	redisConn, err := redisPool.Dial()
	if err != nil {
		return
	}

	//处理login模块的db数据component
	loginComponent := login.NewLogin(this.GetServer(), redisConn)
	loginComponent.Register()

	recordBjlComponent := record_bjl.NewRecordBjl(this.GetServer(), redisConn)
	recordBjlComponent.Register()
}

func (this *SV_DB) Run(closeSig chan bool) {

}

func (this *SV_DB) OnDestroy() {
	if err := this.GetServer().OnDestroy(); err != nil {
		log.Warning("Module server destroy with err: %v", err)
	}
}

