package module_http

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/registry"
	mqrpc "github.com/liangdas/mqant/rpc"
	"github.com/liangdas/mqant/selector"
	"io"
	"math/rand"
	"net/http"
	"server/pb/pb_common"
	"server/pb/pb_lobby"
	"sync"
	"time"
)

var key = "天王盖地虎"

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (self *moduleHttp) entry(w http.ResponseWriter, r *http.Request) {
	secret := md5V(key)
	req := &pb_lobby.ReqEntry{
		Secret: r.PostFormValue("Secret"),
	}

	log.Info("[entry],secret=%s, req.Secret=%s\n", secret, req.Secret)
	resp := new(pb_lobby.RespEntry)
	//if secret == req.Secret {
		resp.ErrCode = pb_common.ErrorCode_OK
		resp.LoginUrl = self.loginUrl
		resp.RegisterUrl = self.registerUrl
		resp.WebSocketUrl = self.websocketUrl
		resp.TcpUrl = self.tcpUrl
	//} else {
	//	resp.ErrCode = pb_common.ErrorCode_EntryError
	//}

	jsonData, err := json.Marshal(resp)
	log.Info("[entry] result=%v\n", resp)
	if err != nil {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
}

//login 用户通过http登录
func (self *moduleHttp) login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Info("parse http form error", err)
		return
	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	result, err := mqrpc.String(
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

	log.Info("RpcCall %v , err %v", result, err)
	if err != nil {
		_, _ = io.WriteString(w, err.Error())
	}
	_, _ = io.WriteString(w, result)
}
