package sv_gate

import (
	"bufio"
	"context"
	"fmt"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/gate/base/mqtt"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/network"
	mqrpc "github.com/liangdas/mqant/rpc"
	mqanttools "github.com/liangdas/mqant/utils"
	"github.com/pkg/errors"
	"runtime"
	"strings"
	"sync"
	"time"
)

//NewCustomAgent 自定义agent复制来源于源mqant_agent,仅对recoverWorker该方法进行了修改
func NewCustomAgent(module module.RPCModule) *CustomAgent {
	a := &CustomAgent{
		module: module,
	}
	return a
}

type CustomAgent struct {
	gate.Agent
	module                       module.RPCModule
	session                      gate.Session
	conn                         network.Conn
	r                            *bufio.Reader
	w                            *bufio.Writer
	gate                         gate.Gate
	client                       *mqtt.Client
	ch                           chan int //控制模块可同时开启的最大协程数
	isclose                      bool
	protocol_ok                  bool
	lock                         sync.Mutex
	lastStorageHeartbeatDataTime time.Duration //上一次发送存储心跳时间
	revNum                       int64
	sendNum                      int64
	connTime                     time.Time
}

func (this *CustomAgent) OnInit(gate gate.Gate, conn network.Conn) error {
	this.ch = make(chan int, gate.Options().ConcurrentTasks)
	this.conn = conn
	this.gate = gate
	this.r = bufio.NewReaderSize(conn, gate.Options().BufSize)
	this.w = bufio.NewWriterSize(conn, gate.Options().BufSize)
	this.isclose = false
	this.protocol_ok = false
	this.revNum = 0
	this.sendNum = 0
	this.lastStorageHeartbeatDataTime = time.Duration(time.Now().UnixNano())
	return nil
}
func (this *CustomAgent) IsClosed() bool {
	return this.isclose
}

func (this *CustomAgent) ProtocolOK() bool {
	return this.protocol_ok
}

func (this *CustomAgent) GetSession() gate.Session {
	return this.session
}

func (this *CustomAgent) Wait() error {
	// 如果ch满了则会处于阻塞，从而达到限制最大协程的功能
	select {
	case this.ch <- 1:
	//do nothing
	default:
		//warnning!
		return fmt.Errorf("the work queue is full!")
	}
	return nil
}
func (this *CustomAgent) Finish() {
	// 完成则从ch推出数据
	select {
	case <-this.ch:
	default:
	}
}

func (this *CustomAgent) Run() (err error) {
	defer func() {
		if err := recover(); err != nil {
			buff := make([]byte, 1024)
			runtime.Stack(buff, false)
			log.Error("conn.serve() panic(%v)\n info:%s", err, string(buff))
		}
		this.Close()

	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				buff := make([]byte, 1024)
				runtime.Stack(buff, false)
				log.Error("OverTime panic(%v)\n info:%s", err, string(buff))
			}
		}()
		select {
		case <-time.After(this.gate.Options().OverTime):
			if this.GetSession() == nil {
				//超过一段时间还没有建立mqtt连接则直接关闭网络连接
				this.Close()
			}
		}
	}()

	//握手协议
	var pack *mqtt.Pack
	pack, err = mqtt.ReadPack(this.r, this.gate.Options().MaxPackSize)
	if err != nil {
		log.Error("Read login pack error %v", err)
		return
	}
	if pack.GetType() != mqtt.CONNECT {
		log.Error("Recive login pack's type error:%v \n", pack.GetType())
		return
	}
	conn, ok := (pack.GetVariable()).(*mqtt.Connect)
	if !ok {
		log.Error("It's not this mqtt connection package.")
		return
	}
	c := mqtt.NewClient(conf.Conf.Mqtt, this, this.r, this.w, this.conn, conn.GetKeepAlive(), this.gate.Options().MaxPackSize)
	this.client = c
	addr := this.conn.RemoteAddr()
	this.session, err = this.gate.NewSessionByMap(map[string]interface{}{
		"Sessionid": mqanttools.GenerateID().String(),
		"Network":   addr.Network(),
		"IP":        addr.String(),
		"Serverid":  this.module.GetServerID(),
		"Settings":  make(map[string]string),
	})

	log.Info("on new session, session.sessionId=%s, ip=%s\n,", this.session.GetSessionID(), this.session.GetIP())

	if err != nil {
		log.Error("gate create agent fail", err.Error())
		return
	}
	this.session.JudgeGuest(this.gate.GetJudgeGuest())
	this.session.CreateTrace() //代码跟踪
	//回复客户端 CONNECT
	err = mqtt.WritePack(mqtt.GetConnAckPack(0), this.w)
	if err != nil {
		log.Error("ConnAckPack error %v", err.Error())
		return
	}
	this.connTime = time.Now()
	this.protocol_ok = true
	this.gate.GetAgentLearner().Connect(this) //发送连接成功的事件
	c.Listen_loop()                           //开始监听,直到连接中断
	return nil
}

func (this *CustomAgent) OnClose() error {
	defer func() {
		if err := recover(); err != nil {
			buff := make([]byte, 1024)
			runtime.Stack(buff, false)
			log.Error("agent OnClose panic(%v)\n info:%s", err, string(buff))
		}
	}()
	this.isclose = true
	this.gate.GetAgentLearner().DisConnect(this) //发送连接断开的事件
	return nil
}

