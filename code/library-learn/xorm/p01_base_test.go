package xorm_learn

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func setup() *xorm.Engine {
	var engine *xorm.Engine

	engine, err := xorm.NewEngine("sqlite3", "./xorm_learn.db")

	if err != nil {
		panic("connect or create database failed")
	}

	//设置日志输出到文件
	//f, err := os.Create("sql.log")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	panic("create log file failed")
	//}
	//engine.SetLogger(log.NewSimpleLogger(f))

	// 默认显示日志级别是INFO
	// 设置在控制台打印sql语句
	engine.ShowSQL(true)
	engine.Logger().SetLevel(log.LOG_DEBUG)

	engine.SetMaxIdleConns(2)
	engine.SetMaxOpenConns(10)
	engine.SetConnMaxLifetime(time.Minute * 10)

	// 给 表格添加前缀 tbl_
	tblMapper := names.NewPrefixMapper(names.SnakeMapper{}, "tbl_")
	engine.SetTableMapper(tblMapper)

	// 建立数据库后非立刻连接
	err = engine.Ping()
	if err != nil {
		panic("ping database failed")
	}
	return engine
}

//创建 engine
func TestCreateEngine(t *testing.T) {
	engine := setup()
	engine.Sync2(new(User))
	// 一般情况下不关闭，程序关闭会自动关闭
	//engine.Close()

}

func TestInsert(t *testing.T) {
	user := new(User)
	user.Name = "stt"
	user.Age = 18
	user.Extend.Alias = "xx2"

	engine := setup()
	re, err := engine.Insert(user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("re", re)
	fmt.Println("user id", user.Id)
}

// 查询使用原生查询，返回map key 是 字段名称 value是该字段的byte[]
func TestQuery(t *testing.T) {
	engine := setup()
	results, err := engine.Query(`select * 
							from tbl_user`)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results)

	// result
	//=== RUN   TestQuery
	//[map[age:[49 56] created:[50 48 50 48 45 49 48 45 49 55 84 48 55 58 50 53 58 50 51 90] extend:[123 34 65 108 105 97 115 34 58 34 120 120 34 125] id:[49] name:[115 116 116] passwd:[] salt:[] updated:[49 54 48 50 57 49 57 53 50 51]]]
	//--- PASS: TestQuery (0.02s)
	//PASS

	// 返回的也是byte[]
	results2, err := engine.Table(&User{}).Where("name = ?", "stt").Query()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results2)

}

// 查询返回的是string
func TestQueryString(t *testing.T) {
	engine := setup()
	results, err := engine.QueryString(`select * from tbl_user`)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results)
}

// 返回的是interface 对象
func TestQueryInterface(t *testing.T) {
	engine := setup()
	results, err := engine.QueryInterface(`select * from tbl_user`)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results)
}

func TestGet(t *testing.T) {
	engine := setup()
	user := new(User)
	has, err := engine.Get(user)
	if has {
		fmt.Println(err)
		fmt.Println(has)
		fmt.Println(user)
	}
	fmt.Println("-------")
	// 可以给table设置一个别名
	has, err = engine.Alias("o").Where("o.name = ?", "stt").Get(user)
	if has {
		fmt.Println(user)
	}
}

func TestWhereAndGet(t *testing.T) {
	engine := setup()
	//user := new(User)
	//has, err := engine.Desc("id").Where("name = ?", "stt").And("age = ?", 18).Get(user)
	var user User
	has, err := engine.Desc("id").Where("name = ?", "stt").And("age = ?", 18).Get(&user)
	if err != nil {
		panic(err.Error())
	}
	if has {
		fmt.Println(user)
	}
}

// 查询得到多个结果
func TestFind(t *testing.T) {
	engine := setup()
	var users []User
	err := engine.Find(&users)
	if err != nil {
		panic(err.Error())
	}
	for i := range users {
		fmt.Println(users[i])
	}
}

func TestUpdate(t *testing.T) {
	engine := setup()
	// 条件
	targetUser := &User{Name: "stt"}
	// 更新的值
	valueUser := &User{Name: "xx", Age: 12}
	// 除非使用Cols,AllCols函数指明，默认只更新非空和非0的字段
	re, err := engine.Update(valueUser, targetUser)
	if err != nil {
		panic(err.Error())
	}
	// 返回修改成功个数
	fmt.Println(re)
}

func TestDelete(t *testing.T) {
	engine := setup()
	// 强制删除 id = 1 的记录
	re, err := engine.ID(1).Delete(&User{})
	fmt.Println(err)
	fmt.Println(re)
}

func TestCount(t *testing.T) {
	engine := setup()
	counts, err := engine.Count(&User{})
	fmt.Println(err)
	fmt.Println(counts)
}

func TestTransaction(t *testing.T) {
	user := &User{Name: "yy", Salt: "11"}
	engine := setup()
	re, err := engine.Transaction(func(session *xorm.Session) (interface{}, error) {
		// 在此处进行事务的处理
		re, err := session.Insert(user)
		return re, err
	})

	fmt.Println(err)
	fmt.Println(re)
}

func TestQueryJson(t *testing.T) {
	// 启动的时候需要使用 go build --tags json1 启动，切换sqlite3不通的配置
	engine := setup()
	user := &User{}
	has, err := engine.Where("JSON_EXTRACT(extend,'$.Alias') = ? ", "xx").Get(user)
	fmt.Println(err)
	fmt.Println(has)
	fmt.Println(user)
}
