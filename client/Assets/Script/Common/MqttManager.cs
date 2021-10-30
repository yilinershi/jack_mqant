using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using M2MqttUnity;
using uPLibrary.Networking.M2Mqtt.Messages;
using Google.Protobuf;
using UnityEngine;

public class MqttManager : M2MqttUnityClient
{
    public static MqttManager Instance;

    protected void Start()
    {
        Instance = this;
        // this.brokerPort=Session.WebSocketUrl
    }

    public void Init()
    {
        base.Connect();
        Instance.RegisterPush<Pb.Bjl.PushRoomInfoChange>("SV_DB/Record_Bjl/RoomInfoChange", BjlRoomController.OnPushRoomInfo);
        Instance.RegisterPush<Pb.Bjl.BroadcastTablePlayerChange>("SV_Bjl/Table/BroadcastTablePlayerChange", BjlTableController.BroadcastTablePlayerChange);
        Instance.RegisterPush<Pb.Bjl.BroadcastStatusReady>("SV_Bjl/Table/BroadcastStateReady", BjlTableController.BroadcastGameStatusReady);
        Instance.RegisterPush<Pb.Bjl.BroadcastStatusBet>("SV_Bjl/Table/BroadcastStateBet", BjlTableController.BroadcastGameStatusBet);
        Instance.RegisterPush<Pb.Bjl.BroadcastStatusSend>("SV_Bjl/Table/BroadcastStateSend", BjlTableController.BroadcastGameStatusSend);
        Instance.RegisterPush<Pb.Bjl.BroadcastStatusShow>("SV_Bjl/Table/BroadcastStateShow", BjlTableController.BroadcastGameStatusShow);
        Instance.RegisterPush<Pb.Bjl.BroadcastStatusSettle>("SV_Bjl/Table/BroadcastStateSettle", BjlTableController.BroadcastGameStatusSettle);
        Instance.RegisterPush<Pb.Bjl.BroadcastPlayerBet>("SV_Bjl/Table/BroadcastTablePlayerBet", BjlTableController.BroadcastPlayerBet);
    }

    private readonly Dictionary<string, Action<byte[]>> _dicCall = new Dictionary<string, Action<byte[]>>();
    private readonly Dictionary<string, Action<byte[]>> _dicPush = new Dictionary<string, Action<byte[]>>();

    public Task<TResp> Call<TReq, TResp>(string topic, TReq req) where TReq : IMessage<TReq>, new() where TResp : IMessage<TResp>, new()
    {
        var data = req.ToByteArray();
        client.Publish(topic, data, MqttMsgBase.QOS_LEVEL_EXACTLY_ONCE, false);
        var t = GenTask<TResp>(topic);
        return t;
    }

    public void Input<TInput>(string topic, TInput input) where TInput : IMessage<TInput>, new()
    {
        var data = input.ToByteArray();
        client.Publish(topic, data, MqttMsgBase.QOS_LEVEL_EXACTLY_ONCE, false);
    }

    private void RegisterPush<TPush>(string topic, Action<TPush> resp) where TPush : IMessage<TPush>, new()
    {
        // if (!Instance._dicPush.ContainsKey(topic))
        // {
            Instance._dicPush.Add(topic, (msg) =>
            {
                TPush k = new MessageParser<TPush>(() => new TPush()).ParseFrom(msg);
                resp?.Invoke(k);
            });
        // }
    }

    private static Task<T> GenTask<T>(string topic) where T : IMessage<T>, new()
    {
        var task = TaskUtil.GenSendTask(out Action<T> action, $"{typeof(T).Name} error");
        void Callback(byte[] msg)
        {
            T resp = new T();
            resp.MergeFrom(msg);
            action.Invoke(resp);
            action = null;
        }
        Instance._dicCall[topic] = Callback;
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
        if (_dicPush.ContainsKey(topic))
        {
            Debug.Log("Received push >> topic = " + topic);
            _dicPush[topic](message);
            _dicCall.Remove(topic);
            return;
        }

        if (topic == "System/Error")
        {
            Debug.LogError(System.Text.Encoding.UTF8.GetString(message));
            return;
        }

        string msg = System.Text.Encoding.UTF8.GetString(message);
        Debug.LogWarning("Received Unknown >> topic = " + topic + " ,msg=" + msg);
    }

    protected override void OnConnected()
    {
        base.OnConnected();
        Debug.Log("连接成功");
        LobbyController.CallAuth();
    }
}