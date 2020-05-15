package dao

// 通用连接
import (
	"database/sql"

	// 用于添加mysql的driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

// 调用该go文件，会先执行该init方法
func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(192.168.205.10:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
