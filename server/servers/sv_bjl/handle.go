package sv_bjl

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_bjl"
	"server/pb/pb_common"
	"server/pb/pb_enum"
)

//registerRouter 注册客户端请求服务
func (this *SV_Bjl) registerRouter() {
	this.GetServer().RegisterGO("Call_JoinTable", this.callJoinTable)
	this.GetServer().RegisterGO("Call_Bet", this.callBet)
	this.GetServer().RegisterGO("Call_TableHeartbeat", this.callHeartbeat)
}

func (this *SV_Bjl) callHeartbeat(session gate.Session, topic string, req *pb_common.ReqHeartbeat) {
	log.Info("[callHeartbeat]  data=%+v\n", req)
	tableId := session.Get("tableId")
	if tableId == "" {
		log.Error("tableId is un exist")
		return
	}
	table := this.room.GetTable(tableId)
	table.PutQueue("Table/CallHeartbeat", session, topic, req)
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
	bytes, _ := proto.Marshal(resp)
	session.Send(topic, bytes)

	session.SetPush("tableId", req.TableId) //加入桌子后，绑定session所在的tableId
	table.PutQueue("Table/CallPlayerJoin", session, topic)
}

func (this *SV_Bjl) callBet(session gate.Session, topic string, data *pb_bjl.ReqBet) {
	log.Info("[callBet]  data=%+v\n", data)
	tableId := session.Get("tableId")
	if tableId == "" {
		log.Error("tableId is un exist")
		return
	}
	table := this.room.GetTable(tableId)
	table.PutQueue("Table/CallBet", session, topic, data)
}
