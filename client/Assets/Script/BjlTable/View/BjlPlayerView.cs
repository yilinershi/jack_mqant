using System;
using System.Collections.Generic;
using Pb.Bjl;
using UnityEngine;
using UnityEngine.UI;

public class BjlPlayerView : MonoBehaviour
{
    private BjlTableModel.BjlPlayerData data;
    private Text labelNickName;
    private Text labelUid;
    private Text labelGold;
    private Text labelXian;
    private Text labelZhuang;
    private Text labelHe;
    private Text labelZhuangDui;
    private Text labelXianDui;
    private Text labelWin;


    public void Init(BjlTableModel.BjlPlayerData pushData)
    {
        data = pushData;
        labelNickName = transform.Find("NickName").GetComponent<Text>();
        labelUid = transform.Find("UID").GetComponent<Text>();
        labelGold = transform.Find("Gold").GetComponent<Text>();
        labelXian = transform.Find("Xian").GetComponent<Text>();
        labelZhuang = transform.Find("Zhuang").GetComponent<Text>();
        labelHe = transform.Find("He").GetComponent<Text>();
        labelZhuangDui = transform.Find("ZhuangDui").GetComponent<Text>();
        labelXianDui = transform.Find("XianDui").GetComponent<Text>();
        labelWin = transform.Find("Win").GetComponent<Text>();


        RefreshAll();
    }

    private void RefreshAll()
    {
        RefreshUid();
        RefreshNickName();
        RefreshGold();
        RefreshPlayerBetInfo();
        RefreshBetWin();
    }

    public void RefreshPlayerBetInfo()
    {
        RefreshBetXian();
        RefreshBetZhuang();
        RefreshBetHe();
        RefreshBetZhuangDui();
        RefreshBetXianDui();
        RefreshGold();
    }

    
    
    
    private void RefreshNickName()
    {
        labelNickName.text = data.BaseInfo.NickName;
    }

    private void RefreshUid()
    {
        labelUid.text = data.BaseInfo.UID.ToString();
    }

    public void RefreshGold()
    {
        labelGold.text = data.BaseInfo.Gold.ToString();
    }

    private void RefreshBetXian()
    {
        labelXian.text = data.XianTotal.ToString();
    }

    private void RefreshBetZhuang()
    {
        labelZhuang.text = data.ZhuangTotal.ToString();
    }

    private void RefreshBetHe()
    {
        labelHe.text = data.HeTotal.ToString();
    }

    private void RefreshBetXianDui()
    {
        labelXianDui.text = data.XianDuiTotal.ToString();
    }

    private void RefreshBetZhuangDui()
    {
        labelZhuangDui.text = data.ZhuangDuiTotal.ToString();
    }

    public void RefreshBetWin()
    {
        labelWin.text = data.WinCount.ToString();
    }
}