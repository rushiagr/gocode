package main

import (
    "fmt"
)

func arr() {
    var a [5]int
    b := [5]int{1, 2, 3, 4, 5}
    c := [...]int{1, 2, 3, 4, 5, 6}
    d := [5]byte{'a', 'b', 'c', 'd', 'e'}
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}
