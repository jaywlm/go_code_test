package main

import "fmt"

type Persion struct {
	Name  string
	Age   int
	Hobby map[string]string
}

func main() {
	t1 := &Persion{
		Name:  "zll",
		Age:   29,
		Hobby: make(map[string]string),
	}
	t3 := &Persion{
		Name: "zll",
		Age:  29,
	}
	t2 := &Persion{
		Name:  "zrj",
		Age:   30,
		Hobby: map[string]string{"ccc": "ddd", "eee": "fff"},
	}
	fmt.Printf("%+v\n", t1)
	fmt.Printf("%+v %+v\n", t3.Hobby, t1.Hobby)
	t1.Hobby["aaa"] = "bbb"
	fmt.Printf("%+v\n", t1)
	fmt.Printf("%+v\n", t2)
	fmt.Printf("%+v\n", t3)
	fmt.Printf("%+v %+v\n", t3.Hobby, t1.Hobby)
}
