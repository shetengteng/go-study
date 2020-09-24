package session

import (
	"reflect"
	"stt_orm/clause"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	// 数组的make 必须要有个默认大小
	//sqlParams := make([]interface{}, 0)
	var sqlParams []interface{}
	for _, value := range values {
		table := s.Model(value).RefTable()
		// insert 设置需要 表名，以及字段名称 合成INSERT INTO tableName (field1,field2...)
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		// 从对象中分析每个字段，获取值
		sqlParams = append(sqlParams, table.RecordValues(value))
	}
	// 给insert语句添加上values(?,?,?)
	s.clause.Set(clause.VALUES, sqlParams...)

	// 构造最后的sql
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)

	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// 接收的是数组对象，给接收的数组对象赋值
func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// 找到数组的元素的类型
	destType := destSlice.Type().Elem()
	// 初始化 schema
	// reflect.New 依据type 创建一个对象，并返回该对象的值的并转换为interface
	// 此处的Elem和Type().Elem不同
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()
	// set select 但是没有set 其他如 where orderby等
	// 因此build的时候只有 select 有效果
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		// 获取dest内字段的引用地址
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		// 将dest 的每个字段的地址放在values中，传递给Scan赋值
		err := rows.Scan(values...)
		if err != nil {
			return err
		}
		// 此时dest 赋值完成
		// 将dest 塞到 slice中去
		// destSlice 是一个reflect.Value对象，需要用reflect.Append进行切片的操作
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
