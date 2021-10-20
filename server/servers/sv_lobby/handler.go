package sv_lobby

import (
	"context"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	mqrpc "github.com/liangdas/mqant/rpc"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
	"server/pb/pb_rpc"
	"strconv"
	"time"
)

//handler 用来处理来自客户端的消息
//rpc  用来处理来自其它服务器的rpc消息

//registerHandle 注册客户端请求服务
func (this *SV_Lobby) registerHandle() {
	this.GetServer().RegisterGO("HD_OnAuth", this.onAuth)
}

func (this *SV_Lobby) onAuth(session gate.Session, msg []byte) ([]byte, error) {
	log.Info("onAuth, session=%+v\n", session.GetSessionID())
	req := new(pb_lobby.ReqAuth)
	if err := proto.Unmarshal(msg, req); err != nil {
		log.Info("err---------")
		return nil, err
	}

	log.Info("onAuth req =%+v\n", req)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	a := new(pb_rpc.DbAccount)
	err := mqrpc.Proto(a, func() (reply interface{}, errStr interface{}) {
		return this.Call(ctx, "SV_DB", "rpcLoadAccount", mqrpc.Param(req.Account))
	})
	log.Info("RpcCall ,account=%+v ,err= %v", a, err)

	resp := new(pb_lobby.RespAuth)
	if a.Token == req.Token {
		u := new(pb_rpc.DbUser)
		err2 := mqrpc.Proto(u, func() (reply interface{}, errStr interface{}) {
			return this.Call(ctx, "SV_DB", "rpcLoadUser", mqrpc.Param(a.UID))
		})
		log.Info("RpcCall ,a.UID=%d ,err= %v", a.UID, err2)
		resp.ErrCode = pb_enum.ErrorCode_OK
		resp.UID = u.UID
		resp.Diamond = u.Diamond
		resp.Gold = u.Gold
		resp.Icon = u.Icon
		resp.NickName = u.NickName
		resp.Sex = u.Sex
	} else {
		resp.ErrCode = pb_enum.ErrorCode_AuthFailed
	}

	session.Set("isLogin", "true")
	session.Set("account", a.Account)
	session.Set("uid", strconv.Itoa(int(a.UID)))
	session.Push()
	respByte, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return respByte, nil
}

