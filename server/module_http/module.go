/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package module_http

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"net/http"
)

var Module = func() module.Module {
	this := new(moduleHttp)
	return this
}

type moduleHttp struct {
	basemodule.BaseModule
}

func (self *moduleHttp) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "module_http"
}

func (self *moduleHttp) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

func (self *moduleHttp) OnInit(app module.App, settings *conf.ModuleSettings) {
	self.BaseModule.OnInit(self, app, settings)
}

func (self *moduleHttp) startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8088"}

	http.HandleFunc("/login/", self.login)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}

func (self *moduleHttp) Run(closeSig chan bool) {
	log.Info("module_http: starting HTTP server :8088")
	srv := self.startHttpServer()
	<-closeSig
	log.Info("module_http: stopping HTTP server")

	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}
	log.Info("module_http: done. exiting")
}

func (self *moduleHttp) OnDestroy() {
	self.BaseModule.OnDestroy()
}
