package dto

// DataC response data
type DataC struct {
	Data Categories `json:"data"`
}

// Categories response categories
type Categories struct {
	Categories []Category `json:"categories"`
}

// CategoryDTO response category
type CategoryDTO struct {
	CategoryID int64  `json:"category_id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
}
