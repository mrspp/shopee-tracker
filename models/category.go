package models

// Data response data
type Data struct {
	Data Categories `json:"data"`
}

// Categories response categories
type Categories struct {
	Categories []Category `json:"categories"`
}

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
