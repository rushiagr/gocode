package main

import (
    "sync"
    "fmt"
    "math/rand"
    "runtime"
    "time"
)

var mutex = &sync.Mutex{}

var state = make(map[int64]int64)

func writer() {
    for {
        mutex.Lock()
        state[rand.Int63()] = rand.Int63()
        mutex.Unlock()
        runtime.Gosched()
    }
}

func reader() {
    var temp int64;
    for {
        mutex.Lock()
        temp = state[rand.Int63()]
        mutex.Unlock()
        runtime.Gosched()
    }
    fmt.Println(temp) //This line never gets printed
}

func main() {
    for i:=0; i<10; i++ {
        go writer()
        go reader()
    }

    time.Sleep(time.Second)
    fmt.Println(len(state))
}
