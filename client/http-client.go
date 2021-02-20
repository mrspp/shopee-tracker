package client

import "net/http"

// HTTPClient ...
type HTTPClient interface {
	Get(url string) ([]byte, error)
}

// NewShopeeClient ...
func NewShopeeClient() HTTPClient {
	return shopeeClient{
		http.Header{
			"Referer": []string{"Shopee.vn"},
		},
	}
}
