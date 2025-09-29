package Service

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func getResponse(url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		end := time.Since(start)
		ch <- fmt.Sprintf("(%s)-(%s) Error fetching URL: %s\n","404 BAD",end, err)
		return
	}
	end := time.Since(start)
	ch <- fmt.Sprintf("(%s)-(%s) Fetched URL: (%s)\n", response.Status, end, url)
	defer response.Body.Close()

}
func RequestParallel(urls []string)[]string {
	ch := make(chan string, len(urls))
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go getResponse(url, ch, &wg)

	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	res := []string{}
	for result := range ch {
		res = append(res, result)
	}
	return res
}
