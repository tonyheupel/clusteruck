package main

import (
	"fmt"
	"net/http"
	"time"
)

func slowHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(250 * time.Millisecond)
	w.Header().Set("Server", "Go/net/http")
	fmt.Fprint(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", slowHandler)

	http.ListenAndServe(":8081", nil)
}
