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
	app  module.App
	players map[string]room.BasePlayer
}

func NewTable(app module.App, opts ...room.Option) *Table {
	t := &Table{
		app:  app,
		players: map[string]room.BasePlayer{},
	}
	opts = append(opts, room.TimeOut(60))
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
	//t.Register("/room/say", this.doSay)
	//t.Register("/room/join", this.doJoin)
	return t
}

func (this *Table) GetApp() module.App {
	return this.app
}


func (this *Table) OnCreate() {
	//可以加载数据
	log.Info("MyTable OnCreate")
	//一定要调用QTable.OnCreate()
	this.QTable.OnCreate()
}


func (this *Table) Update(ds time.Duration) {

}
