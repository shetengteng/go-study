package schema

import (
	"go/ast"
	"reflect"
	"stt_orm/dialect"
)

// 对一个字段有3种描述
type Field struct {
	Name string // 字段的名称，作为表的列名
	Type string // 字段的类型，作为列的类型
	Tag  string // 字段的Tag，作为列的主键以及其他描述，约束条件
}

// 描述一个表
type Schema struct {
	Model      interface{}       // 映射的对象
	Name       string            // 表的名称
	Fields     []*Field          // 表中含有的字段
	FieldNames []string          // 表的列名
	fieldMap   map[string]*Field // 表的列名和字段的映射
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// 将对象解析为Schema对象
// 接收的对象是结构体的指针
func Parse(dest interface{}, d dialect.Dialect) *Schema {

	// Indirect 相当于获取value值的Elem
	// TypeOf() 和 ValueOf() 是 reflect 包最为基本也是最重要的 2 个方法，分别用来返回入参的类型和值
	// 因为设计的入参是一个对象的指针，因此需要 reflect.Indirect() 获取指针指向的实例
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(), // 结构体的名称作为表名
		fieldMap: make(map[string]*Field),
	}
	// 获取类型下所有的Field
	for i := 0; i < modelType.NumField(); i++ {
		structField := modelType.Field(i)
		// 判断字段是否是大写开头，并且非匿名字段
		if !structField.Anonymous && ast.IsExported(structField.Name) {
			field := &Field{
				Name: structField.Name,                                              // 字段名称
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(structField.Type))), // 获取字段的类型
			}
			// 解析tag
			if v, ok := structField.Tag.Lookup("stt"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, field.Name)
			schema.fieldMap[field.Name] = field
		}

	}
	return schema
}
