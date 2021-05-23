/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package module_gate

import (
	//"encoding/json"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/module"
	//"google.golang.org/protobuf/proto"
	"log"
	"server/module_gate/router_custom"
	//"server/pb/pb_lobby"
)

var Module = func() module.Module {
	gate := new(Gate)
	return gate
}

type Gate struct {
	basegate.Gate //继承
}

func (this *Gate) GetType() string {
	return "SV_Gate"
}

func (this *Gate) Version() string {
	return "1.0.0"
}

func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
	//注意这里一定要用 gate.Gate 而不是 module.BaseModule
	wsUrl := app.GetSettings().Settings["WebsocketUrl"].(string)
	tcpUrl := app.GetSettings().Settings["TcpUrl"].(string)
	this.Gate.OnInit(this, app, settings,
		gate.WsAddr(wsUrl),
		gate.TCPAddr(tcpUrl),
		gate.SetRouteHandler(router_custom.NewRouterCustom(this)),
	)
}

func (this *Gate) Connect(session gate.Session) {
	log.Println("gate connect")
	agent, err := this.GetGateHandler().GetAgent(session.GetSessionID())
	if err != nil {

	}
	agent.ConnTime()
}
