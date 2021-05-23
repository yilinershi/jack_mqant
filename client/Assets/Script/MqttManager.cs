using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using System.Xml.Serialization;
using M2MqttUnity;
using uPLibrary.Networking.M2Mqtt.Messages;
using Google.Protobuf;
using Google.Protobuf.Reflection;
using UnityEngine;

public class MqttManager : M2MqttUnityClient
{
    protected override void Start()
    {
        autoConnect = true;
        base.Start();
    }

    private readonly Dictionary<string, Action<byte[]>> _dicCall = new Dictionary<string, Action<byte[]>>();
    private readonly Dictionary<string, Action<byte[]>> _dicSync = new Dictionary<string, Action<byte[]>>();

    public Task<TResp> Call<TReq, TResp>(string topic, TReq req) where TReq : IMessage<TReq>, new() where TResp : IMessage<TResp>, new()
    {
        var data = req.ToByteArray();
        client.Publish(topic, data, MqttMsgBase.QOS_LEVEL_EXACTLY_ONCE, false);
        return GenTask<TResp>(topic);
    }
    
    public void Input<TInput>(string topic, TInput input) where TInput : IMessage<TInput>, new()
    {
        var data = input.ToByteArray();
        client.Publish(topic, data, MqttMsgBase.QOS_LEVEL_EXACTLY_ONCE, false);
    }


    public void OnSync<TResp>(string topic, Action<TResp> resp) where TResp : IMessage<TResp>, new()
    {
        if (_dicSync.ContainsKey(topic))
        {
            _dicSync[topic] += (msg) =>
            {
                TResp k = new MessageParser<TResp>(() => new TResp()).ParseFrom(msg);
                resp?.Invoke(k);
            };
        }
        else
        {
            _dicSync[topic] = (msg) =>
            {
                TResp k = new MessageParser<TResp>(() => new TResp()).ParseFrom(msg);
                resp?.Invoke(k);
            };
        }
    }
    

    private Task<T> GenTask<T>(string topic) where T : IMessage<T>, new()
    {
        var task = TaskUtil.GenSendTask(out Action<T> action, $"{typeof(T).Name} error");

        void Callback(byte[] msg)
        {
            T resp = new T();
            resp.MergeFrom(msg);
            action.Invoke(resp);
            action = null;
        }

        _dicCall.Add(topic, Callback);
        return task;
    }


    protected override void DecodeMessage(string topic, byte[] message)
    {
        
        //收到call类型的消息
        if (_dicCall.ContainsKey(topic))
        {
            Debug.Log("Received Call >> topic = " + topic);
            _dicCall[topic](message);
            _dicCall.Remove(topic);
            return;
        }

        //收到sync类型的消息
        if (_dicSync.ContainsKey(topic))
        {
            Debug.Log("Received Sync >> topic = " + topic);
            _dicSync[topic](message);
            _dicSync.Remove(topic);
            return;
        }
        
        string msg = System.Text.Encoding.UTF8.GetString(message);
        Debug.Log("Received Unknown >> topic = " + topic + " ,msg=" + msg);
    }

    protected override void OnConnected()
    {
        base.OnConnected();
        Debug.Log("连接成功");
        CallLogin();
    }

    public async void CallLogin()
    {
        var req = new Pb.Lobby.ReqLogin {Account = "ming", Password = "abc"};
        var resp = await Call<Pb.Lobby.ReqLogin, Pb.Lobby.RespLogin>("SV_Lobby/HD_Login/Call", req);
        Debug.Log("-------------收到回复消息---------");
        Debug.Log(resp);
    }
}