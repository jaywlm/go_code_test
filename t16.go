package main

import (
	"fmt"
)

type A struct {
	name string
	age  int
}
type B struct {
	t A
}
type C struct {
	A
}

func main() {
	t1 := &B{
		t: A{
			name: "zll",
			age:  28,
		},
	}
	t2 := &C{
		A{
			name: "zrl",
			age:  29,
		},
	}
	t3 := &C{
		A: A{
			name: "zrl",
			age:  29,
		},
	}
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
}
