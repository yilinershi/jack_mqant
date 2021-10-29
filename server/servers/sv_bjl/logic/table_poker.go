package logic

import (
	"fmt"
	"server/pb/pb_bjl"
)

type tablePoker struct {
	table *Table

	xian1 *Poker
	xian2 *Poker
	xian3 *Poker

	zhuang1 *Poker
	zhuang2 *Poker
	zhuang3 *Poker
}

func NewTablePoker(table *Table) *tablePoker {
	target := &tablePoker{
		table: table,

		xian1: table.allPoker.FaPai(),
		xian2: table.allPoker.FaPai(),

		zhuang1: table.allPoker.FaPai(),
		zhuang2: table.allPoker.FaPai(),
	}

	isXianBuPai, xian3 := target.IsXianBuPai()
	if isXianBuPai {
		target.xian3 = xian3
	}

	isZhuangBuPai, zhuang3 := target.IsZhuangBuPai(isXianBuPai)
	if isZhuangBuPai {
		target.zhuang3 = zhuang3
	}
	return target
}

func (this *tablePoker) TotalXianScore() int {
	totalScore := 0
	totalScore += this.xian1.Score + this.xian2.Score
	if this.xian3 != nil {
		totalScore += this.xian3.Score
	}
	totalScore = totalScore % 10
	return totalScore
}

func (this *tablePoker) TotalZhuangScore() int {
	totalScore := 0
	totalScore += this.zhuang1.Score + this.zhuang2.Score
	if this.zhuang3 != nil {
		totalScore += this.zhuang3.Score
	}
	totalScore = totalScore % 10
	return totalScore
}

//IsXianBuPai 闲是否补牌
func (this *tablePoker) IsXianBuPai() (bool, *Poker) {
	if this.TotalXianScore() <= 7 {
		return true, this.table.allPoker.FaPai()
	}
	return false, nil
}

//IsZhuangBuPai 庄是否补牌
func (this *tablePoker) IsZhuangBuPai(isXianBuPai bool) (bool, *Poker) {
	zhuangScore := 0
	zhuangScore += this.zhuang1.Score + this.zhuang2.Score
	zhuangScore = zhuangScore % 10

	switch zhuangScore {
	case 0, 1, 2:
		return true, this.table.allPoker.FaPai()
	case 3:
		if isXianBuPai && this.xian3.Score == 8 {
			return false, nil
		}
		return true, this.table.allPoker.FaPai()
	case 4:
		if isXianBuPai && (this.xian3.Score == 0 || this.xian3.Score == 1 || this.xian3.Score == 8 || this.xian3.Score == 9) {
			return false, nil
		}
		return true, this.table.allPoker.FaPai()
	case 5:
		if isXianBuPai && (this.xian3.Score == 0 || this.xian3.Score == 1 || this.xian3.Score == 2 || this.xian3.Score == 3 || this.xian3.Score == 8 || this.xian3.Score == 9) {
			return false, nil
		}
		return true, this.table.allPoker.FaPai()
	case 6:
		if isXianBuPai && (this.xian3.Score == 6 || this.xian3.Score == 7) {
			return false, nil
		}
		return true, this.table.allPoker.FaPai()
	default:
		return false, nil
	}
}

func (this *tablePoker) String() string {
	str := ""
	if this.xian3 == nil {
		str += fmt.Sprintf("闲=[%s,%s]", this.xian1, this.xian2)
	} else {
		str += fmt.Sprintf(",闲=[%s,%s,%s]", this.xian1, this.xian2, this.xian3)
	}

	if this.zhuang3 == nil {
		str += fmt.Sprintf(" ,庄=[%s,%s]", this.zhuang1, this.zhuang2)
	} else {
		str += fmt.Sprintf(",庄=[%s,%s,%s]", this.zhuang1, this.zhuang2, this.zhuang3)
	}
	return str
}

func (this *tablePoker) CalResult() *pb_bjl.Result {
	result := new(pb_bjl.Result)
	if this.TotalXianScore() > this.TotalZhuangScore() {
		result.WinType = pb_bjl.EnumWinType_Xian
	} else if this.TotalZhuangScore() > this.TotalXianScore() {
		result.WinType = pb_bjl.EnumWinType_Zhuang
	} else {
		result.WinType = pb_bjl.EnumWinType_He
	}

	if this.xian1.Point == this.xian2.Point {
		result.IsXianDui = true
	} else {
		result.IsXianDui = false
	}

	if this.zhuang1.Point == this.zhuang2.Point {
		result.IsZhuangDui = true
	} else {
		result.IsZhuangDui = false
	}

	return result
}
