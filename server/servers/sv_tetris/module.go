package sv_tetris

import (
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/server"
	"time"
)

type SV_Tetris struct {
	basemodule.BaseModule
	room    *room.Room
	proTime int64
	gameId  int
}

func (this *SV_Tetris) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "SV_Tetris"
}

func (this *SV_Tetris) Version() string {
	return "1.0.0"
}

func (this *SV_Tetris) OnInit(app module.App, settings *conf.ModuleSettings) {
	this.BaseModule.OnInit(this, app, settings,
		server.RegisterInterval(15*time.Second),
		server.RegisterTTL(30*time.Second),
	)
	this.room = room.NewRoom(this.GetApp())
	this.registerHandle()
}

func (this *SV_Tetris) Run(closeSig chan bool) {

}

func (this *SV_Tetris) OnDestroy() {
	//一定别忘了关闭RPC
	this.GetServer().OnDestroy()
}
