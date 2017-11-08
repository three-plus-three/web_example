package models

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"github.com/runner-mei/orm"
)

type DB struct {
	Engine  *xorm.Engine
	session *xorm.Session
}

func (db *DB) WithSession(sess *xorm.Session) *DB {
	return &DB{Engine: db.Engine, session: sess}
}

func (db *DB) Begin() (*DB, error) {
	if db.session != nil {
		return nil, errors.New("run in the transaction")
	}
	session := db.Engine.NewSession()
	return &DB{Engine: db.Engine, session: session}, nil
}

func (db *DB) Commit() error {
	if db.session == nil {
		return sql.ErrTxDone
	}
	err := db.session.Commit()
	db.session = nil
	return err
}

func (db *DB) Rollback() error {
	if db.session == nil {
		return sql.ErrTxDone
	}
	err := db.session.Rollback()
	db.session = nil
	return err
}

func (db *DB) Close() error {
	if db.session == nil {
		return sql.ErrTxDone
	}
	db.session.Close()
	db.session = nil
	return nil
}

func (db *DB) Query(sqlStr string, args ...interface{}) orm.Queryer {
	return orm.NewWithNoInstance()(db.Engine).
		WithSession(db.session).
		Query(sqlStr, args...)
}
func (db *DB) OnlineUsers() *orm.Collection {
	return orm.New(func() interface{} {
		return &OnlineUser{}
	}, KeyForOnlineUsers)(db.Engine).WithSession(db.session)
}
func (db *DB) AuthAccounts() *orm.Collection {
	return orm.New(func() interface{} {
		return &AuthAccount{}
	}, KeyForAuthAccounts)(db.Engine).WithSession(db.session)
}

func DropTables(engine *xorm.Engine) error {
	beans := []interface{}{
		&OnlineUser{},
		&AuthAccount{},
	}

	for _, bean := range beans {
		if err := engine.DropIndexes(bean); err != nil {
			if !strings.Contains(err.Error(), "does not exist") &&
				!strings.Contains(err.Error(), "不存在") {
				return err
			}
		}
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
			if !strings.Contains(err.Error(), "already exists") &&
				!strings.Contains(err.Error(), "已经存在") {
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
