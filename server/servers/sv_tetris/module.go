package sv_tetris

import (
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/server"
	"server/pb/pb_tetris"
	"time"
)

func NewServerTetris() module.Module {
	return new(SV_Tetris)
}

type SV_Tetris struct {
	basemodule.BaseModule
	room           *room.Room
	proTime        int64
	gameId         int
	subscribeGroup map [string]gate.Session //订阅该游戏房间信息的session集合
	allTableInfo   map[string]*pb_tetris.TableInfo
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
	this.subscribeGroup =make(map[string]gate.Session,0)
	this.allTableInfo=make(map[string]*pb_tetris.TableInfo,0)
	this.registerRouter()
}

func (this *SV_Tetris) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", this.GetType())
	<-closeSig
	log.Info("%v模块已停止...", this.GetType())
}

func (this *SV_Tetris) OnDestroy() {
	//一定继承
	this.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", this.GetType())
}
