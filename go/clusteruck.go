package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func slowHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(250 * time.Millisecond)
	w.Header().Set("Server", "Go/net/http")
	fmt.Fprint(w, "hello, world!\n")
}

func fastHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello, world!\n")
}

func resizeHandler(w http.ResponseWriter, req *http.Request) {
	procs, err := strconv.Atoi(req.FormValue("workers"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Resize did not work.  Please use '?workers=<number>'")
		return
	}

	fmt.Printf("sizing to %d procs...\n", procs)
	fmt.Println("previous number of procs used:", runtime.GOMAXPROCS(procs))
	fmt.Fprintf(w, "Resized to %d procs", procs)
}

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("number of procs:", numCPU)
	fmt.Println(runtime.GOMAXPROCS(numCPU))

	http.HandleFunc("/", slowHandler)
	http.HandleFunc("/resize", resizeHandler)

	http.ListenAndServe(":8081", nil)
}
