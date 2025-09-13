package main

import (
    "crypto/md5"
    "fmt"
    "os"
    "sync"
)

func main() {
    faili := []string{"test1.txt", "test2.txt"}
    var wg sync.WaitGroup
    for _, f := range faili {
        wg.Add(1)
        go func(file string) {
            defer wg.Done()
            soderjimoe, _ := os.ReadFile(file)
            fmt.Println(file, md5.Sum(soderjimoe))
        }(f)
    }
    wg.Wait()
}
