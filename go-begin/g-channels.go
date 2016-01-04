package main
import (
    "fmt"
    "time"
)

func sum2(a []int, c chan int) {
    time.Sleep(80 * time.Millisecond)
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum // send sum to c
}

func sum(a []int, c chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum // send sum to c
}

func channels() {
    a := []int{234, 2, 8, -2342349, 123423424, 23423520}

    c := make(chan int)
    go sum2(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}
