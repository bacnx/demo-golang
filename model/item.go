package model

type Item struct {
	Id   int64  `gorm:"column:id;primary_key;type:serial" json:"id"`
	Name string `gorm:"column:name"                       json:"name"`
}

func (Item) TableName() string {
	return "items"
}
