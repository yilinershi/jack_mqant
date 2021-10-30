using System.Collections.Generic;
using Pb.Bjl;
using Pb.Common;

public class BjlTableModel
{
    public  List<BjlPlayerData> AllPlayer;
    public  List<Poker> Xian;
    public  List<Poker> Zhaung;
    public  Result RoundResult;

    public  BjlPlayerData FindPlayerByUid(long uid)
    {
        foreach (var p in AllPlayer)
        {
            if (p.BaseInfo.UID == uid)
            {
                return p;
            }
        }

        return null;
    }

    public class BjlPlayerData
    {
        public BjlPlayer BaseInfo;
        public List<BetInfo> BetWaterList = new List<BetInfo>(); //玩家本局的下注流水

        public float WinCount=0;
        
        public float XianTotal
        {
            get
            {
                float total = 0;
                foreach (var item in BetWaterList)
                {
                    if (item.Area == Pb.Bjl.EnumBetArea.AreaXian)
                    {
                        total += item.Count;
                    }
                }

                return total;
            }
        }
        
        public float ZhuangTotal
        {
            get
            {
                float total = 0;
                foreach (var item in BetWaterList)
                {
                    if (item.Area == Pb.Bjl.EnumBetArea.AreaZhuang)
                    {
                        total += item.Count;
                    }
                }

                return total;
            }
        }
        
        public float HeTotal
        {
            get
            {
                float total = 0;
                foreach (var item in BetWaterList)
                {
                    if (item.Area == Pb.Bjl.EnumBetArea.AreaHe)
                    {
                        total += item.Count;
                    }
                }

                return total;
            }
        }
        
        public float ZhuangDuiTotal
        {
            get
            {
                float total = 0;
                foreach (var item in BetWaterList)
                {
                    if (item.Area == Pb.Bjl.EnumBetArea.AreaZhuangDui)
                    {
                        total += item.Count;
                    }
                }

                return total;
            }
        }
        
        public float XianDuiTotal
        {
            get
            {
                float total = 0;
                foreach (var item in BetWaterList)
                {
                    if (item.Area == Pb.Bjl.EnumBetArea.AreaXianDui)
                    {
                        total += item.Count;
                    }
                }

                return total;
            }
        }
        
        
       
    }
}