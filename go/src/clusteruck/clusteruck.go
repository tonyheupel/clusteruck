package main

import (
	"net/http"
	"io"
	"runtime"
	"fmt"
	"time"
	"strconv"
)


func SlowHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(250 * time.Millisecond);
	io.WriteString(w, "hello, world!\n")
}


func FastHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}


func ResizeHandler(w http.ResponseWriter, req *http.Request) {
	procs, err := strconv.Atoi(req.FormValue("workers"))

	if (err != nil) {
		io.WriteString(w, "Resize did not work.  Please use '?workers=<number>'")
		return
	}

	fmt.Println("sizing to", procs, "procs...")
	fmt.Println(runtime.GOMAXPROCS(procs))
	io.WriteString(w, "Resized to " + strconv.Itoa(procs) + " procs")
}


func main() {
	numCPU := runtime.NumCPU();
	fmt.Println("number of procs:", numCPU)
	fmt.Println(runtime.GOMAXPROCS(numCPU))

	http.HandleFunc("/", SlowHandler)
	http.HandleFunc("/resize", ResizeHandler)

	http.ListenAndServe(":8081", nil)
}
