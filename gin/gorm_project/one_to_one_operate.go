package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_project/models/relate_tables"
	"fmt"
)

// 一对一操作
func main() {

	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8&parseTime=True&loc=Local
	db,err := gorm.Open("mysql","root:Qazwsx123@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()   // 关闭空闲的连接



	// 增加

	//user_profile := relate_tables.UserProfile{
	//	Pic:"1.jpg",
	//	CPic:"2.jpg",
	//	Phone:"13411223412",
	//	User:relate_tables.User{
	//		Name:"hallen",
	//		Age:18,
	//		Addr:"湖南长沙",
	//	},
	//}

	//db.Create(&user_profile)

	// 查询
	// 第一种方式

	var user_profile relate_tables.UserProfile

	db.Debug().First(&user_profile,1)

	db.Debug().Model(&user_profile).Association("User").Find(&user_profile.User)
	fmt.Println("++++++++++第一种查询方式")
	fmt.Println(user_profile)

	// 第二种
	var user_profile2 relate_tables.UserProfile
	db.Preload("User").Find(&user_profile2,2)
	fmt.Println("++++++++++第二种查询方式")
	fmt.Println(user_profile2)

	// 第三种
	var user_profile3 relate_tables.UserProfile
	db.First(&user_profile3,1)
	fmt.Println("++++++++++第三种查询方式")
	fmt.Println(user_profile3)

	var user relate_tables.User
	db.Model(&user_profile3).Related(&user,"User")

	fmt.Println(user)
	fmt.Println(user_profile3)


	// 更新：一定要加条件

	var user_profile4 relate_tables.UserProfile

	db.Preload("User").First(&user_profile4,2)
	fmt.Println("++++++++++更新操作")
	fmt.Println(user_profile4)
	// 更新单字段
	//db.Model(&user_profile4.User).Update("name","hallen3")
	db.Model(&user_profile4.User).Update(relate_tables.User{Name:"hallen4",Age:19,Addr:"湖南长沙芙蓉区"})


	// 删除操作
	var user_profile5 relate_tables.UserProfile
	db.Preload("User").First(&user_profile5,1)
	fmt.Println("+++++++++删除操作")
	fmt.Println(user_profile5)

	db.Debug().Delete(&user_profile5.User)








}
