package sv_bjl

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_bjl"
	"server/pb/pb_enum"
)

//registerRouter 注册客户端请求服务
func (this *SV_Bjl) registerRouter() {
	this.GetServer().RegisterGO("Call_JoinTable", this.callJoinTable)
}

//callJoinTable 用户加入桌子
func (this *SV_Bjl) callJoinTable(session gate.Session, topic string, req *pb_bjl.ReqJoinTable) {
	log.Info("[callJoinTable]  req=%+v\n", req)
	table := this.room.GetTable(req.TableId)
	resp := new(pb_bjl.RespJoinTable)
	if table == nil {
		resp.ErrCode = pb_enum.ErrorCode_TableUnExistent
		b, _ := proto.Marshal(resp)
		session.Send(topic, b)
		return
	}

	resp.ErrCode = pb_enum.ErrorCode_OK
	b, _ := proto.Marshal(resp)
	session.Send(topic, b)
	table.PutQueue("Table/CallPlayerJoin", session, topic)
}
