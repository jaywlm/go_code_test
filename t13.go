package main

import "fmt"

func main() {
	a := []byte{'a', 'b', 'c'}
	b := []byte("rust")
	s := "golang"
	a = append(a, s...)
	b = append(b, s...)
	fmt.Println(string(a))
	fmt.Println(string(b))
}
