package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_project/models/relate_tables"
	"fmt"
)

// 一对多操作
func main() {

	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8&parseTime=True&loc=Local
	db,err := gorm.Open("mysql","root:Qazwsx123@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()   // 关闭空闲的连接


	// 添加
	// 第一种方式
	//user2 := relate_tables.User2{
	//	Name:"hallen",
	//	Age:18,
	//	Addr:"长沙",
	//	Articles:[]relate_tables.Article{
	//		{
	//			Title:"beego详解",
	//			Content:"beego详解内容",
	//			Desc:"beego详解描述",
	//		},
	//		{
	//			Title:"beego详解2",
	//			Content:"beego详解内容2",
	//			Desc:"beego详解描述2",
	//		},
	//	},
	//}
	//db.Create(&user2)


	// 第二种添加方式：推荐使用第一种方式

	//article := relate_tables.Article{
	//	Title:"beego详解3",
	//	Content:"beego详解内容3",
	//	Desc:"beego详解描述3",
	//}
	//db.Create(&article)
	//article2 := relate_tables.Article{
	//	Title:"beego详解4",
	//	Content:"beego详解内容4",
	//	Desc:"beego详解描述4",
	//}
	//db.Create(&article2)


	//user2 := relate_tables.User2{
	//	Name:"hallen",
	//	Age:18,
	//	Addr:"长沙",
	//	Articles:[]relate_tables.Article{
	//		article,
	//		article2,
	//	},
	//}
	//
	//db.Create(&user2)



	// 查询

	// preload
	var user2 relate_tables.User2

	db.Preload("Articles").Find(&user2,1)
	fmt.Println(user2)

	// 第二种方式
	var user3 relate_tables.User2
	db.First(&user3,1)
	db.Model(&user2).Association("Articles").Find(&user3.Articles)
	fmt.Println(user3)

	// 第三种方式
	var user4 relate_tables.User2
	db.First(&user4,1)

	var articles []relate_tables.Article
	db.Model(&user4).Related(&articles,"Articles")

	fmt.Println(user4)
	fmt.Println(articles)


	// 更新操作
	var user5 relate_tables.User2
	db.Preload("Articles").Find(&user5,1)
	fmt.Println("========更新操作")
	fmt.Println(user5)
	// 更新操作需要加条件
	db.Model(&user5.Articles).Where("id = ?",1).Update("title","beego详解1")

	// 删除:一定要加限制条件
	var user6 relate_tables.User2
	db.Preload("Articles").Find(&user6,1)

	db.Where("id = ?",1).Delete(&user6.Articles)







}
