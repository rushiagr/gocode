package main

import (
    "fmt"
    "math/rand"
    "time"
)

type readOp struct {
    key int64
    resp chan int64
}

type writeOp struct {
    key int64
    value int64
    resp chan bool
}

var state = make(map[int64]int64)
var writesChan = make(chan *writeOp)
var readsChan = make(chan *readOp)

func stateManager() {
    for {
        select{
        case read := <-readsChan:
            read.resp  <- state[read.key]
        case write := <-writesChan:
            state[write.key] = write.value
            write.resp <- true
        }
    }
}

func writer() {
    for {
        write := &writeOp{
            key: rand.Int63(),
            value: rand.Int63(),
            resp: make(chan bool),
        }
        writesChan <- write
        <-write.resp
    }
}

func reader() {
    for {
        read := &readOp{
            key: rand.Int63(),
            resp: make(chan int64),
        }
        readsChan <- read
        <-read.resp
    }
}

func main() {
    go stateManager()
    for i:=0; i<10; i++ {
        go writer()
        go reader()
    }

    time.Sleep(time.Second)
    fmt.Println(len(state))
}
