package client

import (
	"io/ioutil"
	"net/http"
)

type shopeeClient struct {
	Header http.Header
}

func (s shopeeClient) Get(url string) ([]byte, error) {
	client := http.Client{}
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, nil
	}
	rq.Header = s.Header
	response, err := client.Do(rq)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
