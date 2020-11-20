package relation_tables

type User2 struct {
	Id      int
	Name    string
	Age     int
	Address string
}

type Article struct {
	Id      int
	Title   string
	Content string
	Desc    string
	User2   []User2 `gorm:"ForeignKey:Uid;AssociationForeignKey:Id"`
	Uid     int
}
