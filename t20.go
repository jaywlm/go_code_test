package main

import (
    "fmt"
    "strconv"
)

func makeCakeAndSend(cs chan string, count int) {
    for i := 1; i <= count; i++ {
        cakeName := "Strawberry Cake " + strconv.Itoa(i)
        cs <- cakeName //send a strawberry cake
    }
    close(cs)
}

func receiveCakeAndPack(cs chan string, flag chan int) {
    for s := range cs {
        fmt.Println("Packing received cake: ", s)
    }
    flag <- 1
}

func main() {
    flag := make(chan int)
    cs := make(chan string, 3)
    go makeCakeAndSend(cs, 12)
    go receiveCakeAndPack(cs, flag)
    <- flag
}
