package mysql

import "github.com/dinzhen12306/rpc-one/config"

func TX(fun func() error) {
	tx := config.XDB.NewSession()
	tx.Begin()
	if err := fun(); err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
