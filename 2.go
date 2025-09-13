package main

import (
 "fmt"
 "sync"
)

func main() {
 const Workers = 3
 const Jobs = 10

 jobs := make(chan int, Jobs)
 results := make(chan int, Jobs)

 var wg sync.WaitGroup
 for i := 1; i <= Workers; i++ {
  wg.Add(1)
  go func(workerID int) {
   defer wg.Done()
   for job := range jobs {
    result := job * job
    fmt.Printf("Воркер %d: %d² = %d\n", workerID, job, result)
    results <- result
   }
  }(i)
 }

 for i := 1; i <= Jobs; i++ {
  jobs <- i
 }
 close(jobs)

 go func() {
  wg.Wait()
  close(results)
 }()

 var allResults []int
 for result := range results {
  allResults = append(allResults, result)
 }

 fmt.Printf("Результаты: %v\n", allResults)
}
