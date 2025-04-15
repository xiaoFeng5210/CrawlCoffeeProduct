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
	result, err := handler.GetCoffeeProduct()
	if err != nil {
		fmt.Println("get coffee product failed, err:", err)
		return
	}
	wg.Add(1)
	go func() {
		defer close(httpCh)
		defer wg.Done()
		value := <-httpCh

		tsContent := fmt.Sprintf("export const coffeeLocaleProductList = %s", string(value))
		err := handler.CreateFileBufio([]byte(tsContent))
		if err != nil {
			panic(err)
		}
	}()

	httpCh <- result
	// output := fmt.Sprintf("result: %s", string(result))
	// fmt.Println(output)
	wg.Wait()
}
