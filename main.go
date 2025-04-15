package main

import (
	"crawl-coffee-product/handler"
	"fmt"
)

func main() {
	result, err := handler.GetCoffeeProduct()
	if err != nil {
		fmt.Println("get coffee product failed, err:", err)
		return
	}
	output := fmt.Sprintf("result: %s", string(result))
	fmt.Println(output)
}
