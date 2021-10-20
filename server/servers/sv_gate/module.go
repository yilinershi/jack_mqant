package sv_gate

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
)

func NewServerGate() module.Module {
	return new(SV_Gate)
}

type SV_Gate struct {
	basegate.Gate
}

func (this *SV_Gate) GetType() string {
	return "SV_Gate"
}

func (this *SV_Gate) Version() string {
	return "1.0.0"
}

func (this *SV_Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
	//wsUrl := app.GetSettings().Settings["WebsocketUrl"].(string)
	tcpUrl := app.GetSettings().Settings["TcpUrl"].(string)
	this.Gate.OnInit(this, app, settings,
		gate.TCPAddr(tcpUrl),
		//gate.WsAddr(wsUrl),
	)
	this.Gate.SetCreateAgent(this.CreateCustomAgent)
}

func (this *SV_Gate)CreateCustomAgent() gate.Agent{
	agent:= NewCustomAgent(this)
	return agent
}

func (m *SV_Gate) Connect(session gate.Session) {
	log.Info("session connect from %v-%v", session.GetNetwork(), session.GetIP())
}

func (m *SV_Gate) DisConnect(session gate.Session) {
	log.Info("session disconnect from %v-%v", session.GetNetwork(), session.GetIP())
}
