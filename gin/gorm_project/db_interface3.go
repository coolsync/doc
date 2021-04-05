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

	// 所有的都打印sql信息
	db.LogMode(true)


	// FirstOrInit

	var user relate_tables.User

	//db.FirstOrInit(&user,relate_tables.User{Name:"hallen9"})   // 类似于下面的语句
	db.Debug().FirstOrInit(&user,relate_tables.User{Name:"hallen5",Age:18})
	//db.Where("name = ?","hallen9").Find(&user)
	fmt.Println(user)


	// attrs
	var user2 relate_tables.User

	db.Where("name = ?","hallen9").Attrs(relate_tables.User{Name:"hallen10",Age:19,Addr:"xxx"}).FirstOrInit(&user2)
	fmt.Println(user2)

	// Assign

	var user3 relate_tables.User

	db.Debug().Where("name = ?","hallen9").Assign(relate_tables.User{Name:"hallen10"}).FirstOrInit(&user3)
	fmt.Println(user3)

	// pluck

	var ages []int

	var users []relate_tables.User
	db.Where("age < ?",20).Find(&users).Pluck("age",&ages)
	fmt.Println(ages)

	// Scopes

	// 无参数的
	var users_ok []relate_tables.User
	var users_not_ok []relate_tables.User
	db.Scopes(GetStatusOk).Find(&users_ok)
	db.Scopes(GetStatusNotOk).Find(&users_not_ok)
	fmt.Println(users_ok)
	fmt.Println(users_not_ok)

	// 有参数的
	var users2 []relate_tables.User

	//db.Scopes(GetRowsByName("hallen5")).Find(&users2)
	db.Scopes(GetRowsByName("hallen9")).Find(&users2)
	fmt.Println(users2)


}

func GetStatusOk(db *gorm.DB) *gorm.DB {  // 状态为1
	//return db.Where("status = ?",1)   // 正常的
	return db.Where("p_id = ?",1)  // 为了演示使用
}



func GetStatusNotOk(db *gorm.DB) *gorm.DB {  // 状态为1
	//return db.Where("status = ?",0)  // 正常的
	return db.Where("p_id = ?",0)  // 为了演示使用
}

func GetRowsByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?",name)
	}
}

