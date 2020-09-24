package stt_orm

import (
	"database/sql"
	"stt_orm/log"
	"stt_orm/session"
)

// 负责与用户进行交互的入口，session 负责与数据库交互
// 用于关闭链接等
type Engine struct {
	db *sql.DB
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
	e = &Engine{db: db}
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
	return session.New(e.db)
}
