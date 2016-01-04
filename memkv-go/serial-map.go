package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    state := make(map[int64]int64)
    go func() {
        for {
            state[rand.Int63()] = rand.Int63()
        }
    }()

   time.Sleep(time.Second)
   fmt.Println(len(state))
}
