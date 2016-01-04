package main

import (
    "fmt"
)

func main1() {
    var a [5]int
    fmt.Println("a: ", a)
    for i := 0; i<len(a); i++ {
        fmt.Println(a[i])
    }
}
