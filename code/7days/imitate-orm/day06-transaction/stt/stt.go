package stt_orm

import (
	"database/sql"
	"stt_orm/dialect"
	"stt_orm/log"
	"stt_orm/session"
)

// 负责与用户进行交互的入口，session 负责与数据库交互
// 用于关闭链接等
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// 发送一个ping 保证数据库链接通畅
	err = db.Ping()
	if err != nil {
		log.Error(err)
		return
	}

	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s not found", driver)
		return
	}
	e = &Engine{db: db, dialect: dialect}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	err := e.db.Close()
	if err != nil {
		log.Error("failed to close database")
	}
	log.Info("close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}

type TxFunc func(*session.Session) (interface{}, error)

func (e *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := e.NewSession()
	err = s.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			// 执行过程中有异常
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			// 执行完成返回值有异常
			_ = s.Rollback()
		} else {
			// 最后提交
			err = s.Commit()
		}
	}()

	return f(s)
}
