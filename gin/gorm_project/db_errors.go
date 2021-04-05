package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_project/models/relate_tables"
	"fmt"
)

func main() {
	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8&parseTime=True&loc=Local
	db,err := gorm.Open("mysql","root:Qazwsx123@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()   // 关闭空闲的连接


	var users []relate_tables.User

	result := db.Where("name = ?","hallen10").Find(&users)
	fmt.Println(users)
	fmt.Println(result.RowsAffected)

	fmt.Println(result.Error)

	if result.Error != nil { // 发生错误
		// 发生错误
	}

	//
	errs := result.GetErrors()

	for _,err := range errs{
		fmt.Println(err)
	}



	// ErrRecordNotFound
	var user relate_tables.User
	ret := db.Where("name = ?","hallen10").First(&user)
	fmt.Println(ret.Error)

	// ErrInvalidSQL
	var user2 relate_tables.User
	ret2 := db.Where("name = ?","hallen10").First(&user2)
	fmt.Println(ret2.Error)


	// 事务

	ct := db.Begin() // 开启事务
	ret3 := ct.Commit()

	if ret3.Error != nil {
		ct.Rollback()   // 回滚
	}



}
