package table

type User struct {
	Id   int    `gorm:"column:user_id;type:int;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:varchar(20)"`
}

func (u *User) TableName() string {
	return "users"
}
