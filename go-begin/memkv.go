package main

import (
    "fmt"
    "math/rand"
    "time"
)

func creator(c chan int, thread int) {
    for {
        randNo := rand.Intn(1000000)
        time.Sleep(time.Duration(rand.Intn(100))  * time.Millisecond)
        c<-randNo
        fmt.Println("Sent", randNo, "from thread", thread)
    }
    return
}

func receiver(c chan int) {
    for {
        fmt.Println("Received",  <-c)
    }
    return
}

func main() {
    c := make(chan int)
    for i:=0; i<10; i++ {
        go creator(c, i+1)
    }
    go receiver(c)
    time.Sleep(100000*time.Millisecond)
}
