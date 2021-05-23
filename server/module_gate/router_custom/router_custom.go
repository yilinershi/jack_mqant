package router_custom

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/module"
	"github.com/pkg/errors"
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
	if len(topics) < 3 {
		return true, nil, errors.Errorf("Topic must be [moduleID]/[handler]/[msgType]")
	}
	moduleType := topics[0]
	handlerFunc := topics[1]
	messageType := topics[2]
	if startsWith := strings.HasPrefix(moduleType, "SV_"); !startsWith {
		return true, nil, errors.Errorf("Server(%s) must begin with 'SV_'", moduleType)
	}
	if startsWith := strings.HasPrefix(handlerFunc, "HD_"); !startsWith {
		return true, nil, errors.Errorf("Handler(%s) must begin with 'HD_'", handlerFunc)
	}
	if _, err := this.module.GetRouteServer(moduleType);err != nil {
		return true, nil, errors.Errorf("Service(type:%s) not found", moduleType)
	}
	if startsWith := strings.HasPrefix(messageType, "Call"); startsWith {
		if e := this.module.InvokeNR(moduleType, handlerFunc, session, topic, msg); e != nil {
			return true, nil, e
		}
	}
	if startsWith := strings.HasPrefix(messageType, "Sync"); startsWith {
		if e := this.module.InvokeNR(moduleType, handlerFunc, session, topic, msg); e != nil {
			return true, nil, e
		}
	}

	return false, nil, nil
}
