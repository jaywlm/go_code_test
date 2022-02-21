package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func readfile(file string) []string {
	li := []string{}
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		res := strings.Trim(string(buf), "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(" err = ", err)
		}
		li = append(li, res)
	}
	return li
}
func healthy(li []string) {
	for _, url := range li {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		fmt.Printf("%v %s  %s\n", time.Now().Unix(), url, resp.Status)
		fmt.Println(resp.Header)
	}
}
func tcron(li []string) {
	for {
		healthy(li)
		time.Sleep(time.Second * 2)
	}
}
func main() {
	li := readfile("tmp")
	go tcron(li)
	select {}
}
