package relate_tables


// 一对多

type User2 struct {
	Id int
	Name string
	Age int
	Addr string
	Articles []Article `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`
}

type Article struct {
	Id int
	Title string
	Content string
	Desc string
	// 外键
	UId int
}



