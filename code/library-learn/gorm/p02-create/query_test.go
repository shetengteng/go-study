package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

// 查询第一条记录，按照主键升序查询，自动添加LIMIT
//  SELECT * FROM users ORDER BY id LIMIT 1;
func TestQueryFirst(t *testing.T) {
	db := setup()
	var user User
	db.First(&user)
	fmt.Println(user)

	// 获取一条记录，没有排序 SELECT * FROM users LIMIT 1;
	var user2 User
	db.Take(&user2)
	fmt.Println(user2)

	// 获取最后一条记录
	// SELECT * FROM users ORDER BY id DESC LIMIT 1
	var user3 User
	db.Last(&user3)
	fmt.Println(user3)

	result := db.First(&user)
	fmt.Println("result row affected", result.RowsAffected)
	fmt.Println("result error", result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) != true {
		fmt.Println("---> Success Query")
	}

}
