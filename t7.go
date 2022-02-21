package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

func SingleHost(handler http.Handler, allowedHost string) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if host == allowedHost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(ourFunc)
}

func main() {
	singleHosted := SingleHost(http.HandlerFunc(index), "example.com")
	http.Handle("/", singleHosted)
	http.ListenAndServe(":8000", singleHosted)
}
