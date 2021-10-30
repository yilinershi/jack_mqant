using System;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Pb.Bjl;
using Pb.Common;
using Pb.Enum;

public class BjlTableView : MonoBehaviour
{
    private GameObject playerGo;
    private Transform playerRoot;
    private InputField inputCount;
    private Button btn;
    private Text tip;
    private Text zhuangPoker;
    private Text xianPoker;
    private EnumBetArea curSelectBetArea=EnumBetArea.AreaNone;
    private uint betCount = 0;
    private const float HEARTBEAT_CD = 30;
    private float heartbeatTimer = 0;
    private Dictionary<long, BjlPlayerView> allPlayerView = new Dictionary<long, BjlPlayerView>();
    private void Awake()
    {
        playerGo = transform.Find("Canvas/Player").gameObject;
        playerRoot = transform.Find("Canvas/PlayerRoot");
        tip = transform.Find("Canvas/Tip").GetComponent<Text>();
        zhuangPoker = transform.Find("Canvas/ZhuangPoker").GetComponent<Text>();
        xianPoker = transform.Find("Canvas/XianPoker").GetComponent<Text>();
        inputCount = transform.Find("Canvas/InputField").GetComponent<InputField>();
        inputCount.onValueChanged.AddListener((value) =>
        {
            var count = uint.Parse(value);
            Debug.Log("count="+count);
            betCount = count;
        });
        
        
        for (var i = 0; i < 5; i++)
        {
            var index = i + 1;
            var toggle = transform.Find("Canvas/ToggleGroup/Toggle_" + index).GetComponent<Toggle>();
            toggle.onValueChanged.AddListener((isSelect) =>
            {
                if (isSelect)
                {
                     Debug.Log("toggle.name="+toggle.name);
                    if (toggle.name == "Toggle_1")
                    {
                        curSelectBetArea = EnumBetArea.AreaXian;
                    }

                    if (toggle.name == "Toggle_2")
                    {
                        curSelectBetArea = EnumBetArea.AreaZhuang;
                    }

                    if (toggle.name == "Toggle_3")
                    {
                        curSelectBetArea = EnumBetArea.AreaHe;
                    }

                    if (toggle.name == "Toggle_4")
                    {
                        curSelectBetArea = EnumBetArea.AreaXianDui;
                    }

                    if (toggle.name == "Toggle_5")
                    {
                        curSelectBetArea = EnumBetArea.AreaZhuangDui;
                    }
                }
            });
        }

       
        btn = transform.Find("Canvas/Button").GetComponent<Button>();
        btn.onClick.AddListener(() =>
        {
            if (betCount<=0)
            {
                Debug.Log("下注额度不能为空");
                return;
            }

            if (curSelectBetArea == EnumBetArea.AreaNone)
            {
                Debug.Log("下注区域不能为空");
                return;
            }

           
            BjlTableController.CallBet(curSelectBetArea, betCount);
        });
    }

    public void Refresh()
    {
        //清掉节点
        List<GameObject> list = new List<GameObject>();
        for (int i = 0; i < playerRoot.childCount; i++)
        {
            list.Add(playerRoot.GetChild(i).gameObject);
        }

        foreach (var child in list)
        {
            Destroy(child.gameObject);
        }


        allPlayerView = new Dictionary<long, BjlPlayerView>();
        foreach (var p in BjlTableController.model.AllPlayer)
        {
            var go = Instantiate(playerGo, playerRoot, true);
            go.SetActive(true);
            var playerView = go.AddComponent<BjlPlayerView>();
            playerView.Init(p);
            allPlayerView.Add(p.BaseInfo.UID,playerView);
        }
    }

    public BjlPlayerView FindPlayerView(long uid)
    {
        if (!allPlayerView.ContainsKey(uid))
        {
            return null;
        }

        return allPlayerView[uid];
    }

    public void OnGameStateReady()
    {
        btn.enabled = false;
        tip.text = "游戏开始";
        foreach (var item in this.allPlayerView)
        {
            item.Value.RefreshBetWin();
            item.Value.RefreshGold();
            item.Value.RefreshPlayerBetInfo();
        }
        
        zhuangPoker.gameObject.SetActive(false);
        xianPoker.gameObject.SetActive(false);
    }

    public void OnGameStateBet()
    {
        btn.enabled = true;
        tip.text = "请下注";
    }
    
    public void OnGameStateSend()
    {
        btn.enabled = false;
        tip.text = "发牌中";
    }
    
    public void OnGameStateShow()
    {
        btn.enabled = false;
        tip.text = "开牌";

        var strZhuang = "庄：";
        foreach (var item in BjlTableController.model.Zhaung)
        {
            strZhuang += pokerToStr(item);
        }
        zhuangPoker.gameObject.SetActive(true);
        zhuangPoker.text = strZhuang;
        
        var strXian = "闲：";
        foreach (var item in BjlTableController.model.Xian)
        {
            strXian += pokerToStr(item);
        }
        xianPoker.gameObject.SetActive(true);
        xianPoker.text = strXian;
    }

    private string pokerToStr(Poker p)
    {
        var str = "";
        switch (p.Hua)
        {
            case PokerHua.Tao:
                str += "♠";
                break;
            case PokerHua.Xin:
                str += "♥";
                break;
            case PokerHua.Mei:
                str +=  "♣";
                break;
            case PokerHua.Fang:
                str += "♦";
                break;
        }

        switch (p.Point)
        {
            case PokerPoint.Point2:
                str += "2";
                break;
            case PokerPoint.Point3:
                str += "3";
                break;
            case PokerPoint.Point4:
                str += "4";
                break;
            case PokerPoint.Point5:
                str += "5";
                break;
            case PokerPoint.Point6:
                str += "6";
                break;
            case PokerPoint.Point7:
                str += "7";
                break;
            case PokerPoint.Point8:
                str += "8";
                break;
            case PokerPoint.Point9:
                str += "9";
                break;
            case PokerPoint.PointT:
                str += "10";
                break;
            case PokerPoint.PointJ:
                str += "J";
                break;
            case PokerPoint.PointQ:
                str += "Q";
                break;
            case PokerPoint.PointK:
                str += "K";
                break;
            case PokerPoint.PointA:
                str += "A";
                break;
        }
        
        return str;
    }

    public void OnGameStateSettle()
    {
        btn.enabled = false;
        var str = "【";
        if (BjlTableController.model.RoundResult.IsZhuangDui )
        {
            str += "庄对";
        }
        if (BjlTableController.model.RoundResult.IsXianDui )
        {
            str += "闲对";
        }
        if (BjlTableController.model.RoundResult.WinType == EnumWinType.Xian)
        {
            str += "闲赢";
        }
        if (BjlTableController.model.RoundResult.WinType == EnumWinType.Zhuang)
        {
            str += "庄赢";
        } 
        if (BjlTableController.model.RoundResult.WinType == EnumWinType.He)
        {
            str += "和";
        }

        str += "】";

        tip.text = "结算:" + str;
        foreach (var item in this.allPlayerView)
        {
            item.Value.RefreshBetWin();
            item.Value.RefreshGold();
        }
    }

    private void Update()
    {
        heartbeatTimer += Time.deltaTime;
        if (heartbeatTimer >= HEARTBEAT_CD)
        {
            heartbeatTimer = 0;
            BjlTableController.CallTableHeartbeat();
        }
    }
}