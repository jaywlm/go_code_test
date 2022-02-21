package main

import "fmt"

type stu struct {
	id   int64
	name string
}

type stuM struct {
	allst map[int64]stu
}

func (s stuM) del() {
	delete(s.allst, 200)
}

func main() {
	var s = stuM{
		allst: map[int64]stu{100: {100, "zll"}, 200: {200, "test"}},
	}
	fmt.Println(s)
	s.del()
	fmt.Println(s)
}
