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

//registerHandle 注册客户端请求服务
func (this *SV_Lobby) registerHandle() {
	this.GetServer().RegisterGO("Call_Auth", this.callAuth)
}

func (this *SV_Lobby) callAuth(session gate.Session, topic string, req *pb_lobby.ReqAuth) {
	log.Info("[callAuth] session=%+v,req =%+v\n", session.GetSessionID(), req)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	a := new(pb_rpc.DbAccount)
	err := mqrpc.Proto(a, func() (reply interface{}, errStr interface{}) {
		return this.Call(ctx, "SV_DB", "rpcLoadAccount", mqrpc.Param(req.Account))
	})
	if err != nil {
		return
	}
	resp := new(pb_lobby.RespAuth)
	if a.Token == req.Token {
		u := new(pb_rpc.DbUser)
		mqrpc.Proto(u, func() (reply interface{}, errStr interface{}) {
			return this.Call(ctx, "SV_DB", "rpcLoadUser", mqrpc.Param(a.UID))
		})
		resp.ErrCode = pb_enum.ErrorCode_OK
		resp.UID = u.UID
		resp.Diamond = u.Diamond
		resp.Gold = u.Gold
		resp.Icon = u.Icon
		resp.NickName = u.NickName
		resp.Sex = u.Sex

		session.Bind(strconv.Itoa(int(a.UID)))
		session.Set("isLogin", "true")
		session.Set("account", a.Account)
		session.Set("uid", strconv.Itoa(int(a.UID)))
		session.Set("nickName", u.NickName)
		session.Push()
	} else {
		resp.ErrCode = pb_enum.ErrorCode_AuthFailed
	}
	log.Info("[callAuth] result=%v\n", resp)
	b, _ := proto.Marshal(resp)
	session.Send(topic, b)
}
