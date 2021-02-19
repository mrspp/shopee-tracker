package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"shopee-tracker/models"
	"shopee-tracker/repository"
)

// Crawler crawl data
func Crawler() {
	URL := "https://shopee.vn/api/v4/official_shop/get_categories?tab_type=1"
	res, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close() // for garbage collection

	responseBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(string(responseBodyBytes))

	var resp models.Data
	err = json.Unmarshal(responseBodyBytes, &resp)
	if err != nil {
		log.Fatal(err)
	}

	cateRepo := repository.GetCategoryRepo()
	for _, c := range resp.Data.Categories {
		cateRepo.Save(c)
	}
	// fmt.Println(resp.Data.Categories)
}
