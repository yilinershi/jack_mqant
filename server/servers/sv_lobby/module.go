/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package sv_lobby

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
)

func NewServerLobby() module.Module {
	return new(SV_Lobby)
}

type SV_Lobby struct {
	basemodule.BaseModule
}

func (this *SV_Lobby) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "SV_Lobby"
}

func (this *SV_Lobby) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

func (this *SV_Lobby) OnAppConfigurationLoaded(app module.App) {
	//当App初始化时调用，这个接口不管这个模块是否在这个进程运行都会调用
	this.BaseModule.OnAppConfigurationLoaded(app)
}

func (this *SV_Lobby) OnInit(app module.App, settings *conf.ModuleSettings) {
	this.BaseModule.OnInit(this, app, settings)
	this.GetServer().Options().Metadata["state"] = "alive"
	log.Info("%v模块初始化完成...", this.GetType())

	this.registerHandle()
}

func (this *SV_Lobby) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", this.GetType())
	<-closeSig
	log.Info("%v模块已停止...", this.GetType())
}

func (this *SV_Lobby) OnDestroy() {
	//一定继承
	this.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", this.GetType())
}
