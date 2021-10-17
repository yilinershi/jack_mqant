using uPLibrary.Networking.M2Mqtt;
using uPLibrary.Networking.M2Mqtt.Messages;
using System.Security.Cryptography.X509Certificates;
using System.Net.Security;
using UnityEngine;

public static class MqttNet
{
    private static MqttClient mqttClient;

    public static void InitMqtt()
    {
        // ProtocolVersion = MqttProtocolVersion.V3_1;
        
        
        // 加载证书
        // var cert = Resources.Load("Txt/cacert") as TextAsset;
        // 使用TLS证书连接
        mqttClient = new MqttClient("localhost", 3563, true,null);
   //          , new X509Certificate(cert.bytes), new RemoteCertificateValidationCallback
			// (
			// 	// 测试服务器未设置公钥证书，返回true即跳过检查，直接通过，否则抛出服务器证书无效Error
			// 	(srvPoint, certificate, chain, errors) => true
			// ));  
            
            
            
        // 消息接收事件
        mqttClient.MqttMsgPublishReceived += msgReceived;
        // 连接
        mqttClient.Connect("client id","adimn","password");
        // 发送登录消息
        // mqttClient.Publish("Login/HD_Login/1", Encoding.UTF8.GetBytes("{\"userName\": \"username\",\"passWord\": \"Hello,anyone!\"}"));
    }

    private static void msgReceived(object sender, MqttMsgPublishEventArgs e)
    {
        Debug.Log("服务器返回数据");
        string msg = System.Text.Encoding.Default.GetString(e.Message);
        Debug.Log(msg);
    }

    public static void OnDisable()
    {
        if (mqttClient != null && mqttClient.IsConnected)
            mqttClient.Disconnect();
    }
}