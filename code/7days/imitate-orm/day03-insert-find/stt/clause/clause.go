package clause

import "strings"

type Type int

// 子句的类型
const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

type Clause struct {
	sql       map[Type]string
	sqlParams map[Type][]interface{}
}

// 从执行器中获取执行子句
func (c *Clause) Set(name Type, params ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlParams = make(map[Type][]interface{})
	}
	sql, sqlParams := generatorMap[name](params...)
	c.sql[name] = sql
	c.sqlParams[name] = sqlParams
}

// 构建子句，将所有子句整合在一起
func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string
	var params []interface{}
	for _, order := range orders {
		sql, ok := c.sql[order]
		if ok {
			sqls = append(sqls, sql)
			params = append(params, c.sqlParams[order]...)
		}
	}
	return strings.Join(sqls, " "), params
}
