package main

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
	Role     string
	Updated  int64  `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Extend   Extend `gorm:"type:json"`
}

type Extend struct {
	Alias string
}

// 创建前的钩子函数
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Role == "Admin" {
		return errors.New("invalid role")
	}
	if user.Role == "" {
		user.Age = 90
	}
	return
}
