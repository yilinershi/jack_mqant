package logic

import (
	"github.com/liangdas/mqant-modules/room"
)

type Player struct {
	room.BasePlayerImp
	UserID int64
	NickName string
	winCount float32
	gold     float32
}

func NewPlayer() *Player {
	p := &Player{
		gold:     1000,
		winCount: 0,
	}
	return p
}

func (this *Player) Reset() {
	this.winCount = 0
}

func (this *Player) GetUserID() int64 {
	return this.UserID
}
