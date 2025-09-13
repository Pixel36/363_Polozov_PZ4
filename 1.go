package main

import (
 "fmt"
 "sync"
 "time"
)

func main() {
 var wg sync.WaitGroup
 wg.Add(1)

 go func() {
  defer wg.Done()
  for i := 1; i <= 5; i++ {
   fmt.Printf("%d\n", i)
   time.Sleep(1 * time.Second)
  }
 }()

 wg.Wait()
 fmt.Println("Горутина завершена")
}
