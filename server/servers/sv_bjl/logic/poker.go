package logic

import (
	"server/pb/pb_common"
	"server/pb/pb_enum"
)

type Poker struct {
	pb_common.Poker     //继承protobuf
	Score           int //分值,每个游戏的分值不同
}

// NewPoker 新建一张牌
func NewPoker(hua pb_enum.PokerHua, point pb_enum.PokerPoint) *Poker {
	p := new(Poker)
	p.Hua = hua
	p.Point = point
	p.Score = toBjlPokerScore(point)
	return p
}

//String 转换成字符串
func (this *Poker) String() string {
	str := ""
	// 花色
	switch this.Hua {
	case pb_enum.PokerHua_Tao:
		{
			str = "♠"
		}
	case pb_enum.PokerHua_Xin:
		{
			str = "♥"
		}
	case pb_enum.PokerHua_Mei:
		{
			str = "♣"
		}
	case pb_enum.PokerHua_Fang:
		{
			str = "♦"
		}
	}

	// 点数
	switch this.Point {
	case pb_enum.PokerPoint_Point3:
		{
			str = str + "3"
		}
	case pb_enum.PokerPoint_Point4:
		{
			str = str + "4"
		}
	case pb_enum.PokerPoint_Point5:
		{
			str = str + "5"
		}
	case pb_enum.PokerPoint_Point6:
		{
			str = str + "6"
		}
	case pb_enum.PokerPoint_Point7:
		{
			str = str + "7"
		}
	case pb_enum.PokerPoint_Point8:
		{
			str = str + "8"
		}
	case pb_enum.PokerPoint_Point9:
		{
			str = str + "9"
		}
	case pb_enum.PokerPoint_PointT:
		{
			str = str + "T"
		}
	case pb_enum.PokerPoint_PointJ:
		{
			str = str + "J"
		}
	case pb_enum.PokerPoint_PointQ:
		{
			str = str + "Q"
		}
	case pb_enum.PokerPoint_PointK:
		{
			str = str + "K"
		}
	case pb_enum.PokerPoint_PointA:
		{
			str = str + "A"
		}
	case pb_enum.PokerPoint_Point2:
		{
			str = str + "2"
		}
	}
	return str
}

func toBjlPokerScore(pokerPoint pb_enum.PokerPoint) int {
	switch pokerPoint {
	case pb_enum.PokerPoint_Point3:
		return 3
	case pb_enum.PokerPoint_Point4:
		return 4
	case pb_enum.PokerPoint_Point5:
		return 5
	case pb_enum.PokerPoint_Point6:
		return 6
	case pb_enum.PokerPoint_Point7:
		return 7
	case pb_enum.PokerPoint_Point8:
		return 8
	case pb_enum.PokerPoint_Point9:
		return 9
	case pb_enum.PokerPoint_PointT:
		return 0
	case pb_enum.PokerPoint_PointJ:
		return 0
	case pb_enum.PokerPoint_PointQ:
		return 0
	case pb_enum.PokerPoint_PointK:
		return 0
	case pb_enum.PokerPoint_PointA:
		return 1
	case pb_enum.PokerPoint_Point2:
		return 2
	default:
		return 0
	}
}
