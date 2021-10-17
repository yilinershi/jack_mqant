package router_custom

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/module"
	argsutil "github.com/liangdas/mqant/rpc/util"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"strings"
)

type Option func(*RouterCustom)

type RouterCustom struct {
	module module.RPCModule
}

func NewRouterCustom(module module.RPCModule, opts ...Option) *RouterCustom {
	route := &RouterCustom{
		module: module,
	}
	if opts != nil {
		for _, o := range opts {
			o(route)
		}
	}
	return route
}

func (this *RouterCustom) OnRoute(session gate.Session, topic string, msg []byte) (bool, interface{}, error) {
	_, err := url.Parse(topic)
	if err != nil {
		return true, nil, errors.Errorf("topic is not url %v", err.Error())
	}
	topics := strings.Split(topic, "/")
	if len(topics) < 2 {
		return true, nil, errors.Errorf("Topic must be [SV]/[HD]/[Function]")
	}
	routerServer := topics[0]
	routerHandler := topics[1]
	//routerFunction := topics[2]

	if startsWith := strings.HasPrefix(routerServer, "SV_"); !startsWith {
		return true, nil, errors.Errorf("Server(%s) must begin with 'Server'", routerServer)
	}

	serverSession, err := this.module.GetRouteServer(routerServer)
	if err != nil {
		return true, nil, errors.Errorf("Service(type:%s) not found", routerServer)
	}

	if startsWith := strings.HasPrefix(routerHandler, "HD_"); !startsWith {
		return true, nil, errors.Errorf("Handler(%s) must begin with 'Handler'", routerHandler)
	}

	var ArgsType = make([]string, 2)
	ArgsType[0] = gate.RPCParamSessionType //默认，同default的内容
	ArgsType[1] = argsutil.BYTES           //所有客户端发给服务器的消息都为byte类型，节约流量
	var args = make([][]byte, 2)
	args[0] = msg //客户端发过来的参数，第1个消息为该内容（我们默认客户端给服务器只发一个参数的消息，该自定义路由不支持多参数）

	session = session.Clone()
	session.SetTopic(topic)

	//ctx, _ := context.WithTimeout(context.TODO(), this.module.GetApp().Options().RPCExpired)
	if b, err := session.Serializable(); err == nil {
		args[0] = b
	}
	e := serverSession.CallNRArgs(routerHandler, ArgsType, args)
	if e != nil {
		log.Println("Gate rpc", e.Error())
	}
	return false, nil, nil
	//
	///*
	//	Call：为req-resp结构，即客户端发消息给服务器，同时服务器对该消息进行响应
	//    Notify:为req-noResp结构，为客户端发消息给服务器，但服务器不作响应
	//    Push:为noReq-Resp结构,即服务器主动推送消息到客户端，不需要客户端作出响应
	//*/
	//
	//if isStartsWithCall := strings.HasPrefix(routerFunction, "Call"); isStartsWithCall {
	//	if b, err := session.Serializable(); err == nil {
	//		args[0] = b
	//	}
	//	result, e := serverSession.CallArgs(ctx, routerHandler, ArgsType, args)
	//	return false, result, errors.Errorf(e)
	//} else if isStartsWithNotify := strings.HasPrefix(routerFunction, "Notify"); isStartsWithNotify {
	//	if b, err := session.Serializable(); err == nil {
	//		args[0] = b
	//	}
	//	e := serverSession.CallNRArgs(topics[1], ArgsType, args)
	//	if e != nil {
	//		log.Println("Gate rpc", e.Error())
	//	}
	//	return false, nil, nil
	//}

	return false, nil, nil
}
