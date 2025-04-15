package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	url    = "https://shop.lebai.ltd/api/product/list"
	cookie = "eyJfcGVybWFuZW50Ijp0cnVlLCJ1c2VyX2lkIjoicWluZ2ZlbmcuemhhbmdAbGViYWkubHRkIn0.Z_3UVw.o5ajQ-QxbKQYAYyM7F4M-o6c_NE"
)

type GetCoffeeProductRequest struct {
	Dealer     string `json:"dealer"`
	Device     string `json:"device"`
	IsTemplate bool   `json:"is_template"`
	Keyword    string `json:"keyword"`
	Pn         int    `json:"pn"`
	Ps         int    `json:"ps"`
}

func GetCoffeeProduct() ([]byte, error) {
	var req = GetCoffeeProductRequest{
		Dealer:     "674e9e28a95a7d676b494b16",
		Device:     "",
		IsTemplate: false,
		Keyword:    "",
		Pn:         1,
		Ps:         20,
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return nil, err
	}
	var request *http.Request
	request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("new request failed, err:", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.AddCookie(&http.Cookie{
		Name:  "l-shop-token",
		Value: cookie,
	})

	var client = &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("http do failed, err:", err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed, err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}
