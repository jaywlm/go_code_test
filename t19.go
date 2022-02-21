package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(cs chan string, wg *sync.WaitGroup) {
	for s := range cs {
		time.Sleep(time.Second)
		fmt.Println("Packing received cake: ", s)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	cs := make(chan string, 2)
	go makeCakeAndSend(cs, 12)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go receiveCakeAndPack(cs, &wg)
	}

	wg.Wait()
}
