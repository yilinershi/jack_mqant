package sv_bjl

import (
	"fmt"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/server"
	"server/servers/sv_bjl/logic"
	"strconv"
	"time"
)

func NewServerBjl() module.Module {
	return new(SV_Bjl)
}

type SV_Bjl struct {
	basemodule.BaseModule
	room   *room.Room
	tables map[string]*logic.Table
}

func (this *SV_Bjl) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "SV_Bjl"
}

func (this *SV_Bjl) Version() string {
	return "1.0.0"
}

func (this *SV_Bjl) OnInit(app module.App, settings *conf.ModuleSettings) {
	this.BaseModule.OnInit(this, app, settings,
		server.RegisterInterval(15*time.Second),
		server.RegisterTTL(30*time.Second),
	)
	this.room = room.NewRoom(this.GetApp())
	this.tables = make(map[string]*logic.Table)
	const tableCount = 1 //房间数量
	for i := 0; i < tableCount; i++ {
		tableId := strconv.Itoa(10001 + i)
		table, err := this.room.CreateById(this.GetApp(), tableId, this.NewTable)
		if err != nil {
			return
		}
		table.Run()
		log.Info("创建table tableId=%s", tableId)
	}
	this.registerRouter()
}

func (this *SV_Bjl) NewTable(module module.App, tableId string) (room.BaseTable, error) {
	table, err := logic.NewTable(
		this,
		room.TableId(tableId),
		room.Router(func(TableId string) string {
			return fmt.Sprintf("%v://%v/room", this.GetType(), this.GetServer().ID())
		}),
		room.DestroyCallbacks(func(table room.BaseTable) error {
			log.Info("回收了房间: %v", table.TableId())
			_ = this.room.DestroyTable(table.TableId())
			return nil
		}),
	)
	return table, err
}

func (this *SV_Bjl) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", this.GetType())
	<-closeSig
	log.Info("%v模块已停止...", this.GetType())
}

func (this *SV_Bjl) OnDestroy() {
	//一定继承
	this.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", this.GetType())
}
