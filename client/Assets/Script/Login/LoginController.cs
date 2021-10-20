using System;
using Google.Protobuf;
using UnityEngine;
using uPLibrary.Networking.M2Mqtt;
using uPLibrary.Networking.M2Mqtt;
using uPLibrary.Networking.M2Mqtt.Messages;
using System.Text;
using System.Security.Cryptography.X509Certificates;
using System.Net.Security;

public static class LoginController
{
    public static void Register(string account, string password, Action callback = null)
    {
        var req = new Pb.Http.ReqRegister() {Account = account, Password = password};
        UnityHTTP.Request theRequest = new UnityHTTP.Request("post", "Http://" + Session.ResgisterUrl, req.ToByteArray());
        theRequest.Send((request) =>
        {
            var resp = new Pb.Http.RespRegister();
            resp.MergeFrom(request.response.bytes);

            if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
            {
                Debug.LogError("register error," + resp);
                return;
            }

            Debug.Log("register result=" + resp);
            callback?.Invoke();
        });
    }


    public static void Login(string account, string password, Action callback = null)
    {
        var req = new Pb.Http.ReqLogin() {Account = account, Password = password};
        UnityHTTP.Request theRequest = new UnityHTTP.Request("post", "Http://" + Session.LoginUrl, req.ToByteArray());
        Debug.Log("[http Login],req = " + req.ToString());
        theRequest.Send((request) =>
        {
            var resp = new Pb.Http.RespLogin();
            resp.MergeFrom(request.response.bytes);

            if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
            {
                Debug.LogError(resp.ErrCode.ToString());
                return;
            }

            Debug.Log("[http Login],resp = " + resp.ToString());
            Session.Account = account;
            Session.Token = resp.Token;
            // MqttNet.InitMqtt();
            MqttManager.Instance.Init();

            callback?.Invoke();
        });
    }


    
}