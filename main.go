package main

import (
	"flag"
	"fmt"
	"net/http"
)

var serverHost string
var serverDirectory string
var serverPort int

func main() {

	flag.StringVar(&serverHost, "host", "0.0.0.0", "Server Host")
	flag.IntVar(&serverPort, "port", 3000, "Server Port")
	flag.StringVar(&serverDirectory, "dir", "index.html", "Server start file")
	flag.Parse()

	http.HandleFunc("/", handler)
	fmt.Println("Server Running...")
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, fmt.Sprintf("%s", serverDirectory))
}
