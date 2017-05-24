package models

import (
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"github.com/runner-mei/orm"
)

type DB struct {
	Engine *xorm.Engine
}

func (db *DB) OnlineUsers() *orm.Collection {
	return orm.New(func() interface{} {
		return &OnlineUser{}
	})(db.Engine)
}
func (db *DB) AuthAccounts() *orm.Collection {
	return orm.New(func() interface{} {
		return &AuthAccount{}
	})(db.Engine)
}

func DropTables(engine *xorm.Engine) error {
	beans := []interface{}{
		&OnlineUser{},
		&AuthAccount{},
	}

	return engine.DropTables(beans...)
}

func InitTables(engine *xorm.Engine) error {
	beans := []interface{}{
		&OnlineUser{},
		&AuthAccount{},
	}

	if err := engine.CreateTables(beans...); err != nil {
		return err
	}

	for _, bean := range beans {
		if err := engine.CreateIndexes(bean); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return err
			}
			revel.WARN.Println(err)
		}

		if err := engine.CreateUniques(bean); err != nil {
			if !(strings.Contains(err.Error(), "already exists") ||
				strings.Contains(err.Error(), "已经存在")) &&
				!(strings.Contains(err.Error(), "create unique index") ||
					strings.Contains(err.Error(), "UQE_")) {
				return err
			}
			revel.WARN.Println(err)
		}
	}
	return nil
}
