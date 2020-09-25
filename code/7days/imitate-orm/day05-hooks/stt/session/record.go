package session

import (
	"errors"
	"reflect"
	"stt_orm/clause"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {

	// 数组的make 必须要有个默认大小
	//sqlParams := make([]interface{}, 0)
	var sqlParams []interface{}
	for _, value := range values {
		// 添加钩子函数 对每个对象都插入前都进行判断操作
		s.CallMethod(BeforeInsert, value)
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
	// 添加钩子函数
	s.CallMethod(AfterInsert, nil)

	return result.RowsAffected()
}

// 接收的是数组对象，给接收的数组对象赋值
func (s *Session) Find(values interface{}) error {

	// 执行钩子函数
	s.CallMethod(BeforeQuery, nil)

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
		// 执行查询完成一条调用一次，使用的是dest对象
		// 可以对每个操作对象进行操作，在该dest对象中添加AfterQuery 方法就可以生效
		s.CallMethod(AfterQuery, dest.Addr().Interface())

		// 此时dest 赋值完成
		// 将dest 塞到 slice中去
		// destSlice 是一个reflect.Value对象，需要用reflect.Append进行切片的操作
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}

// 第一种：默认支持map的键值对
// 第二种：支持 k1,v1,k2,v2的数组入参
func (s *Session) Update(kv ...interface{}) (int64, error) {
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		// 说明是第二种入参情况，需要转换为第一种情况
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	sql, params := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.Raw(sql, params...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Delete() (int64, error) {
	s.clause.Set(clause.DELETE, s.RefTable().Name)
	sql, params := s.clause.Build(clause.DELETE, clause.WHERE)
	res, err := s.Raw(sql, params...).Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, params := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, params...).QueryRow()
	var temp int64
	err := row.Scan(&temp)
	if err != nil {
		return 0, err
	}
	return temp, nil
}

func (s *Session) Limit(num int) *Session {
	s.clause.Set(clause.LIMIT, num)
	return s
}

func (s *Session) Where(desc string, args ...interface{}) *Session {
	var vars []interface{}
	vars = append(vars, desc)
	s.clause.Set(clause.WHERE, append(vars, args...)...)
	return s
}

func (s *Session) OrderBy(desc string) *Session {
	s.clause.Set(clause.ORDERBY, desc)
	return s
}

// 得到第一个值，入参是返回值的引用 &User{}，这里是查询函数
func (s *Session) First(value interface{}) error {
	dest := reflect.Indirect(reflect.ValueOf(value))
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()

	err := s.Limit(1).Find(destSlice.Addr().Interface())
	if err != nil {
		return err
	}
	if destSlice.Len() == 0 {
		return errors.New("NOT FOUND")
	}
	dest.Set(destSlice.Index(0))
	return nil
}
