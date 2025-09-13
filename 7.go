package main

import "fmt"

func main() {
    set, get := make(chan string), make(chan string)
    go func() {
        var tmp string
        for {
            select {
            case tmp = <-set:
            case get <- tmp:
            }
        }
    }()
  	set <- "тест"
    fmt.Println(<-get)
}
