package relation_tables

type User2 struct {
	Id      int
	Name    string
	Age     int
	Address string
	Article []Article `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`
}

type Article struct {
	Id      int
	Title   string
	Content string
	Desc    string
	UId     int
}
