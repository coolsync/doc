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



	// First
    // 1
	var user relate_tables.User
	db.First(&user)  // 默认使用id查询
	fmt.Println(user)
	// 2
	var user2 relate_tables.User
	db.First(&user2,"name = ?","hallen4")
	fmt.Println(user2)
	// 3
	var user3 relate_tables.User
	ret := db.Where("name = ?","hallen4").First(&user3)
	fmt.Println(user3)
	fmt.Println(ret.RowsAffected)   // 收影响的行数
	fmt.Println(ret.Error)	// 错误信息
	fmt.Println(user.Id) // 返回的主键

	// FirstOrCreate

	var user5 relate_tables.User
	user4 := relate_tables.User{
		Name:"hallen6",
		Age:19,
		Addr:"湖南长沙芙蓉区",

	}

	ret1 := db.FirstOrCreate(&user5,user4)
	fmt.Println(user5)

	fmt.Println(ret1.RowsAffected)


	// Last
	var user6 relate_tables.User
	db.Last(&user6)
	fmt.Println(user6)

	// Take

	var user7 relate_tables.User
	db.Take(&user7,3)
	fmt.Println(user7)

	// Find

	var user8 relate_tables.User

	id_arr := []int{1,23}
	db.Debug().Find(&user8,id_arr)
	fmt.Println(user8)

}
