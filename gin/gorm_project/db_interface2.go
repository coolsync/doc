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


	// where
	var user relate_tables.User

	db.Debug().Where("name = ?","hallen4").Find(&user)
	fmt.Println(user)

	// select

	var user2 relate_tables.User

	db.Select("name,age").Where("name = ?","hallen4").Find(&user2)
	fmt.Println(user2)


	// create

	//user3 := relate_tables.User{
	//	Name:"hallen8",
	//	Age:19,
	//	Addr:"xxx",
	//}
	//
	//db.Create(&user3)


	// 不支持批量插入

	//user4s := []relate_tables.User{
	//	{
	//		Name:"hallen9",
	//		Age:19,
	//		Addr:"xxx",
	//	},
	//	{
	//		Name:"hallen10",
	//		Age:19,
	//		Addr:"xxx",
	//	},
	//}
	//
	//db.Create(&user4s)

	// save
	//user4 := relate_tables.User{
	//	Name:"hallen9",
	//	Age:19,
	//	Addr:"xxx",
	//}
	//
	//db.Save(&user4)

	//var user5 relate_tables.User
	//db.Where("name = ?","hallen9").First(&user5)
	//fmt.Println(user5)
	//user5.Name = "hallen10"
	//db.Save(&user5)

	// update
	//var user6 relate_tables.User
	//db.Where("name = ?","hallen10").First(&user6)
	//fmt.Println(user6)
	//db.Model(&user6).Update("name","hallen12")

	//var user7 relate_tables.User
	//db.Where("name = ?","hallen12").Find(&user7).Update("name","hallen13")
	//db.Where("name = ?","hallen13").Find(&user7).Update(relate_tables.User{Name:"hallen14",Age:12})
	//db.Where("name = ?","hallen14").Find(&user7).Update(map[string]interface{}{"name":"hallen15","age":20})


	//var user8 relate_tables.User
	//
	//db.Where("name = ?","hallen8").Delete(&user8)


	// not
	var user9 []relate_tables.User
	db.Debug().Not("name = ?","hallen8").Find(&user9)
	fmt.Println(user9)

	// or
	var user10 []relate_tables.User

	db.Where("name = ?","hallen8").Or("name = ?","hallen7").Find(&user10)
	fmt.Println(user10)

	// order

	var user11 []relate_tables.User

	db.Debug().Where("name LIKE ?","%ha%").Order("name desc").Find(&user11)
	fmt.Println(user11)

	// limit offset
	var user12 []relate_tables.User
	db.Debug().Limit(2).Offset(2).Find(&user12)
	fmt.Println("=======limit and offset")
	fmt.Println(user12)

	// scan
	type UserBak struct {
		Name string
		//Age int
		Addr string
	}
	var user_bak UserBak
	var user13 []relate_tables.User
	db.Find(&user13).Scan(&user_bak)
	fmt.Println("======scan")
	fmt.Println(user13)
	fmt.Println(user_bak)

	// count
	var user14 []relate_tables.User
	var count int
	//db.Find(&user14).Count(&count)
	db.Model(&user14).Debug().Where("name like ?","%hal%").Count(&count)
	fmt.Println(count)

	// group  having
	var user15 []relate_tables.User
	type GroupData struct {
		Name string
		Age int
		Count int
	}

	var group_data []GroupData
	db.Select("name,age,count(*) as count").Group("name").Having("age = ?",18).Find(&user15).Scan(&group_data)
	fmt.Println(group_data)
	fmt.Println(user15)

	// join
	type UserProfile struct {
		Id int
		PId int
		Name string
		Age int
		Pic string
		CPic string
		Phone string
	}

	var user_profile []UserProfile
	var user16 []relate_tables.User

	db.Debug().Select("users.id,user_profiles.*").Joins("left join user_profiles on users.p_id = user_profiles.id").Find(&user16).Scan(&user_profile)
	fmt.Println("==========join")
	fmt.Println(user_profile)
	fmt.Println(user16)







}
