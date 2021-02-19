package models

// Category response category
type Category struct {
	CategoryID int64  `json:"category_id" gorm:"primarykey"`
	Name       string `json:"name"`
	Image      string `json:"image"`
}

// TableName name of table
func (c Category) TableName() string {
	return "category"
}
