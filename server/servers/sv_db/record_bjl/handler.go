package record_bjl

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_bjl"
	"server/pb/pb_enum"
)

func (this *RecordBjlComponent) callSubscribeRoomInfo(session gate.Session, topic string, req *pb_bjl.ReqSubscribeRoomInfo) {
	uid := session.Get("uid")
	log.Info("[callSubscribeRoomInfo] uid=%s, req=%+v\n", uid, req)
	resp := new(pb_bjl.RespSubscribeRoomInfo)
	if req.IsSubscribe == false {
		if _, isOk := this.subscribeGroup[session.GetSessionID()]; isOk {
			delete(this.subscribeGroup, session.GetSessionID())
		}
		resp.ErrCode = pb_enum.ErrorCode_OK
	} else {
		this.subscribeGroup[session.GetSessionID()] = session
		resp.ErrCode = pb_enum.ErrorCode_OK
	}
	log.Info("[callSubscribeRoomInfo]  resp=%+v\n", resp)
	b, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	session.Send(topic, b)
}

func (this *RecordBjlComponent) pushRoomInfoChang() {
	for _, session := range this.subscribeGroup {
		session.Send("Record_Bjl/RoomInfoChange", nil)
	}
}
