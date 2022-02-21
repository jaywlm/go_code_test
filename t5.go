package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limit := rate.Every(1 * time.Second)
	limiter := rate.NewLimiter(limit, 0)
	prev := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d %v", i, time.Sub(time))
		prev = time.Now()
	}
}
