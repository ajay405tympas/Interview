package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchUrl(ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{}

	

	for req := range ch {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, req, nil)
		if err != nil {
			fmt.Println("POST Request failed")
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Request to %s failed: %v\n", req, err)
			continue
		}
		fmt.Println("Response", resp)

		//fmt.Sprintf("<ResponseURL %s>", req.url)
		//fmt.Sprintf("<ResponseURL %s>", req.statuscode)
	}
}
func main() {
	var wg sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	inputUrl := []string{
		"https://example.com",
		"https://httpbin.org/status/404",
		"https://httpbin.org/delay/3",
		"http://this-domain-does-not-exist.invalid",
		"http://httpstat.us/200?sleep=3000",
	}
	urlChan := make(chan string, 3)
	for i := 0; i < len(inputUrl); i++ {
		wg.Add(1)
		urlChan <- inputUrl[i]
		go fetchUrl(urlChan, ctx, &wg)
	}
	close(urlChan)
	wg.Wait()
	fmt.Println("Hello")
}
