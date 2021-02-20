package models

type Item struct {
	Itemid int64 `json:"itemid" gorm:"primarykey;column:ID"`
}

func (i Item) TableName() string {
	return "item"
}
