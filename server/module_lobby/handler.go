package module_lobby

import (
	"fmt"
	"github.com/liangdas/mqant/gate"
	"log"
)

func (self *module_lobby) onRegister(name string) (r string, err error) {
	return fmt.Sprintf("hi %v", name), nil
}

func (self *module_lobby) onLogin(session gate.Session, msg map[string]interface{}) (r string, err error) {
	//session.Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))
	log.Printf("name = %s login success",msg["name"])
	return fmt.Sprintf("hi %v 你在网关 %v", msg["name"], session.GetSessionID()), nil
}
