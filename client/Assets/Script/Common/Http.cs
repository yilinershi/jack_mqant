using System;
using Google.Protobuf;
using UnityEngine;

public static class HttpNet
{
    public static void Handshake(Action callback)
    {
        var req = new Pb.Http.ReqHandShake {Secret = "天王盖地虎,宝塔镇河妖"};
        UnityHTTP.Request theRequest = new UnityHTTP.Request("post", "http://127.0.0.1:8088/handshake", req.ToByteArray());
        Debug.Log( "[http entry],req = "+req.ToString());
        theRequest.Send((request) =>
        {
            var resp = new Pb.Http.RespHandShake();
            resp.MergeFrom(request.response.bytes);
            if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
            {
                Debug.LogError(resp.ErrCode.ToString());
                return;
            }
            
            Session.LoginUrl = resp.LoginUrl;
            Session.ResgisterUrl = resp.RegisterUrl;
            Session.TcpUrl = resp.TcpUrl;
            Session.WebSocketUrl = resp.WebSocketUrl;
            Debug.Log( "[http entry],resp = "+resp.ToString());

            callback();
        });
    }

}