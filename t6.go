package main

import (
	"fmt"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	w.(http.Flusher).Flush()
	time.Sleep(time.Second)
	fmt.Fprintln(w, "hello world")
	w.(http.Flusher).Flush()
}
func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world2")
	w.(http.Flusher).Flush()
	time.Sleep(time.Second)
	fmt.Fprintln(w, "hello world2")
	fmt.Fprintln(w, r.URL.Path, r.URL.RawQuery)
	fmt.Fprintln(w, r.URL)
	w.(http.Flusher).Flush()
}

func main() {
	http.HandleFunc("/a/", IndexHandler)
	http.HandleFunc("/aa/", IndexHandler2)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
