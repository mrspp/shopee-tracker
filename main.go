package main

import (
	"shopee-tracker/config"
	helper "shopee-tracker/helper/crawler"
	"shopee-tracker/models"
)

func main() {
	config.GetDbConnection().AutoMigrate(&models.Category{}, models.Item{}, models.Shopeemall{})
	helper.Crawler()
}
