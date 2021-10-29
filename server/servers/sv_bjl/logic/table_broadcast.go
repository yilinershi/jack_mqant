package logic

import (
	"github.com/liangdas/mqant/log"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_bjl"
	"server/pb/pb_common"
	"time"
)

func (this *Table) broadcastTablePlayerChange(onChangePlayer *Player, changeType pb_bjl.BroadcastTablePlayerChange_EnumChangeType) {
	var pbAllPlayer []*pb_bjl.BjlPlayer
	for _, p := range this.players {
		pbAllPlayer = append(pbAllPlayer, &pb_bjl.BjlPlayer{
			NickName: p.NickName,
			Gold:     p.gold,
		})
	}
	var pbOnChangePlayer = &pb_bjl.BjlPlayer{
		NickName: onChangePlayer.NickName,
		Gold:     onChangePlayer.gold,
	}
	data := &pb_bjl.BroadcastTablePlayerChange{
		ChangeType:     changeType,
		OnChangePlayer: pbOnChangePlayer,
		AllPlayer:      pbAllPlayer,
	}
	bytes, err := proto.Marshal(data)
	if err != nil {
		return
	}
	log.Info("[broadcastTablePlayerChange]  data=%+v\n", data)

	this.NotifyCallBackMsgNR("SV_Bjl/Table/BroadcastTablePlayerChange", bytes)
}

func (this *Table) broadcastStateReady(isShuffle bool) {
	data := &pb_bjl.BroadcastStatusReady{
		GameStatus: this.GetCurState(),
		IsShuffle:  isShuffle,
		RoundId:    this.curRoundId,
		Time:       uint32(this.fsmTimer / time.Second),
	}
	bytes, err := proto.Marshal(data)
	if err != nil {
		return
	}
	log.Info("[broadcastStateReady]  data=%+v\n", data)
	this.NotifyCallBackMsgNR("SV_Bjl/Table/BroadcastStateReady", bytes)
}

func (this *Table) broadcastStateBet() {
	data := &pb_bjl.BroadcastStatusBet{
		GameStatus: this.GetCurState(),
		Time:       uint32(this.fsmTimer / time.Second),
	}
	bytes, err := proto.Marshal(data)
	if err != nil {
		return
	}
	log.Info("[broadcastStateBet]  data=%+v\n", data)
	this.NotifyCallBackMsgNR("SV_Bjl/Table/BroadcastStateBet", bytes)
}

func (this *Table) broadcastStateSend() {
	data := &pb_bjl.BroadcastStatusSend{
		GameStatus: this.GetCurState(),
		Time:       uint32(this.fsmTimer / time.Second),
	}
	bytes, err := proto.Marshal(data)
	if err != nil {
		return
	}
	log.Info("[broadcastStateSend]  data=%+v\n", data)
	this.NotifyCallBackMsgNR("SV_Bjl/Table/BroadcastStateSend", bytes)
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

	data := &pb_bjl.BroadcastStatusShow{
		GameStatus: this.GetCurState(),
		Time:       uint32(this.fsmTimer / time.Second),
		Xian:       xian,
		Zhuang:     zhuang,
	}
	bytes, err := proto.Marshal(data)
	if err != nil {
		return
	}
	log.Info("[broadcastStateReady]  data=%+v\n", data)
	this.NotifyCallBackMsgNR("SV_Bjl/Table/BroadcastStateShow", bytes)
}

func (this *Table) broadcastStateSettle(result *pb_bjl.Result) {
	data := &pb_bjl.BroadcastStatusSettle{
		GameStatus: this.GetCurState(),
		Time:       uint32(this.fsmTimer / time.Second),
		Result:     result,
	}
	for _, p := range this.players {
		data.GoldChange = p.winCount
		data.Gold = p.gold
		log.Info("[broadcastStateSettle]  data=%+v\n", data)
		bytes, err := proto.Marshal(data)
		if err != nil {
			return
		}
		this.SendCallBackMsgNR([]string{p.Session().GetSessionID()}, "SV_Bjl/Table/BroadcastStateSettle", bytes)
	}
}
