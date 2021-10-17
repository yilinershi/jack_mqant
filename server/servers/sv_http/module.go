/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package sv_http

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"net/http"
)

func NewServerHttp() module.Module {
	return new(ServerHttp)
}

type ServerHttp struct {
	basemodule.BaseModule
	loginUrl     string
	registerUrl  string
	websocketUrl string
	tcpUrl       string
}

func (this *ServerHttp) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "SV_Http"
}

func (this *ServerHttp) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

func (this *ServerHttp) OnInit(app module.App, settings *conf.ModuleSettings) {
	this.BaseModule.OnInit(this, app, settings)
	this.SetListener(this)
	this.loginUrl = app.GetSettings().Settings["LoginUrl"].(string)
	this.registerUrl = app.GetSettings().Settings["RegisterUrl"].(string)
	this.websocketUrl = app.GetSettings().Settings["WebsocketUrl"].(string)
	this.tcpUrl = app.GetSettings().Settings["TcpUrl"].(string)
}

func (this *ServerHttp) startHttpServer() *http.Server {
	srv := &http.Server{
		Addr: ":8088",
	}
	http.HandleFunc("/entry", this.entry)
	http.HandleFunc("/login", this.login)
	http.HandleFunc("/register", this.register)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}

func (this *ServerHttp) Run(closeSig chan bool) {
	log.Info("ServerHttp: starting HTTP server :8088")
	srv := this.startHttpServer()
	<-closeSig
	log.Info("ServerHttp: stopping HTTP server")

	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}
	log.Info("ServerHttp: done. exiting")
}

func (this *ServerHttp) OnDestroy() {
	this.BaseModule.OnDestroy()
}
