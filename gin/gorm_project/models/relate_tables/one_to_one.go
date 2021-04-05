package relate_tables



// 一对一

// 属于
//type User struct {
//	Id int
//	Name string
//	Age int
//	Addr string
//
//}
//
//
//type UserProfile struct {
//	Id int
//	Pic string
//	CPic string
//	Phone string
//	User User `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`  // 关联关系
//	//UserID int  // 默认关联字段
//	UId int // uid
//}


// 包含

type User struct {
	Id int
	Name string
	Age int
	Addr string
	PId int

}


type UserProfile struct {
	Id int
	Pic string
	CPic string
	Phone string
	User User `gorm:"ForeignKey:PId;AssociationForeignKey:Id"`  // 关联关系

}


// 属于：关系和外键在同一方，有关系的那一方属于另外一个模型
// 包含：关系和外键不在同一方，有关系的那一方包含另外一个有外键的模型