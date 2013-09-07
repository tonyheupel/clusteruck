package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func SlowHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(250 * time.Millisecond)
	fmt.Fprint(w, "hello, world!\n")
}

func FastHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello, world!\n")
}

func ResizeHandler(w http.ResponseWriter, req *http.Request) {
	procs, err := strconv.Atoi(req.FormValue("workers"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Resize did not work.  Please use '?workers=<number>'")
		return
	}

	fmt.Printf("sizing to %s procs...\n", procs)
	fmt.Println("previous number of procs used:", runtime.GOMAXPROCS(procs))
	fmt.Fprintf(w, "Resized to %d procs", strconv.Itoa(procs))
}

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("number of procs:", numCPU)
	fmt.Println(runtime.GOMAXPROCS(numCPU))

	http.HandleFunc("/", SlowHandler)
	http.HandleFunc("/resize", ResizeHandler)

	http.ListenAndServe(":8081", nil)
}
