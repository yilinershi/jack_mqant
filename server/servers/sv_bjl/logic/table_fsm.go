package logic

import (
	"fmt"
	"github.com/liangdas/mqant/log"
	"github.com/looplab/fsm"
	"server/pb/pb_bjl"
	"time"
)

//fsmState 状态机的状态
var fsmState = struct {
	none   string
	ready  string
	bet    string
	send   string
	show   string
	settle string
}{
	none:   "none",   //空状态，状态的起点
	ready:  "ready",  //准备阶段，客户端表现一局开始特效，洗牌等动画
	bet:    "bet",    //下注阶段，客户端可在该阶段下注，并发送相应消息到客户端
	send:   "send",   //发牌阶段，客户端表现发牌动画
	show:   "show",   //开牌阶段，客户端表现开牌动画
	settle: "settle", //结算阶段，服务器发送结算结果
}

//fsmEvent 状态机的状态切换事件
var fsmEvent = struct {
	enterAny    string
	enterReady  string
	enterBet    string
	enterSend   string
	enterShow   string
	enterSettle string
}{
	enterAny:    "enter_state", //进入任意阶段的key
	enterReady:  "enter_ready",
	enterBet:    "enter_bet",
	enterSend:   "enter_send",
	enterShow:   "enter_show",
	enterSettle: "enter_settle",
}

//fsmTimeOut 状态机各个状态的超时时间
var fsmTimeOut = struct {
	none   time.Duration
	ready  time.Duration
	bet    time.Duration
	send   time.Duration
	show   time.Duration
	settle time.Duration
}{
	none:   3 * time.Second,
	ready:  3 * time.Second,
	bet:    10 * time.Second,
	send:   3 * time.Second,
	show:   5 * time.Second,
	settle: 5 * time.Second,
}

//OnEnterStatusAny 进入任意阶段,都给房间内的所有session广播状态机状态改变
func (this *Table) OnEnterStatusAny(e *fsm.Event) {
	log.Info(fmt.Sprintf("on enter_any，tableId=%s,当前state=%s\n", this.TableId(), this.tableFSM.Current()))
}

func (this *Table) OnEnterStatusReady(e *fsm.Event) {
	log.Info(fmt.Sprintf("on enter_ready，tableId=[%s],当前state=[%s]，event=%+v\n", this.TableId(), this.tableFSM.Current(), e))
	//每5局重取一副牌，洗一次牌
	isShuffle := false
	if this.curRoundIndex%5 == 0 {
		this.allPoker = NewPokerBjl()
		this.allPoker.Shuffle()
		this.fsmTimer += time.Second * 3
		isShuffle = true
		log.Info("[洗牌], poker=%s", this.allPoker.String())
	}
	this.curRoundIndex++

	str := time.Now().Format("2006-01-02-15-04-05") //这里格式必需是这个格式，据说是go的生日，记忆方法6-1-2-3-4-5
	this.curRoundId = fmt.Sprintf("%s-%s", this.TableId(), str)
	//每局开始时，先清掉当局的下注流水
	this.curRoundBetWaterList = make([]*pb_bjl.BetInfo, 0)

	//每局开始，每个玩家下注信息重置
	for _, p := range this.players {
		p.Reset()
	}

	this.broadcastStateReady(isShuffle)
}

func (this *Table) OnEnterStatusBet(e *fsm.Event) {
	log.Info(fmt.Sprintf("状态机状态切换，tableId=[%s],当前state=[%s]，[下注]\n", this.TableId(), this.tableFSM.Current()))

	this.broadcastStateBet()
}

func (this *Table) OnEnterStatusSend(e *fsm.Event) {
	log.Info(fmt.Sprintf("状态机状态切换，tableId=[%s],当前state=[%s]，[发牌]\n", this.TableId(), this.tableFSM.Current()))

	this.broadcastStateSend()
}

func (this *Table) OnEnterStatusShow(e *fsm.Event) {
	this.curTablePoker = NewTablePoker(this)
	log.Info(fmt.Sprintf("状态机状态切换，tableId=[%s],当前state=[%s]，[开牌] => %s", this.TableId(), this.tableFSM.Current(), this.curTablePoker.String()))

	this.broadcastStateShow()
}

func (this *Table) OnEnterStatusSettle(e *fsm.Event) {
	result := this.curTablePoker.CalResult()
	history := NewHistory(this.curRoundIndex, this.curRoundBetWaterList, this.curTablePoker, result)
	//只记录该桌子上最近20局的记录，如果记录太多，删掉前记录
	if len(this.histories) >= 20 {
		this.histories = this.histories[1:]
	}
	this.histories = append(this.histories, history)
	log.Info(fmt.Sprintf("状态机状态切换，tableId=[%s],当前state=[%s]，[结算]=[%+v]\n", this.TableId(), this.tableFSM.Current(), resultToString(result)))

	for _, info := range this.curRoundBetWaterList {
		if p, isOk := this.players[info.UID]; isOk {
			if result.WinType == pb_bjl.EnumWinType_Xian && info.Area == pb_bjl.EnumBetArea_AreaXian {
				p.winCount += 0.95 * info.Count
				p.gold += p.winCount + info.Count //加钱时，要把下注的钱也加回来
			}
			if result.WinType == pb_bjl.EnumWinType_Zhuang && info.Area == pb_bjl.EnumBetArea_AreaZhuang {
				p.winCount += info.Count
				p.gold += p.winCount + info.Count
			}
			if result.WinType == pb_bjl.EnumWinType_He && info.Area == pb_bjl.EnumBetArea_AreaHe {
				p.winCount += info.Count * 8
				p.gold += p.winCount + info.Count
			}
			if result.IsZhuangDui && info.Area == pb_bjl.EnumBetArea_AreaZhuangDui {
				p.winCount += info.Count * 10
				p.gold += p.winCount + info.Count
			}
			if result.IsXianDui && info.Area == pb_bjl.EnumBetArea_AreaXianDui {
				p.winCount += info.Count * 10
				p.gold += p.winCount + info.Count
			}
		}
	}

	//todo: save player gold to redis

	this.broadcastStateSettle(result)
}
