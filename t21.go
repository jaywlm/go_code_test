package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		time.Sleep(time.Second)
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string, 2)
	go makeCakeAndSend(cs, 12)
	receiveCakeAndPack(cs)
}
