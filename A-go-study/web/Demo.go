package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func myHandler(writer http.ResponseWriter, request *http.Request) {
	// fmt.Printf("URL.Path = %q\n", request.URL.Path)
	// fmt.Println(writer)
	fmt.Fprintf(writer, "URL.Path = %q,method:%s,proto:%s\nurl:%s\nremotAddr:%s",
		request.URL.Path, request.Method, request.Proto, request.URL, request.RemoteAddr)
}
