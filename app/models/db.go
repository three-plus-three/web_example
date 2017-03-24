package models

import (
	"strings"

	"github.com/go-xorm/xorm"
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
		}

		if err := engine.CreateUniques(bean); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return err
			}
		}
	}
	return nil
}
