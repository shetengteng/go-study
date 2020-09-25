package session

import (
	"database/sql"
	"strings"
	"stt_orm/clause"
	"stt_orm/dialect"
	"stt_orm/log"
	"stt_orm/schema"
)

type Session struct {
	db        *sql.DB
	dialect   dialect.Dialect
	tx        *sql.Tx // 增加对事务的支持，如果为nil，则使用db进行操作
	refTable  *schema.Schema
	clause    clause.Clause // 子句
	sql       strings.Builder
	sqlParams []interface{}
}

// 定义接口 db 执行的最小集合
type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

// 重写 DB 接口
func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlParams = nil
	s.clause = clause.Clause{} // 重置 clause
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

// 查询一行
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	return s.DB().QueryRow(s.sql.String(), s.sqlParams...)
}

// 查询多行
func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	rows, err := s.DB().Query(s.sql.String(), s.sqlParams...)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}
