package logic

import (
	"errors"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/looplab/fsm"
	"reflect"
	"server/pb/pb_bjl"
	"time"
)

type Table struct {
	room.QTable                                //基类
	tableFSM        *fsm.FSM                   //桌子的状态机
	module          module.RPCModule           //所在module
	players         map[string]room.BasePlayer //桌子上的玩家
	fsmTimer        time.Duration              //状态机计时器
	curRoundIndex   uint32                     //当前桌子的第几轮
	curRoundId      string						//当前桌子的Id
	curRoundBetInfo []*pb_bjl.BetInfo          //当前桌子上当前局的下注流水
	minBet          uint32                     //最小下注额度
	maxBet          uint32                     //最大下注额度
	histories       []*deskHistory             //桌子上历史结果，只记录近20局的记录
	allPoker        *PokerBjl                  //所有的扑克牌
	curTablePoker   *tablePoker                //当前桌子上发的牌
}

func NewTable(module module.RPCModule, opts ...room.Option) (*Table, error) {
	t := &Table{
		module:     module,
		players:    map[string]room.BasePlayer{},
		curRoundIndex:0,
	}
	opts = append(opts, room.TimeOut(5*60)) //房间内没有消息活动的超时时间
	opts = append(opts, room.Update(t.Update))
	opts = append(opts, room.RunInterval(500*time.Millisecond)) // 时间周期设置为30毫秒
	opts = append(opts, room.NoFound(func(msg *room.QueueMsg) (value reflect.Value, e error) {
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
	return t, nil
}

func (this *Table) registerRouter() {
	this.Register("Table/CallPlayerJoin", this.onPlayerJoin)
}

func (this *Table) onPlayerJoin(session gate.Session, topic string) {
	nickName := session.Get("nickName")
	player := NewPlayer(nickName)
	player.Bind(session)
	player.OnRequest(session)
	this.players[session.GetSessionID()] = player
	//向桌子上广播玩家进入
	this.broadcastTablePlayerChange(player, pb_bjl.PushTablePlayerChange_Join)
}
