package logic

import (
	"server/pb/pb_common"
	"server/pb/pb_enum"
)

type Poker struct {
	pb_common.Poker
	Score int //分值
}

// NewPoker 新建卡牌
func NewPoker(hua pb_enum.PokerHua, point pb_enum.PokerPoint) *Poker {
	p := new(Poker)
	p.Hua = hua
	p.Point = point
	p.Score = toBjlPokerScore(point)
	return p
}

//String 转换成字符串
func (this *Poker) String() string {
	strResult := ""
	// 花色
	switch this.Hua {
	case pb_enum.PokerHua_Tao:
		{
			strResult = "♠"
		}
	case pb_enum.PokerHua_Xin:
		{
			strResult = "♥"
		}
	case pb_enum.PokerHua_Mei:
		{
			strResult = "♣"
		}
	case pb_enum.PokerHua_Fang:
		{
			strResult = "♦"
		}
	}

	// 点数
	switch this.Point {
	case pb_enum.PokerPoint_Point3:
		{
			strResult = strResult + "3"
		}
	case pb_enum.PokerPoint_Point4:
		{
			strResult = strResult + "4"
		}
	case pb_enum.PokerPoint_Point5:
		{
			strResult = strResult + "5"
		}
	case pb_enum.PokerPoint_Point6:
		{
			strResult = strResult + "6"
		}
	case pb_enum.PokerPoint_Point7:
		{
			strResult = strResult + "7"
		}
	case pb_enum.PokerPoint_Point8:
		{
			strResult = strResult + "8"
		}
	case pb_enum.PokerPoint_Point9:
		{
			strResult = strResult + "9"
		}
	case pb_enum.PokerPoint_PointT:
		{
			strResult = strResult + "T"
		}
	case pb_enum.PokerPoint_PointJ:
		{
			strResult = strResult + "J"
		}
	case pb_enum.PokerPoint_PointQ:
		{
			strResult = strResult + "Q"
		}
	case pb_enum.PokerPoint_PointK:
		{
			strResult = strResult + "K"
		}
	case pb_enum.PokerPoint_PointA:
		{
			strResult = strResult + "A"
		}
	case pb_enum.PokerPoint_Point2:
		{
			strResult = strResult + "2"
		}

	}
	return strResult
}

func toBjlPokerScore(pokerPoint pb_enum.PokerPoint) int {
	switch pokerPoint {
	case pb_enum.PokerPoint_Point3:
		{
			return 3
		}
	case pb_enum.PokerPoint_Point4:
		{
			return 4
		}
	case pb_enum.PokerPoint_Point5:
		{
			return 5
		}
	case pb_enum.PokerPoint_Point6:
		{
			return 6
		}
	case pb_enum.PokerPoint_Point7:
		{
			return 7
		}
	case pb_enum.PokerPoint_Point8:
		{
			return 8
		}
	case pb_enum.PokerPoint_Point9:
		{
			return 9
		}
	case pb_enum.PokerPoint_PointT:
		{
			return 0
		}
	case pb_enum.PokerPoint_PointJ:
		{
			return 0
		}
	case pb_enum.PokerPoint_PointQ:
		{
			return 0
		}
	case pb_enum.PokerPoint_PointK:
		{
			return 0
		}
	case pb_enum.PokerPoint_PointA:
		{
			return 1
		}
	case pb_enum.PokerPoint_Point2:
		{
			return 2
		}
	default:
		return 0
	}

}

//
//// 从牌值获取花色
//func toFlower(id int) int {
//	if id <= 0 || id > 54 {
//		return flowerNIL
//	}
//	return ((id - 1) / 13) + 1
//}
////
////// 从牌值获取点数
////func toPointAndScore(id int) (point int, score int) {
////	if id <= 0 {
////		return cardPointNIL, 0
////	}
////	if id == 53 {
////		return cardPointX, 0 // 小王
////	}
////	if id == 54 {
////		return cardPointY, 0 // 大王
////	}
////	point = (id-1)%13 + 1
////	if point == cardPointT || point == cardPointJ || point == cardPointQ || point == cardPointK {
////		score = 0
////	} else {
////		score = point
////	}
////	return point, score
////}
