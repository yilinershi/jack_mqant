package logic

import (
	"github.com/liangdas/mqant-modules/room"
)

type Player struct {
	room.BasePlayerImp
	Score    uint32 //分数
	NickName string
}

func NewPlayer(nickName string) *Player {
	p := &Player{
		Score:    1000,
		NickName: nickName,
	}
	return p
}
