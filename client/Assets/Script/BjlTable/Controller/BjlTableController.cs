using System.Collections.Generic;
using Pb.Bjl;
using UnityEngine;

public static class TetrisTableController
{
    private static TetrisTableView view;

    public static async void CallJoinTetrisTable(string tableId)
    {
        const string topic = "SV_Tetris/Call_JoinTable";
        var req = new ReqJoinTable {TableId = tableId};
        var resp = await MqttManager.Instance.Call<ReqJoinTable, RespJoinTable>(topic, req);
        Debug.Log("【CallJoinTetrisTable】，[topic]=" + topic + "\n[req]=" + req + "\n[resp]=" + resp);
        if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
        {
            return;
        }

        var prefab = Resources.Load("Prefab/UITetrisTable");
        var go = Object.Instantiate(prefab) as GameObject;
        view = go.AddComponent<TetrisTableView>();
    }


    public static void PushTablePlayerChange(PushTablePlayerChange data)
    {
        
        TetrisTableModel.AllPlayer = new List<BjlPlayer>();
        foreach (var p in data.AllPlayer)
        {
            TetrisTableModel.AllPlayer.Add(p);
        }

        view.Refresh();
    }
}