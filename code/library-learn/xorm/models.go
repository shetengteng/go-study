package xorm_learn

import (
	"time"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`                 // 转换为UTC时间
	Updated int64     `xorm:"updated comment('更新时间')"` // 转换为时间戳秒数
	Extend  Extend    `xorm:"json"`
}

type Extend struct {
	Alias string
}
