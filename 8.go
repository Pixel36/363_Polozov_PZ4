package main

import "fmt"

func main() {
    workers := []string{"Матвей", "Серёга", "Карина"}
    requests := make(chan string)
    go func() {
        for i := 0; ; i++ {
            requests <- workers[i%len(workers)]
        }
    }()
		for i := 1; i < 43; i++ {
        fmt.Println("Запрос обработал", <-requests)
    }
}
