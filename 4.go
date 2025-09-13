package main

import (
 "fmt"
 "net/http"
 "sync"
)

func main() {
 urls := []string{
  "http://google.com",
  "httptest.org/status/404",
  "httptest.org/status/500",
  "httptest.org/status/201",
  "httptest.org/status/403",
 }

 const maxWorkers = 3
 jobs := make(chan string, len(urls))
 results := make(chan string, len(urls))

 var wg sync.WaitGroup
 for i := 1; i <= maxWorkers; i++ {
  wg.Add(1)
  go func(workerID int) {
   defer wg.Done()
   for url := range jobs {
    resp, err := http.Get(url)
    if err != nil {
     results <- fmt.Sprintf("URL: %s, Ошибка: %v", url, err)
     continue
    }
    results <- fmt.Sprintf("URL: %s, Статус: %d", url, resp.StatusCode)
    resp.Body.Close()
   }
  }(i)
 }

 for _, url := range urls {
  jobs <- url
 }
 close(jobs)

 go func() {
  wg.Wait()
  close(results)
 }()

 for result := range results {
  fmt.Println(result)
 }
}
