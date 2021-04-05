package main

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_project/models"
)




func main() {

	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", "root:Qazwsx123@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close() // 关闭空闲的连接

	db.AutoMigrate(&models.User{})

	// 增
	//db.Create(&models.User{Name:"李四",Age:18,Addr:"xxx",Pic:"/static/upload/pic111.jpg",Phone:"13411232312"})

	// 查询

	//var user models.User

	//db.First(&user,1)  // 默认id
	//db.First(&user,"name=?","张三")   // 指定字段
	//fmt.Println(user)

	// 更新
	//var user models.User
	//db.First(&user,1)
	//db.Model(&user).Update("age",22)
	//db.Model(&user).Update("addr","zs-xxx")

	// 删除
	var user models.User
	db.First(&user,2)

	db.Delete(&user)





}
