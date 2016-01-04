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
    state[12345]=71717171
    for i := 0; i<10000; i++ {
        mutex.Lock()
        state[rand.Int63()] = rand.Int63()
        mutex.Unlock()
        runtime.Gosched()
    }
}

func reader() {
    fmt.Println(state[12345])
    for i := 0; i<10000; i++ {
        mutex.Lock()
        mutex.Unlock()
        runtime.Gosched()
    }
}

func main() {
    for i:=0; i<10; i++ {
        go reader()
        go writer()
    }
    time.Sleep(time.Second)
    fmt.Println(len(state))
}
