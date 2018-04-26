package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func cleanup() {
    if r := recover(); r != nil {
        fmt.Println("Recovered in cleanup: ", r)
    }
}

func say(s string) {
    defer wg.Done()
    defer cleanup()
    for i := 0; i < 3; i++ {
        time.Sleep(time.Millisecond*100)
        fmt.Println(s)
        if i == 2 {
          panic("Oh dear, it's a 2!")
        }
    }
}

func main() {
    wg.Add(1)
    go say("Hey")
    wg.Add(1)
    go say("There")
    wg.Wait()
}
