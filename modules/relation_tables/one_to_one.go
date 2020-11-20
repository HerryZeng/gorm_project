package relation_tables

type User struct {
	Id      int
	Name    string
	Age     int
	Address string
	PId     int
}

type UserProfile struct {
	Id    int
	Pic   string
	CPic  string
	Phone string
	User  User `gorm:"ForeignKey:PId;AssociationForeignKey:Id"` //关联关系
	//UserID int  // 默认关联字段
}
