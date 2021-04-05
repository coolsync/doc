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
	//db.Raw("select * from users").Find(&users)
	db.Raw("select * from users where name = ?","hallen9").Find(&users)
	fmt.Println(users)

	// 增加
	//db.Exec("insert into users(name,age) values(?,?)","hallen111",18)

	// 修改
	//db.Exec("update users set name = ? where id = ?","hallen222",8)

	// 删除
	db.Exec("delete from users where id = ?",8)


}
