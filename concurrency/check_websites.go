package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
	time.Time
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool, len(urls))
	resultChannel := make(chan result)

	fmt.Println("Goroutine 1: ", runtime.NumGoroutine())
	//for _, url := range urls {
	//	go func(u string) {
	//		resultChannel <- result{u, wc(u)}
	//	}(url)
	//}

	go func() {
		for _, url := range urls {
			resultChannel <- result{url, wc(url), time.Now()}
		}
		close(resultChannel)
	}()

	fmt.Println("Goroutine 2: ", runtime.NumGoroutine())

	//for i := 0; i < len(urls); i++ {
	//	result := <-resultChannel
	//	results[result.string] = result.bool
	//}

	for v := range resultChannel {
		results[v.string] = v.bool
		fmt.Println("inside range: ", v.Time)
	}

	fmt.Println("Goroutine 3: ", runtime.NumGoroutine())


	return results
}