using System.Collections.Generic;
using Pb.Bjl;
using Pb.Common;
using Pb.Enum;
using UnityEngine;

public static class BjlTableController
{
    private static BjlTableView view;
    public static BjlTableModel model;
    public static async void CallJoinTable(string tableId)
    {
        const string topic = "SV_Bjl/Call_JoinTable";
        var req = new ReqJoinTable {TableId = tableId};
        var resp = await MqttManager.Instance.Call<ReqJoinTable, RespJoinTable>(topic, req);
        Debug.Log("【CallJoinTetrisTable】，[topic]=" + topic + "\n[req]=" + req + "\n[resp]=" + resp);
        if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
        {
            return;
        }

        var prefab = Resources.Load("Prefab/UIBjlTable");
        var go = Object.Instantiate(prefab) as GameObject;
        view = go.AddComponent<BjlTableView>();

        model = new BjlTableModel();
    }

    public static async void CallTableHeartbeat()
    {
        const string topic = "SV_Bjl/Call_TableHeartbeat";
        var req = new ReqHeartbeat() {Ping = "ping"};
        var resp = await MqttManager.Instance.Call<ReqHeartbeat, RespHeartbeat>(topic, req);
        Debug.Log("【CallTableHeartbeat】，[topic]=" + topic + "\n[req]=" + req + "\n[resp]=" + resp);
    }

    public static void BroadcastTablePlayerChange(BroadcastTablePlayerChange data)
    {
        Debug.Log("【BroadcastTablePlayerChange】，[data]=" + data);
        model.AllPlayer = new List<BjlTableModel.BjlPlayerData>();
        foreach (var p in data.AllPlayer)
        {
            var newPlayer = new BjlTableModel.BjlPlayerData
            {
                BaseInfo = p
            };
            model.AllPlayer.Add(newPlayer);
        }


        view.Refresh();
    }

    public static void BroadcastGameStatusReady(BroadcastStatusReady data)
    {
        Debug.Log("【BroadcastGameStatusReady】，[data]=" + data);
        foreach (var p in model.AllPlayer)
        {
            p.BetWaterList = new List<BetInfo>();
        }

        view.OnGameStateReady();
    }

    public static void BroadcastGameStatusBet(BroadcastStatusBet data)
    {
        Debug.Log("【BroadcastGameStatusBet】，[data]=" + data);
        view.OnGameStateBet();
    }

    public static void BroadcastGameStatusShow(BroadcastStatusShow data)
    {
        Debug.Log("【BroadcastGameStatusShow】，[data]=" + data);

        model.Xian = new List<Poker>();
        foreach (var item in data.Xian)
        {
            model.Xian.Add(item);
        }
        
        model.Zhaung = new List<Poker>();
        foreach (var item in data.Zhuang)
        {
            model.Zhaung.Add(item);
        }

        
        view.OnGameStateShow();
    }


    public static void BroadcastGameStatusSend(BroadcastStatusSend data)
    {
        Debug.Log("【BroadcastGameStatusSend】，[data]=" + data);

      
        
        view.OnGameStateSend();
    }

    public static void BroadcastGameStatusSettle(BroadcastStatusSettle data)
    {
        Debug.Log("【BroadcastGameStatusSettle】，[data]=" + data);
        model.RoundResult = data.Result;
        foreach (var info in data.Info)
        {
            var player = model.FindPlayerByUid(info.UID);
            if (player != null)
            {
                player.BaseInfo.Gold = info.Gold;
                player.WinCount = info.GoldChange;
            }
        }
        
        view.OnGameStateSettle();
    }

    public static void BroadcastPlayerBet(BroadcastPlayerBet data)
    {
        Debug.Log("【BroadcastPlayerBet】，[data]=" + data);
        var player = model.FindPlayerByUid(data.Info.UID);
        if (player == null)
        {
            return;
        }

        player.BetWaterList.Add(new BetInfo()
        {
            Area = data.Info.Area,
            Count = data.Info.Count,
        });
        
        var playerView = view.FindPlayerView(data.Info.UID);
        playerView.RefreshPlayerBetInfo();
    }

    public static async void CallBet(EnumBetArea area, uint count)
    {
        const string topic = "SV_Bjl/Call_Bet";
        var req = new ReqBet()
        {
            Area = area,
            Count = count,
        };
        var resp = await MqttManager.Instance.Call<ReqBet, RespBet>(topic, req);
        Debug.Log("【NotifyBet】，[topic]=" + topic + "\n[req]=" + req+ "\n[resp]=" + resp);

        if (resp.ErrCode == ErrorCode.Ok)
        {
            var myPlayer = model.FindPlayerByUid(Session.User.UID);
            myPlayer.BaseInfo.Gold = resp.Gold;

            var myPlayerView = view.FindPlayerView(Session.User.UID);
            myPlayerView.RefreshGold();
        }
    }
}