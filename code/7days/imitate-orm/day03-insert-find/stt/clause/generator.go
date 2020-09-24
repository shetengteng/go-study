package clause

import (
	"fmt"
	"strings"
)

// 实现各个子句的生成规则
// 返回sql，含有占位符
// 返回sql的参数，可变参数
type generator func(values ...interface{}) (string, []interface{})

// Type 是int值，表示各个子句的类型
var generatorMap map[Type]generator

func init() {
	generatorMap = make(map[Type]generator)
	generatorMap[INSERT] = _insert
	generatorMap[VALUES] = _values
	generatorMap[SELECT] = _select
	generatorMap[LIMIT] = _limit
	generatorMap[WHERE] = _where
	generatorMap[ORDERBY] = _orderBy
}

// insert子句
// 第一个参数是 string类型，表示tableName
// 第二个参数是 []string类型，表示要插入的值，是列的名称
// 返回sql，以及param,
// 可以直接返回一个sql，而param为空数组
var _insert generator = func(values ...interface{}) (string, []interface{}) {
	// INSERT INTO $tableName ($fields)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

// 注意每个value是一个数组，表示一行
var _values generator = func(values ...interface{}) (string, []interface{}) {
	// VALUES($v1),($v2)
	var bindStr string
	var sql strings.Builder
	var params []interface{}
	sql.WriteString("VALUES ")
	for index, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVarsStr(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if index+1 != len(values) {
			// 添加 , 如(?,?,?,?), (?,?,?,?)
			sql.WriteString(", ")
		}
		params = append(params, v...)
	}
	return sql.String(), params
}

// 第一个参数是tableName
// 第二个参数是数组 []string 要查询的字段
var _select generator = func(values ...interface{}) (string, []interface{}) {
	// SELECT $fields FROM $tableName
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

var _limit generator = func(values ...interface{}) (string, []interface{}) {
	// LIMIT $num
	return "LIMIT ?", values
}

// WHERE desc --> desc 可以是 ? = ? ,params --> 可以是这2个的值
var _where generator = func(values ...interface{}) (string, []interface{}) {
	// WHERE $desc
	desc, params := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), params
}

var _orderBy generator = func(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

// 依照数目产生 防止sql注入的 占位符串，如 ?, ?, ?
func genBindVarsStr(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}
