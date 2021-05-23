package module_lobby

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_common"

	//"github.com/gogo/protobuf/proto"
	"github.com/liangdas/mqant/gate"
	"log"
	"server/pb/pb_lobby"
)

func (self *module_lobby) onTest(session gate.Session, msg map[string]interface{}) (r string, err error) {
	name := msg["name"]
	return fmt.Sprintf("hi %v", name), nil
}

func (self *module_lobby) onRegister(session gate.Session, msg map[string]interface{}) (r string, err error) {
	log.Printf("msg=%+v\n", msg)
	session.Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))
	return fmt.Sprintf("hi %v 你在网关 %v", msg["name"], session.GetSessionID()), nil
}
//
////
//func (self *module_lobby) onLogin(msg *pb_lobby.ReqLogin) (*rpcpb.ResultInfo, error) {
//	log.Printf("hi =%+v\n", msg.Account)
//	//
//	//req := new(pb_lobby.ReqLogin)
//	//if err := proto.Unmarshal(msg.Result, req); err != nil {
//	//	return nil, nil
//	//}
//	//log.Printf("hi =%+v\n", req.Account)
//
//	resp := new(pb_lobby.RespLogin)
//	resp.ErrCode = pb_common.ErrorCode_OK
//	data, _ := proto.Marshal(resp)
//	r := &rpcpb.ResultInfo{
//		Result: data,
//	}
//
//	return r, nil
//}
//
func (this *module_lobby) onLogin(session gate.Session,topic string, msg []byte) error {
	log.Println(msg)
	log.Println("_____________login_______________")
	//解析客户端发送过来的user.LoginRequest结构体
	//req := &pb_lobby.ReqLogin{}
	req := new(pb_lobby.ReqLogin)

	if err := proto.Unmarshal(msg, req); err != nil {
		log.Println("err---------")
		return err
	}

	log.Printf("hi =%+v\n", req.Account)

	resp:=new (pb_lobby.RespLogin)
	resp.ErrCode=pb_common.ErrorCode_OK
	respByte, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	log.Println("respByte=",respByte)
	session.Send(topic,respByte)
	return nil
}

//func (this *module_lobby) login(session gate.Session, msg map[string]interface{}) (result module.ProtocolMarshal, err string) {
//	return this.App.ProtocolMarshal("","login success","")
//}
