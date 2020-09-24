package main

import (
	"database/sql"
	"log"

	// 注意：需要安装gcc https://www.cnblogs.com/zsy/p/5958170.html
	// 导包时会注册sqlite3的驱动
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// 链接数据库，第一个参数是驱动的名称
	// 第二个参数是数据库的名称，也是文件名
	// 返回值是sql.DB的指针
	db, _ := sql.Open("sqlite3", "stt.db")
	defer func() { _ = db.Close() }()
	// Exec 执行sql语句，如果是查询语句，则不返回相应的记录，查询语句使用Query,QueryRow
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) VALUES(?),(?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	} else {
		log.Println(err)
	}
	// 使用QueryRow返回查询的结果
	// Query返回多条记录
	// QueryRow返回一条记录
	// 第一个入参是sql语句，第二入参是占位符 ? 对应的值，占位符一般用来防止SQL注入
	// 返回值是*sql.Row，row.Scan() 接受1个或多个参数，获取对应column的值
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	err = row.Scan(&name)
	if err == nil {
		log.Println(name)
	}
}
