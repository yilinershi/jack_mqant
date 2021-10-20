using Pb.Enum;
using UnityEngine;

public class LobbyController
{
   
    public static async void CallAuth()
    {
        var topic = "SV_Lobby/HD_OnAuth/Call";
        var req = new Pb.Lobby.ReqAuth {Token = Session.Token,Account = Session.Account};
        var resp = await MqttManager.Instance.Call<Pb.Lobby.ReqAuth, Pb.Lobby.RespAuth>(topic, req);
        Debug.Log("【CallAuth】，[topic]=" + topic + "\n[req]=" + req + "\n[resp]=" + resp);
        if (resp.ErrCode != ErrorCode.Ok)
        {
            return;
        }

        Session.User.UID = resp.UID;
        Session.User.Diamond = resp.Diamond;
        Session.User.NickName = resp.NickName;
        Session.User.Sex = resp.Sex;
        Session.User.Icon = resp.Icon;
        Session.User.Gold = resp.Gold;
        //打开lobby面板
        var prefab = Resources.Load("Prefab/UILobby");
        var go =  Object.Instantiate(prefab) as GameObject;
        go.AddComponent<LobbyView>();
    }
}