package logic

import (
	"github.com/liangdas/mqant/log"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_bjl"
	"server/pb/pb_common"
	"time"
)

func (this *Table) broadcastTablePlayerChange(onChangePlayer *Player, changeType pb_bjl.PushTablePlayerChange_EnumChangeType) {
	var pbAllPlayer []*pb_bjl.BjlPlayer
	for _, basePlayer := range this.players {
		p := basePlayer.(*Player)
		pbAllPlayer = append(pbAllPlayer, &pb_bjl.BjlPlayer{
			NickName: p.NickName,
			Score:    p.Score,
		})
	}
	var pbOnChangePlayer = &pb_bjl.BjlPlayer{
		NickName: onChangePlayer.NickName,
		Score:    onChangePlayer.Score,
	}
	pushData := &pb_bjl.PushTablePlayerChange{
		ChangeType:     changeType,
		OnChangePlayer: pbOnChangePlayer,
		AllPlayer:      pbAllPlayer,
	}
	pushDataByte, err := proto.Marshal(pushData)
	if err != nil {
		return
	}
	log.Info("[broadcastTablePlayerChange]  pushData=%+v\n", pushData)
	for _, basePlayer := range this.players {
		basePlayer.Session().Send("SV_Bjl/Push_TablePlayerChange", pushDataByte)
	}
}

func (this *Table) broadcastStateReady(isShuffle bool) {
	broadcastInfo := &pb_bjl.BroadcastStatusReady{
		GameStatus: stateToPb(this.tableFSM.Current()),
		IsShuffle:  isShuffle,
		RoundId:    this.curRoundId,
		Time:       uint32(this.fsmTimer / time.Second),
	}
	b, err := proto.Marshal(broadcastInfo)
	if err != nil {
		return
	}
	log.Info("[broadcastStateReady]  broadcastInfo=%+v\n", broadcastInfo)
	for _, basePlayer := range this.players {
		basePlayer.Session().Send("SV_Bjl/Table/BroadcastStateReady", b)
	}
}

func (this *Table) broadcastStateBet() {
	broadcastInfo := &pb_bjl.BroadcastStatusBet{
		GameStatus: stateToPb(this.tableFSM.Current()),
		Time:       uint32(this.fsmTimer / time.Second),
	}
	b, err := proto.Marshal(broadcastInfo)
	if err != nil {
		return
	}
	log.Info("[broadcastStateBet]  broadcastInfo=%+v\n", broadcastInfo)
	for _, basePlayer := range this.players {
		basePlayer.Session().Send("SV_Bjl/Table/BroadcastStatusBet", b)
	}
}

func (this *Table) broadcastStateSend() {
	broadcastInfo := &pb_bjl.BroadcastStatusSend{
		GameStatus: stateToPb(this.tableFSM.Current()),
		Time:       uint32(this.fsmTimer / time.Second),
	}
	b, err := proto.Marshal(broadcastInfo)
	if err != nil {
		return
	}
	log.Info("[broadcastStateSend]  broadcastInfo=%+v\n", broadcastInfo)
	for _, basePlayer := range this.players {
		basePlayer.Session().Send("SV_Bjl/Table/BroadcastStateSend", b)
	}
}

func (this *Table) broadcastStateShow() {
	xian := make([]*pb_common.Poker, 0)
	xian = append(xian, &this.curTablePoker.xian1.Poker)
	xian = append(xian, &this.curTablePoker.xian2.Poker)
	if this.curTablePoker.xian3 != nil {
		xian = append(xian, &this.curTablePoker.xian3.Poker)
	}
	zhuang := make([]*pb_common.Poker, 0)
	zhuang = append(zhuang, &this.curTablePoker.zhuang1.Poker)
	zhuang = append(zhuang, &this.curTablePoker.zhuang2.Poker)
	if this.curTablePoker.zhuang3 != nil {
		zhuang = append(zhuang, &this.curTablePoker.zhuang3.Poker)
	}

	broadcastInfo := &pb_bjl.BroadcastStatusShow{
		GameStatus: stateToPb(this.tableFSM.Current()),
		Time:       uint32(this.fsmTimer / time.Second),
		Xian:       xian,
	}
	b, err := proto.Marshal(broadcastInfo)
	if err != nil {
		return
	}
	log.Info("[broadcastStateReady]  broadcastInfo=%+v\n", broadcastInfo)
	for _, basePlayer := range this.players {
		basePlayer.Session().Send("SV_Bjl/Table/BroadcastStatusChange", b)
	}
}

func stateToPb(state string) pb_bjl.EnumGameStatus {
	switch state {
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
