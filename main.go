package main

import (
	"crawl-coffee-product/handler"
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	httpCh := make(chan []byte)
	httpCh2 := make(chan []byte)
	result, err := handler.GetCoffeeProduct()
	if err != nil {
		fmt.Println("get coffee product failed, err:", err)
		return
	}
	wg.Add(2)
	go func() {
		defer close(httpCh)
		defer wg.Done()
		value := <-httpCh
		fmt.Println("ts success")
		tsContent := fmt.Sprintf("export const coffeeLocaleProductList = %s", string(value))
		err := handler.CreateFileBufio([]byte(tsContent), "coffee-product.ts")
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer close(httpCh2)
		defer wg.Done()
		value := <-httpCh2
		fmt.Println("json success")
		err := handler.CreateFileBufio(value, "coffee-product.json")
		if err != nil {
			panic(err)
		}
	}()

	httpCh <- result
	httpCh2 <- result
	// output := fmt.Sprintf("result: %s", string(result))
	// fmt.Println(output)
	wg.Wait()
}
