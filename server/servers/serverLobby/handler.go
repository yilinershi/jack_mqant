package serverLobby

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
)

//handler 用来处理来自客户端的消息
//rpc  用来处理来自其它服务器的rpc消息

func (self *ServerLobby) onAuth(session gate.Session,msg []byte) ([]byte, error) {

	log.Info("onAuth, session=%+v\n",session.GetSessionID())
	req := new(pb_lobby.ReqAuth)

	if err := proto.Unmarshal(msg, req); err != nil {
		log.Info("err---------")
		return nil, err
	}

	log.Info("onAuth req =%+v\n", req)

	resp := new(pb_lobby.RespAuth)
	resp.ErrCode = pb_enum.ErrorCode_OK
	respByte, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}
	//session.Send(topic, respByte)
	return respByte, nil
}

//func (this *ServerLobby) onLogin(session gate.Session,topic string, msg []byte) error {
//	log.Println(msg)
//	log.Println("_____________login_______________")
//	//解析客户端发送过来的user.LoginRequest结构体
//	//req := &pb_lobby.ReqLogin{}
//	req := new(pb_lobby.ReqLogin)
//
//	if err := proto.Unmarshal(msg, req); err != nil {
//		log.Println("err---------")
//		return err
//	}
//
//	log.Printf("hi =%+v\n", req.Account)
//
//	resp:=new (pb_lobby.RespLogin)
//	resp.ErrCode=pb_common.ErrorCode_OK
//	respByte, err := proto.Marshal(resp)
//	if err != nil {
//		return err
//	}
//	log.Println("respByte=",respByte)
//	session.Send(topic,respByte)
//	return nil
//}

//func (this *ServerLobby) login(session gate.Session, msg map[string]interface{}) (result module.ProtocolMarshal, err string) {
//	return this.App.ProtocolMarshal("","login success","")
//}
