package helper

import (
	"fmt"
	"log"
	"shopee-tracker/client"
)

// Crawler crawl data
func Crawler() {

	URL1 := "https://shopee.vn/api/v2/item/get?itemid=3163681208&shopid=277411443"
	client := client.NewShopeeClient()
	responseBodyBytes, err := client.Get(URL1)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(responseBodyBytes))
}
