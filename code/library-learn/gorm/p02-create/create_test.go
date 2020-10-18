package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

var user = User{Name: "xx", Age: 88, Birthday: time.Now()}

func setup() *gorm.DB {
	// 创建数据库
	db, err := gorm.Open(sqlite.Open("gorm_test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 创建合并表
	db.AutoMigrate(&User{})
	return db
}

// 创建 插入数据
func TestCreate(t *testing.T) {
	db := setup()

	user := &User{Name: "stt2", Age: 12, Birthday: time.Now()}
	result := db.Create(user)

	fmt.Println("ID------>", user.ID)
	fmt.Println("result.Error", result.Error)
	fmt.Println("result.RowsAffected", result.RowsAffected)
}

// 创建 插入指定字段数据
func TestCreateSelect(t *testing.T) {
	db := setup()
	// 创建一条记录，只插入Name字段
	db.Select("Name").Create(&user)
}

// 插入一条记录，忽略Name字段
func TestCreateOmit(t *testing.T) {
	db := setup()
	db.Omit("Name").Create(&user)
}
