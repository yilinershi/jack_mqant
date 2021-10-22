package sv_tetris

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"google.golang.org/protobuf/proto"
	"server/pb/pb_enum"
	"server/pb/pb_tetris"
	"server/servers/sv_tetris/entity"
)

//registerRouter 注册客户端请求服务
func (this *SV_Tetris) registerRouter() {
	this.GetServer().RegisterGO("Call_SubscribeRoomInfo", this.callSubscribeRoomInfo)
	this.GetServer().RegisterGO("Call_CreateTable", this.callCreateTable)
	this.GetServer().RegisterGO("Call_JoinTable", this.callJoinTable)
}

func (this *SV_Tetris) callSubscribeRoomInfo(session gate.Session, topic string, msg []byte) {
	req := new(pb_tetris.ReqSubscribeRoomInfo)
	if err := proto.Unmarshal(msg, req); err != nil {
		log.Info("proto Unmarshal err=%+v\n", err)
		return
	}
	uid := session.Get("uid")
	log.Info("[callSubscribeRoomInfo] uid=%s, req=%+v\n", uid, req)
	resp := new(pb_tetris.RespSubscribeRoomInfo)
	if req.IsSubscribe == false {
		if _, isOk := this.subscribeGroup[session.GetSessionID()]; isOk {
			delete(this.subscribeGroup, session.GetSessionID())
		}
		resp.ErrCode = pb_enum.ErrorCode_OK
	} else {
		this.subscribeGroup[session.GetSessionID()] = session
		resp.ErrCode = pb_enum.ErrorCode_OK
		temp := make([]*pb_tetris.TableInfo, 0)
		for _, info := range this.allTableInfo {
			temp = append(temp, info)
		}
		resp.AllTableInfo = temp
	}

	respByte, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	session.Send(topic, respByte)
}

//callCreateTable 用户创建桌子
func (this *SV_Tetris) callCreateTable(session gate.Session, topic string, msg []byte) {
	req := new(pb_tetris.ReqCreateTetris)
	if err := proto.Unmarshal(msg, req); err != nil {
		log.Info("proto Unmarshal err=%+v\n", err)
		return
	}
	log.Info("[callCreateTable]  req=%+v\n", req)
	var tableId = uuid.New().String()

	_, err := this.room.CreateById(this.GetApp(), tableId, func(app module.App, tableId string) (room.BaseTable, error) {
		table := entity.NewTable(
			app,
			room.TableId(tableId),
			room.Router(func(TableId string) string {
				return fmt.Sprintf("%v://%v/%v", this.GetType(), this.GetServerID(), tableId)
			}),
			room.DestroyCallbacks(func(table room.BaseTable) error {
				log.Info("【回收房间】: %v", table.TableId())
				_ = this.room.DestroyTable(table.TableId())
				tableInfo, isOk := this.allTableInfo[table.TableId()]
				if isOk {
					delete(this.allTableInfo,table.TableId())
					//这里播报房间销毁消息
					this.BroadcastRoomInfoChange(tableInfo, pb_tetris.PushRoomInfoChange_OnDestroy)
				}
				return nil
			}),
		)
		return table, nil
	})

	if err != nil {
		return
	}
	table := this.room.GetTable(tableId)
	tableInfo := &pb_tetris.TableInfo{
		TableId:         table.TableId(),
		Name:            req.TableName,
		CreatorUId:      session.Get("uid"),
		CreatorNickName: session.Get("nickName"),
	}
	this.allTableInfo[tableInfo.TableId] = tableInfo
	resp := new(pb_tetris.RespCreteTetris)
	resp.ErrCode = pb_enum.ErrorCode_OK
	respByte, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	log.Info("[callCreateTable]  resp=%+v\n", resp)
	session.Send(topic, respByte)

	//向全room广播房间里的桌子变动情况
	this.BroadcastRoomInfoChange(tableInfo, pb_tetris.PushRoomInfoChange_OnCreate)
}

func (this *SV_Tetris) BroadcastRoomInfoChange(tableInfo *pb_tetris.TableInfo, changeType pb_tetris.PushRoomInfoChange_EnumChangeType) {
	temp := make([]*pb_tetris.TableInfo, 0)
	for _, info := range this.allTableInfo {
		temp = append(temp, info)
	}
	pushData := &pb_tetris.PushRoomInfoChange{
		ChangeType:        changeType,
		OnChangeTableInfo: tableInfo,
		AllTableInfo:      temp,
	}
	bytes, err := proto.Marshal(pushData)
	if err != nil {
		return
	}
	for _, g := range this.subscribeGroup {
		g.Send("SV_Tetris/Push_RoomInfoChange", bytes)
	}

	log.Info("[BroadcastRoomInfoChange] pushData=%+v\n ", pushData)
}

//callJoinTable 用户加入桌子
func (this *SV_Tetris) callJoinTable(session gate.Session, topic string, msg []byte) {
	req := new(pb_tetris.ReqCreateTetris)
	if err := proto.Unmarshal(msg, req); err != nil {
		log.Info("proto Unmarshal err=%+v\n", err)
		return
	}
	log.Info("[callJoinTable]  req=%+v\n", req)
	var tableId = uuid.New().String()

	_, err := this.room.CreateById(this.GetApp(), tableId, func(app module.App, tableId string) (room.BaseTable, error) {
		table := entity.NewTable(
			app,
			room.TableId(tableId),
			room.Router(func(TableId string) string {
				return fmt.Sprintf("%v://%v/%v", this.GetType(), this.GetServerID(), tableId)
			}),
			room.DestroyCallbacks(func(table room.BaseTable) error {
				log.Info("【回收房间】: %v", table.TableId())
				_ = this.room.DestroyTable(table.TableId())
				tableInfo, isOk := this.allTableInfo[table.TableId()]
				if isOk {
					delete(this.allTableInfo,table.TableId())
					//这里播报房间销毁消息
					this.BroadcastRoomInfoChange(tableInfo, pb_tetris.PushRoomInfoChange_OnDestroy)
				}
				return nil
			}),
		)
		return table, nil
	})

	if err != nil {
		return
	}
	table := this.room.GetTable(tableId)
	tableInfo := &pb_tetris.TableInfo{
		TableId:         table.TableId(),
		Name:            req.TableName,
		CreatorUId:      session.Get("uid"),
		CreatorNickName: session.Get("nickName"),
	}
	this.allTableInfo[tableInfo.TableId] = tableInfo
	resp := new(pb_tetris.RespCreteTetris)
	resp.ErrCode = pb_enum.ErrorCode_OK
	respByte, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	log.Info("[callCreateTable]  resp=%+v\n", resp)
	session.Send(topic, respByte)

	//向全room广播房间里的桌子变动情况
	this.BroadcastRoomInfoChange(tableInfo, pb_tetris.PushRoomInfoChange_OnCreate)
}
