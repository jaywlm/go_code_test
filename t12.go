package main

import (
	"encoding/json"
	"fmt"
)

type Persion struct {
	Name string
	Age  int
}

func main() {
	var m map[string]interface{}
	p := &Persion{
		Name: "zll",
		Age:  29,
	}
	b, _ := json.Marshal(p)
	_ = json.Unmarshal(b, &m)
	fmt.Printf("%T %+v\n", b, string(b))
	fmt.Printf("%T %+v\n", m, m)
	for k, v := range m {
		fmt.Printf("%v %v %T\n", k, v, v)
	}

	mp := []map[string]interface{}{}
	t1 := map[string]interface{}{"name": "zll", "age": 29}
	t2 := map[string]interface{}{"name": "zrj", "age": 30}
	mp = append(mp, t1, t2)
	j, _ := json.Marshal(mp)
	fmt.Printf("%+v\n", string(j))

}
