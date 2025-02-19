```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int)

        wg.Add(1)
        go func() {
                defer wg.Done()
                fmt.Println("Goroutine 1 started")
                ch <- 1
                fmt.Println("Goroutine 1 exiting")
        }()

        wg.Wait()
        close(ch)
        fmt.Println("Main function exiting")
}
```