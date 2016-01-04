//package main
//
//import (
//    "fmt"
//    "time"
//)
//func delayedPrint() {
//    time.Sleep(time.Millisecond * 1000)
//    fmt.Println("Delayed Print")
//}
//
//func parallel() {
//    for i := 0; i<10; i++ {
//        go delayedPrint()
//    }
//    time.Sleep(time.Millisecond * 4000)
//    fmt.Println("OK, done")
//}

package main

import (
    "fmt"
    "time"
)

func say2(s string){
    for i := 0; i < 5; i++ {
        time.Sleep(80 * time.Millisecond)
        fmt.Println(s)
    }
}

func say(s string){
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func parallel() {
    go say2("world")
    say("hello")
}

