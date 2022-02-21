package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	num := runtime.NumCPU() * 2
	fmt.Printf("groutine num: %d\n", num)

	var wg sync.WaitGroup
	userCount := 200
	ch := make(chan bool, num)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		ch <- true
		go read(ch, i, &wg)
	}
	wg.Wait()
}

func read(ch chan bool, num int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Printf("go func %d\n", num)
	<-ch
	wg.Done()
}
