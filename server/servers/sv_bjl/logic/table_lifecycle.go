package logic

import (
	"fmt"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/module"
	"github.com/looplab/fsm"
	"strconv"
	"time"
)

func (this *Table) GetApp() module.App {
	return this.module.GetApp()
}

func (this *Table) GetSeats() map[string]room.BasePlayer {
	m := map[string]room.BasePlayer{}
	for _, v := range this.players {
		m[fmt.Sprintf("%d", v.UserID)] = v
	}
	return m
}

func (this *Table) OnDestroy() {
	this.QTable.OnDestroy()
}

func (this *Table) OnCreate() {
	//可以加载数据
	fmt.Printf("bjl table OnCreate, table id=%s\n", this.TableId())
	//一定要调用QTable.OnCreate()
	this.QTable.OnCreate()

	//初始化状态机
	this.tableFSM = fsm.NewFSM(
		fsmState.none, //起始状态
		fsm.Events{
			{Name: fsmEvent.enterReady, Src: []string{fsmState.none, fsmState.settle}, Dst: fsmState.ready}, //定义可以由哪些状态切换到下注状态，可以由none或settle切换到ready
			{Name: fsmEvent.enterBet, Src: []string{fsmState.ready}, Dst: fsmState.bet},                     //定义可以由哪些状态切换到下注状态
			{Name: fsmEvent.enterSend, Src: []string{fsmState.bet}, Dst: fsmState.send},
			{Name: fsmEvent.enterShow, Src: []string{fsmState.send}, Dst: fsmState.show},
			{Name: fsmEvent.enterSettle, Src: []string{fsmState.show}, Dst: fsmState.settle},
		},
		fsm.Callbacks{
			fsmEvent.enterAny:    func(e *fsm.Event) { this.OnEnterStatusAny(e) },
			fsmEvent.enterReady:  func(e *fsm.Event) { this.OnEnterStatusReady(e) },
			fsmEvent.enterBet:    func(e *fsm.Event) { this.OnEnterStatusBet(e) },
			fsmEvent.enterSend:   func(e *fsm.Event) { this.OnEnterStatusSend(e) },
			fsmEvent.enterShow:   func(e *fsm.Event) { this.OnEnterStatusShow(e) },
			fsmEvent.enterSettle: func(e *fsm.Event) { this.OnEnterStatusSettle(e) },
		},
	)

	//错锋进入ready状态，错峰的意义在于多个房间不会同时进入同一状态，而在同一时间点增加计算量，空闲时间时又集中空闲
	intId, _ := strconv.Atoi(this.TableId())
	delayTime := time.Second * time.Duration(intId-10000)
	this.fsmTimer = fsmTimeOut.none + delayTime
}

func (this *Table) OnTimeOut() {
	this.Finish()
}

func (this *Table) Update(dt time.Duration) {
	this.fsmTimer -= dt
	switch this.tableFSM.Current() {
	case fsmState.none:
		if this.fsmTimer < 0 {
			this.fsmTimer = fsmTimeOut.ready
			this.tableFSM.Event(fsmEvent.enterReady)
		}
	case fsmState.ready:
		if this.fsmTimer < 0 {
			this.fsmTimer = fsmTimeOut.bet
			this.tableFSM.Event(fsmEvent.enterBet)
		}
	case fsmState.bet:
		if this.fsmTimer < 0 {
			this.fsmTimer = fsmTimeOut.send
			this.tableFSM.Event(fsmEvent.enterSend)
		}
	case fsmState.send:
		if this.fsmTimer < 0 {
			this.fsmTimer = fsmTimeOut.show
			this.tableFSM.Event(fsmEvent.enterShow)
		}
	case fsmState.show:
		if this.fsmTimer < 0 {
			this.fsmTimer = fsmTimeOut.settle
			this.tableFSM.Event(fsmEvent.enterSettle)
		}
	case fsmState.settle:
		if this.fsmTimer < 0 {
			this.fsmTimer = fsmTimeOut.ready
			this.tableFSM.Event(fsmEvent.enterReady)
		}
	}
}
