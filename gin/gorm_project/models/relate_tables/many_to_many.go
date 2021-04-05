package relate_tables

type Article2 struct {
	Id int
	Title string
	Content string
	Desc string
	Tags []Tag `gorm:"many2many:Article2s2Tags"`  // ;ForeignKey:AId;AssociationForeignKey:TId

}

type Tag struct {
	Id int
	Name string
	Desc string
}


