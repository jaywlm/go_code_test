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
	//defer (*wg).Done()
	defer wg.Done()
	for s := range cs {
		time.Sleep(time.Second)
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	num := 4
	var wg sync.WaitGroup
	cs := make(chan string, 0)
	go makeCakeAndSend(cs, 200)
	for i := 0; i < num; i++ {
		wg.Add(1)
		go receiveCakeAndPack(cs, &wg)
	}

	wg.Wait()
}
