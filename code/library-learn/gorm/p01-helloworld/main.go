package main

// 快速入门

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	// 创建数据库
	db, err := gorm.Open(sqlite.Open("gorm_test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移创建表格
	db.AutoMigrate(&Product{})

	// 创建记录
	db.Create(&Product{Code: "Stt", Price: 11})

	// read
	var product Product
	// 通过主键查找
	db.First(&product, 1)
	// 通过条件查找，如果不存在则抛出异常
	db.First(&product, "code = ?", "Stt")
	fmt.Println(product)
	// 更新某个字段
	db.Model(&product).Update("Code", "s2")
	// 更新非空字段 注意是Updates 也可以更新多个字段
	db.Model(&product).Updates(Product{Code: "S3"})
}
