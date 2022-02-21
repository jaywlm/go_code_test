package main

import (
	"encoding/json"
	"fmt"
)

type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	m := make(map[string]interface{})
	p := &Persion{
		Name: "zll",
		Age:  29,
	}
	j, _ := json.Marshal(p)
	json.Unmarshal(j, &m)
	fmt.Printf("%+s\n", j)
	fmt.Printf("%+v\n", m)
}
