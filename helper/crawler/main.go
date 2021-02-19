package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Crawler crawl data
func Crawler() {
	URL0 := "https://shopee.vn/"
	URL1 := "https://shopee.vn/api/v2/item/get?itemid=3163681208&shopid=277411443"
	client := http.Client{}
	res, err := client.Get(URL0)
	res, err = client.Get(URL1)
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

	fmt.Println(string(responseBodyBytes))
}
