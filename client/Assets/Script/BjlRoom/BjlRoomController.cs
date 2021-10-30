using System.Collections.Generic;
using Pb.Bjl;
using UnityEngine;
public static class BjlRoomController
{

    public static List<TableInfo> allTableInfos = new List<TableInfo>();
    
    public static async void CallSubscribeRoomInfo(bool isSubscribe)
    {
        var topic = "SV_DB/Call_SubscribeRoomInfo";
        var req = new ReqSubscribeRoomInfo() {IsSubscribe = isSubscribe};
        var resp = await MqttManager.Instance.Call<ReqSubscribeRoomInfo, RespSubscribeRoomInfo>(topic, req);
        Debug.Log("【CallSubscribeRoomInfo】，[topic]=" + topic + "\n[req]=" + req + "\n[resp]=" + resp);
        
        
        allTableInfos = new List<TableInfo>();
        foreach (var item in resp.AllTableInfo)
        {
            allTableInfos.Add(item);
        }
        
        var prefab = Resources.Load("Prefab/UIBjlRoom");
        var go = Object.Instantiate(prefab) as GameObject;
        go.AddComponent<BjlRoomView>();
    }

    public static void OnPushRoomInfo(PushRoomInfoChange data)
    {
        Debug.Log("【OnPushRoomInfo】，onPush,data=" + data);
        allTableInfos = new List<TableInfo>();
        foreach (var item in data.AllTableInfo)
        {
            allTableInfos.Add(item);
        }
        BjlRoomView.Instance.RefreshTableList();
    }
    
    
    //创建俄罗斯方法
    public static async void CallCreateTable( string tableName)
    {
        var topic = "SV_Bjl/Call_CreateTable";
        var req = new ReqCreateBjl { TableName = tableName};
        var resp = await MqttManager.Instance.Call<ReqCreateBjl, RespCreteBjl>(topic, req);
        Debug.Log("【CallAuth】，[topic]=" + topic + "\n[req]=" + req + "\n[resp]=" + resp);
        if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
        {
            return;
        }
    }
    
    
    
}