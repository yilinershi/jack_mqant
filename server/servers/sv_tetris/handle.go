package sv_tetris

import (
	"fmt"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"server/servers/sv_tetris/entity"
)

//registerHandle 注册客户端请求服务
func (this *SV_Tetris) registerHandle() {
	this.GetServer().RegisterGO("HD_CreateTable", this.onCreateTable)
}

func (this *SV_Tetris) onCreateTable(session gate.Session, msg []byte) ([]byte, error) {
	var tableId = session.Get("uid")
	table := this.room.GetTable(tableId)

	if table != nil {
		return nil, nil
	}
	_, err := this.room.CreateById(this.GetApp(), tableId, this.NewTable)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *SV_Tetris) NewTable(module module.App, tableId string) (room.BaseTable, error) {
	table := entity.NewTable(
		module,
		room.TableId(tableId),
		room.Router(func(TableId string) string {
			return fmt.Sprintf("%v://%v/%v", self.GetType(), self.GetServerID(), tableId)
		}),
		room.DestroyCallbacks(func(table room.BaseTable) error {
			log.Info("回收了房间: %v", table.TableId())
			_ = self.room.DestroyTable(table.TableId())
			return nil
		}),
	)
	return table, nil
}
