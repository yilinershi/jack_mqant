package serverGate

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"server/servers/serverGate/router_custom"
)

func NewServerGate() module.Module {
	return new(ServerGate)
}

type ServerGate struct {
	basegate.Gate
}

func (this *ServerGate) GetType() string {
	return "SV_Gate"
}

func (this *ServerGate) Version() string {
	return "1.0.0"
}

func (this *ServerGate) OnInit(app module.App, settings *conf.ModuleSettings) {
	//wsUrl := app.GetSettings().Settings["WebsocketUrl"].(string)
	tcpUrl := app.GetSettings().Settings["TcpUrl"].(string)
	//gate.WsAddr(":3653"),
	//注意这里一定要用 gate.ServerGate 而不是 module.BaseModule
	this.Gate.OnInit(this, app, settings,
		gate.TCPAddr(tcpUrl),
		//gate.WsAddr(wsUrl),
		gate.SetRouteHandler(router_custom.NewRouterCustom(this)),
	)
	this.Gate.SetCreateAgent(this.CreateCustomAgent)
}

func (this *ServerGate)CreateCustomAgent() gate.Agent{
	agent:= NewCustomAgent(this)
	return agent
}

func (m *ServerGate) Connect(session gate.Session) {
	log.Info("session connect from %v-%v", session.GetNetwork(), session.GetIP())
}

func (m *ServerGate) DisConnect(session gate.Session) {
	log.Info("session disconnect from %v-%v", session.GetNetwork(), session.GetIP())
}
