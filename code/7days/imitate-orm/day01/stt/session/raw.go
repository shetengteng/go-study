package session

import (
	"database/sql"
	"strings"
	"stt_orm/log"
)

type Session struct {
	db        *sql.DB
	sql       strings.Builder
	sqlParams []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlParams = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

// 给原生的sql赋值
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlParams = append(s.sqlParams, values...)
	return s
}

// 执行原生的 sql
func (s *Session) Exec() (sql.Result, error) {
	// 每次执行完，清空sql，便于下次执行
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	// 使用 = 直接返回对应的返回值
	result, err := s.DB().Exec(s.sql.String(), s.sqlParams...)
	if err != nil {
		log.Error(err)
	}
	return result, err
}

// 查询一列
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	return s.DB().QueryRow(s.sql.String(), s.sqlParams...)
}

// 查询多列
func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	rows, err := s.DB().Query(s.sql.String(), s.sqlParams)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}
