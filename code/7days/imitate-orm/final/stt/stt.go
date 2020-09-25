package stt_orm

import (
	"database/sql"
	"fmt"
	"strings"
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

// 找到2个对象字段数组的不通点
// difference returns a - b
func difference(a []string, b []string) (diff []string) {
	mapB := make(map[string]bool)
	for _, v := range b {
		mapB[v] = true
	}
	for _, v := range a {
		if _, ok := mapB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}

// Migrate table
//CREATE TABLE new_table AS SELECT col1, col2, ... from old_table
//DROP TABLE old_table
//ALTER TABLE new_table RENAME TO old_table;
func (engine *Engine) Migrate(value interface{}) error {
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		if !s.Model(value).HasTable() {
			log.Infof("table %s doesn't exist", s.RefTable().Name)
			return nil, s.CreateTable()
		}
		table := s.RefTable()
		rows, _ := s.Raw(fmt.Sprintf("SELECT * FROM %s LIMIT 1", table.Name)).QueryRows()
		columns, _ := rows.Columns()
		addCols := difference(table.FieldNames, columns)
		delCols := difference(columns, table.FieldNames)
		log.Infof("added cols %v, deleted cols %v", addCols, delCols)

		for _, col := range addCols {
			f := table.GetField(col)
			sqlStr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", table.Name, f.Name, f.Type)
			if _, err = s.Raw(sqlStr).Exec(); err != nil {
				return
			}
		}

		if len(delCols) == 0 {
			return
		}
		tmp := "tmp_" + table.Name
		fieldStr := strings.Join(table.FieldNames, ", ")
		s.Raw(fmt.Sprintf("CREATE TABLE %s AS SELECT %s from %s;", tmp, fieldStr, table.Name))
		s.Raw(fmt.Sprintf("DROP TABLE %s;", table.Name))
		s.Raw(fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", tmp, table.Name))
		_, err = s.Exec()
		return
	})
	return err
}
