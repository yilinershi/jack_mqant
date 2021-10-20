package entity

import (
	"github.com/liangdas/mqant-modules/room"
)

type Player struct {
	room.BasePlayerImp
	SeatIndex  int
	Score      int //金币数量
	timeToMove int64
}

func NewPlayer(SeatIndex int) *Player {
	p := &Player{
		SeatIndex: SeatIndex,
		Score:     1000,
	}
	return p
}
