package session

import (
	"fmt"
	"reflect"
	"strings"
	"stt_orm/log"
	"stt_orm/schema"
)

// 数据库表的删除添加操作

func (s *Session) Model(value interface{}) *Session {
	// nil 或者 model有变化，则更新refTable
	// 如果传入的结构体名称不发生变化，则不会更新 refTable 的值
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

// 使用refTable中的数据库字段拼接SQL
func (s *Session) CreateTable() error {
	table := s.refTable
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.refTable.Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, params := s.dialect.TableExistSQL(s.refTable.Name)
	row := s.Raw(sql, params...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.refTable.Name
}
