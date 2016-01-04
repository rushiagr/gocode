package main

import (
    "sync"
    "sync/atomic"
    "fmt"
    "math/rand"
    "runtime"
    "time"
)

var mutex = &sync.Mutex{}

var state = make(map[int64]int64)

var ops int64 = 0

func writer() {
    for {
        mutex.Lock()
        state[rand.Int63()] = rand.Int63()
        mutex.Unlock()
        atomic.AddInt64(&ops, 1)
        runtime.Gosched()
    }
}

func reader() {
    var temp int64;
    for {
        mutex.Lock()
        temp = state[rand.Int63()]
        mutex.Unlock()
        atomic.AddInt64(&ops, 1)
        runtime.Gosched()
    }
    fmt.Println(temp) //This line never gets printed
}

func main() {
    go writer()
    for i:=0; i<10; i++ {
        go reader()
    }

    time.Sleep(time.Second)
    //fmt.Println(len(state))
    fmt.Println("Total operations:", atomic.LoadInt64(&ops))
}
