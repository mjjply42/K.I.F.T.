package main

import (
	"flag"
	"fmt"
	"net/http"
)

import C
var serverHost string
var serverPort int

func main() {

	flag.IntVar(&serverPort, "port", 3000, "Server Port")

	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
	fmt.Println("Server Running...")
}
func handler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
