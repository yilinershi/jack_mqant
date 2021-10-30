package logic

import (
	"errors"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/looplab/fsm"
	"google.golang.org/protobuf/proto"
	"reflect"
	"server/pb/pb_bjl"
	"server/pb/pb_common"
	"server/pb/pb_enum"
	"time"
)

type Table struct {
	room.QTable                            //基类
	tableFSM             *fsm.FSM          //桌子的状态机
	module               module.RPCModule  //所在module
	players              map[int64]*Player //桌子上的玩家
	fsmTimer             time.Duration     //状态机计时器
	curRoundIndex        uint32            //当前桌子的第几轮
	curRoundId           string            //当前桌子的Id
	curRoundBetWaterList []*pb_bjl.BetInfo //当前桌子上当前局的下注流水
	minBet               uint32            //最小下注额度
	maxBet               uint32            //最大下注额度
	histories            []*deskHistory    //桌子上历史结果，只记录近20局的记录
	allPoker             *PokerBjl         //所有的扑克牌
	curTablePoker        *tablePoker       //当前桌子上发的牌
}

func NewTable(module module.RPCModule, opts ...room.Option) (*Table, error) {
	t := &Table{
		module:        module,
		players:       make(map[int64]*Player, 0),
		curRoundIndex: 0,
	}
	opts = append(opts, room.TimeOut(100*365*24*60*60))         //房间内没有消息活动的超时时间，这里100年不自动关闭
	opts = append(opts, room.Update(t.Update))                  //设定update函数
	opts = append(opts, room.RunInterval(500*time.Millisecond)) // 时间周期设置为500毫秒
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
	this.Register("Table/CallBet", this.onPlayerBet)
	this.Register("Table/CallHeartbeat", this.onPlayerHeartbeat)
}

func (this *Table) onPlayerHeartbeat(session gate.Session, topic string,req *pb_common.ReqHeartbeat) {
	basePlayer := this.FindPlayer(session)
	basePlayer.OnRequest(session)  //记录桌子上的玩家在发送消息
	resp:=&pb_common.RespHeartbeat{
		Pong: "pong",
	}
	bytes, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	log.Info("[onPlayerHeartbeat]  data=%+v\n", resp)
	session.Send(topic,bytes)
}

func (this *Table) onPlayerJoin(session gate.Session, topic string) {
	player := NewPlayer()
	player.Bind(session)
	player.OnRequest(session)
	player.UserID = session.GetUserIdInt64()
	this.players[player.UserID] = player
	//向桌子上广播玩家进入
	this.broadcastTablePlayerChange(player, pb_bjl.BroadcastTablePlayerChange_Join)
}

func (this *Table) onPlayerBet(session gate.Session, topic string, req *pb_bjl.ReqBet) {
	basePlayer := this.FindPlayer(session)
	basePlayer.OnRequest(session)  //记录桌子上的玩家在发送消息
	resp := new(pb_bjl.RespBet)
	if this.tableFSM.Current() != fsmState.bet {
		resp.ErrCode = pb_enum.ErrorCode_TableStateCanNotBet
	} else {
		p := basePlayer.(*Player)
		if p.gold < req.Count {
			resp.ErrCode = pb_enum.ErrorCode_TableBetMoneyUnEnough
		} else {
			p.gold = p.gold - req.Count
			resp.ErrCode = pb_enum.ErrorCode_OK
			resp.Gold = p.gold
		}
	}

	bytes, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	log.Info("[onPlayerBet]  resp=%+v\n", resp)
	session.Send(topic, bytes)

	if resp.ErrCode == pb_enum.ErrorCode_OK {
		uid := session.GetUserIdInt64()
		betInfo := &pb_bjl.BetInfo{
			UID:   uid,
			Count: req.Count,
			Area:  req.Area,
		}
		this.curRoundBetWaterList = append(this.curRoundBetWaterList, betInfo)
		this.broadcastPlayerBet(betInfo)
	}
}

//GetCurState 获取当前状态机的状态(protobuf格式)
func (this *Table) GetCurState() pb_bjl.EnumGameStatus {
	switch this.tableFSM.Current() {
	case fsmState.none:
		return pb_bjl.EnumGameStatus_None
	case fsmState.ready:
		return pb_bjl.EnumGameStatus_Ready
	case fsmState.bet:
		return pb_bjl.EnumGameStatus_Bet
	case fsmState.send:
		return pb_bjl.EnumGameStatus_Send
	case fsmState.show:
		return pb_bjl.EnumGameStatus_Show
	case fsmState.settle:
		return pb_bjl.EnumGameStatus_Settle
	default:
		return pb_bjl.EnumGameStatus_None
	}
}
