package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

// 存储不同语言之间的区别，进行解耦
type Dialect interface {
	// 对go中的反射值的类型解析到sql中的类型
	// 将 Go 语言的类型转换为该数据库的数据类型
	DataTypeOf(t reflect.Value) string
	//返回某个表是否存在的 SQL 语句，参数是表名(table)
	// 不同数据库的执行sql有所不同
	TableExistSQL(tableName string) (string, []interface{})
	// 不同数据库的差异还有其他待补充
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
