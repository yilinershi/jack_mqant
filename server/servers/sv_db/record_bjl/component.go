package record_bjl

import (
	"github.com/gomodule/redigo/redis"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/server"
)

type RecordBjlComponent struct {
	server         server.Server
	redisConn      redis.Conn
	subscribeGroup map[string]gate.Session //订阅百家乐时时动态的session集合
}

func NewRecordBjl(server server.Server, conn redis.Conn) *RecordBjlComponent {
	r := &RecordBjlComponent{
		server:         server,
		redisConn:      conn,
		subscribeGroup: make(map[string]gate.Session, 0),
	}
	return r
}

func (this *RecordBjlComponent) Register() {
	this.server.RegisterGO("Call_SubscribeRoomInfo", this.callSubscribeRoomInfo)
}


