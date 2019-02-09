package main

//#include "test.h"
//#include "printNumber.c"
//#include <stdlib.h>
import "C"
import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
)

//var serverDirectory string
var serverPort int
var serverHost string
var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))

func main() {

	flag.StringVar(&serverHost, "host", "0.0.0.0", "Server Host")
	flag.IntVar(&serverPort, "port", 3000, "Server Port")
	//flag.StringVar(&serverDirectory, "dir", "index.html", "Server start file")
	flag.Parse()
	http.HandleFunc("/", handler)
	http.HandleFunc("/callme", func(res http.ResponseWriter, req *http.Request) {
		Cstring := C.printNumber()
		//Cstring:= C.pocketsphinx_continuous("~/Downloads/request.wav")
		GoString := C.GoString(Cstring)
		fmt.Fprintln(res, GoString)
		//os.Remove("~/Download/require.wav")
		GoString = ""

	})
	fmt.Println("Server Running...")
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, fmt.Sprintf("%s", "index.html"))
}
