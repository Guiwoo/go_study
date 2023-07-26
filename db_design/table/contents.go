package table

type Contents struct {
	Id   int    `gorm:"column:contents_id;type:int;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
}

func (c *Contents) TableName() string {
	return "contents"
}
