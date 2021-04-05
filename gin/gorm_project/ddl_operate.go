package main

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_project/models/relate_tables"
	//"gorm_project/models"
)




func main() {

	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8&parseTime=True&loc=Local
	db,err := gorm.Open("mysql","root:Qazwsx123@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()   // 关闭空闲的连接



	// 创建表
	//db.CreateTable(&User{})  // 后面加了s
	// 指定表名
	//db.Table("user").CreateTable(&User{})

	// 删除表

	//db.DropTable("users")
	//db.DropTable(&User{})

	// 判断表存在不存在
	//b := db.HasTable("user")
	//fmt.Println(b)

	//b2 := db.HasTable(&User{})  // 判断users
	//fmt.Println(b2)



	// 统一加前缀或后缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "sys_" + defaultTableName + "sys_"
	//}

	// 自动迁移
	//db.AutoMigrate(&models.User{},&models.GormModel{},&models.Article{})
	db.AutoMigrate(&relate_tables.User{},&relate_tables.UserProfile{},&relate_tables.User2{},&relate_tables.Article{},&relate_tables.Article2{},&relate_tables.Tag{})





}
