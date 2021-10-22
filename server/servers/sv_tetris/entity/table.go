package entity

import (
	"errors"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/looplab/fsm"
	"reflect"
	"time"
)

type Table struct {
	tableFSM *fsm.FSM //桌子的状态机
	room.QTable
	app     module.App
	players map[string]room.BasePlayer
	name    string
	creator room.BasePlayer //房间创建者
}

func NewTable(app module.App, opts ...room.Option) *Table {
	t := &Table{
		app:     app,
		players: map[string]room.BasePlayer{},
	}
	opts = append(opts, room.TimeOut(60)) //房间内没有消息活动的超时时间
	opts = append(opts, room.Update(t.Update))
	opts = append(opts, room.NoFound(func(msg *room.QueueMsg) (value reflect.Value, e error) {
		//return reflect.ValueOf(this.doSay), nil
		return reflect.Zero(reflect.ValueOf("").Type()), errors.New("no found handler")
	}))
	opts = append(opts, room.SetRecoverHandle(func(msg *room.QueueMsg, err error) {
		log.Error("Recover %v Error: %v", msg.Func, err.Error())
	}))
	opts = append(opts, room.SetErrorHandle(func(msg *room.QueueMsg, err error) {
		log.Error("Error %v Error: %v", msg.Func, err.Error())
	}))
	t.OnInit(t, opts...)
	t.registerRouter()
	return t
}

func (this *Table) registerRouter() {
	//this.Register("table/onPlayerCreateTable", this.onPlayerCreateTable)
}

func (this *Table) GetApp() module.App {
	return this.app
}

func (this *Table) GetSeats() map[string]room.BasePlayer {
	return this.players
}

func (this *Table) OnCreate() {
	//可以加载数据
	log.Info("tetris table OnCreate, table id=%s\n",this.TableId())
	//一定要调用QTable.OnCreate()
	this.QTable.OnCreate()
}

func (this *Table) Update(ds time.Duration) {

}
//
//func (this *Table) onPlayerCreateTable(session gate.Session, msg []byte) error {
//	player := &room.BasePlayerImp{}
//	player.Bind(session)
//	player.OnRequest(session)
//	this.players[session.GetSessionID()] = player
//	req := new(pb_tetris.ReqCreateTetris)
//	if err := proto.Unmarshal(msg, req); err != nil {
//		log.Info("err---------")
//		return err
//	}
//	log.Info("table.onPlayerCreateTable req =%+v\n", req)
//	this.name = req.TableName
//	this.creator = player
//	rpcErr :=room.Router( "SV_Lobby", "rpcOnPlayerCreateTable", req.TableName)
//		if rpcErr!=nil {
//
//		}
//
//
//	return nil
//}


