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


	// 增加
	//tag1 := relate_tables.Tag{
	//	Name:"标签1",
	//	Desc:"标签1描述",
	//}
	//tag2 := relate_tables.Tag{
	//	Name:"标签1",
	//	Desc:"标签1描述",
	//}
	//db.Create(&tag1)
	//db.Create(&tag2)

	// 先从数据库查询数据，再绑定标签
	//var tag1 relate_tables.Tag
	//db.First(&tag1,9)
	//
 	//article := relate_tables.Article2{
	//	Title:"文章标题",
	//	Content:"文章内容",
	//	Desc:"文章描述",
	//	Tags:[]relate_tables.Tag{
	//		tag1,
	//		//tag2,
	//	},
	//}
	//
	//db.Create(&article)


	// 查询
	var article2 relate_tables.Article2
	db.Preload("Tags").Find(&article2,5)
	fmt.Println(article2)


	var article3 relate_tables.Article2
	db.First(&article3,5)
	db.Model(&article3).Association("Tags").Find(&article3.Tags)
	fmt.Println(article3)

	// 更新
	//var article4 relate_tables.Article2
	//db.Preload("Tags").Find(&article4,5)
	//
	//db.Model(&article4.Tags).Where("id = ?",9).Update("name","gin标签")

	// 删除
	var article5 relate_tables.Article2
	db.Preload("Tags").Find(&article5,5)
	// 记得加删除条件
	db.Where("id = ?",9).Delete(&article5.Tags)

}