func (this *CustomAgent) GetError() error {
	return this.client.GetError()
}

func (this *CustomAgent) RevNum() int64 {
	return this.revNum
}
func (this *CustomAgent) SendNum() int64 {
	return this.sendNum
}
func (this *CustomAgent) ConnTime() time.Time {
	return this.connTime
}
func (this *CustomAgent) OnRecover(pack *mqtt.Pack) {
	err := this.Wait()
	if err != nil {
		log.Error("Gate OnRecover error [%v]", err)
		pub := pack.GetVariable().(*mqtt.Publish)
		this.toResult(this, *pub.GetTopic(), nil, err.Error())
	} else {
		go this.recoverWorker(pack)
	}
}

func (this *CustomAgent) toResult(a *CustomAgent, Topic string, Result interface{}, Error string) error {
	switch v2 := Result.(type) {
	case module.ProtocolMarshal:
		return a.WriteMsg(Topic, v2.GetData())
	}
	b, err := a.module.GetApp().ProtocolMarshal(a.session.TraceId(), Result, Error)
	if err == "" {
		if b != nil {
			return a.WriteMsg(Topic, b.GetData())
		}
		return nil
	}
	br, _ := a.module.GetApp().ProtocolMarshal(a.session.TraceId(), nil, err)
	return a.WriteMsg(Topic, br.GetData())
}

func (this *CustomAgent) recoverWorker(pack *mqtt.Pack) {
	defer func() {
		this.lock.Lock()
		interval := int64(this.lastStorageHeartbeatDataTime) + int64(this.gate.Options().Heartbeat) //单位纳秒
		this.lock.Unlock()
		if interval < time.Now().UnixNano() {
			if this.gate.GetStorageHandler() != nil {
				this.lock.Lock()
				this.lastStorageHeartbeatDataTime = time.Duration(time.Now().UnixNano())
				this.lock.Unlock()
				this.gate.GetStorageHandler().Heartbeat(this.GetSession())
			}
		}
		this.Finish()
		if r := recover(); r != nil {
			buff := make([]byte, 1024)
			runtime.Stack(buff, false)
			log.Error("Gate recoverWorker error [%v] stack : %v", r, string(buff))
		}
	}()

	//路由服务
	switch pack.GetType() {
	case mqtt.PUBLISH:
		this.lock.Lock()
		this.revNum = this.revNum + 1
		this.lock.Unlock()
		pub := pack.GetVariable().(*mqtt.Publish)
		msg := pub.GetMsg()
		topic := *pub.GetTopic()
		log.Info("[router >> get message from client], topic=" + topic)
		topicSplitStr := strings.Split(topic, "/")
		if len(topicSplitStr) < 2 {
			this.WriteMsg("System/Error", []byte("路由层次不能低于2层"))
			return
		}
		routerSV := topicSplitStr[0]
		routerHD := topicSplitStr[1]
		isStartsWithSV := strings.HasPrefix(routerSV, "SV_")
		if !isStartsWithSV {
			this.WriteMsg("System/Error", []byte("路由第1层必需以SV_开头"))
			return
		}

		isStartsWithCall := strings.HasPrefix(routerHD, "Call_")
		if isStartsWithCall {
			this.module.Call(context.Background(), routerSV, routerHD, mqrpc.Param(this.session, topic, msg))
			return
		}

		isStartsWithNotify := strings.HasPrefix(routerHD, "Notify_")
		if isStartsWithNotify {
			this.module.Call(context.Background(), routerSV, routerHD, mqrpc.Param(this.session,msg))
			return
		}

		this.WriteMsg("System/Error", []byte("路由第2层必需以Call或Notify开头"))
	case mqtt.PINGREQ:
		//客户端发送的心跳包
		//if this.GetSession().GetUserId() != "" {
		//这个链接已经绑定Userid
		//this.lock.Lock()
		//interval := int64(this.lastStorageHeartbeatDataTime) + int64(this.gate.Options().Heartbeat) //单位纳秒
		//this.lock.Unlock()
		//if interval < time.Now().UnixNano() {
		//	if this.gate.GetStorageHandler() != nil {
		//		this.lock.Lock()
		//		this.lastStorageHeartbeatDataTime = time.Duration(time.Now().UnixNano())
		//		this.lock.Unlock()
		//		this.gate.GetStorageHandler().Heartbeat(this.GetSession())
		//	}
		//}
		//}
	}
}

func (this *CustomAgent) WriteMsg(topic string, body []byte) error {
	if this.client == nil {
		return errors.New("mqtt.Client nil")
	}
	this.sendNum++
	if this.gate.Options().SendMessageHook != nil {
		bb, err := this.gate.Options().SendMessageHook(this.GetSession(), topic, body)
		if err != nil {
			return err
		}
		body = bb
	}
	return this.client.WriteMsg(topic, body)
}

func (this *CustomAgent) Close() {
	go func() {
		//关闭连接部分情况下会阻塞超时，因此放协程去处理
		if this.conn != nil {
			this.conn.Close()
		}
	}()
}

func (this *CustomAgent) Destroy() {
	if this.conn != nil {
		this.conn.Destroy()
	}
}
