package module_http

import (
	"context"
	"fmt"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/registry"
	mqrpc "github.com/liangdas/mqant/rpc"
	"github.com/liangdas/mqant/selector"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//login 用户通过http登录
func (self *moduleHttp) login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Info("parse http form error", err)
		return
	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	rstr, err := mqrpc.String(
		self.Call(
			ctx,
			"module_lobby", //要访问的moduleType
			"onLogin",      //访问模块中handler路径
			mqrpc.Param(r.Form.Get("name")),
			selector.WithStrategy(func(services []*registry.Service) selector.Next {
				var nodes []*registry.Node

				// Filter the nodes for datacenter
				for _, service := range services {
					if service.Version != "1.0.0" {
						continue
					}
					for _, node := range service.Nodes {
						nodes = append(nodes, node)
						if node.Metadata["state"] == "alive" || node.Metadata["state"] == "" {
							nodes = append(nodes, node)
						}
					}
				}

				var mtx sync.Mutex
				return func() (*registry.Node, error) {
					mtx.Lock()
					defer mtx.Unlock()
					if len(nodes) == 0 {
						return nil, fmt.Errorf("no node")
					}
					index := rand.Intn(int(len(nodes)))
					return nodes[index], nil
				}
			}),
		),
	)

	log.Info("RpcCall %v , err %v", rstr, err)
	if err != nil {
		_, _ = io.WriteString(w, err.Error())
	}
	_, _ = io.WriteString(w, rstr)
}
