package models

// Shopeemall ...
type Shopeemall struct {
	UserID           int64  `json:"userid"`
	Username         string `json:"username"`
	ShopID           int64  `json:"shopid" gorm:"primarykey;column:ID"`
	ShopName         string `json:"shop_name"`
	Logo             string `json:"logo"`
	LogoPc           string `json:"logo_pc"`
	ShopCollectionID int64  `json:"shop_collection_id"`
	Ctime            int64  `json:"ctime"`
	BrandLabel       int64  `json:"brand_label"`
}

// TableName name of table
func (sm Shopeemall) TableName() string {
	return "shopeemall"
}
