package logic

import (
	"fmt"
	"math/rand"
	"server/pb/pb_enum"
	"time"
)

//PokerBjl 一组牌
type PokerBjl struct {
	index    uint      //当前发牌发到的扑克index
	allPoker []*Poker  //一副扑克
}

//NewPokerBjl 新建一副牌
func NewPokerBjl() *PokerBjl {
	pokerBjl := &PokerBjl{
		index: 0,
	}
	pokerBjl.allPoker = make([]*Poker, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			hua := pb_enum.PokerHua(i+1)
			point := pb_enum.PokerPoint(j+1)
			p := NewPoker(hua, point)
			pokerBjl.allPoker = append(pokerBjl.allPoker, p)
		}
	}
	pokerBjl.Shuffle()
	return pokerBjl
}

//Shuffle 洗牌
func (this *PokerBjl) Shuffle() {
	s := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range this.allPoker {
		j := s.Intn(len(this.allPoker))
		this.allPoker[i], this.allPoker[j] = this.allPoker[j], this.allPoker[i]
	}
}

//FaPai 发牌
func (this *PokerBjl) FaPai() *Poker {
	target := this.allPoker[this.index]
	this.index++
	return target
}

//String 打印一整副牌
func (this *PokerBjl) String() string {
	str := ""
	for i, poker := range this.allPoker {
		str += fmt.Sprintf("[%d]=%s ", i, poker.String())
	}
	return str
}
