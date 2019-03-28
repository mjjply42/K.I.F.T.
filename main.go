package main

//#include "test.h"
//#include "printNumber.c"
//#include <stdlib.h>
import "C"
import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	com "./commands"
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
		// Cstring := C.printNumber()
		//Cstring:= C.pocketsphinx_continuous("~/Downloads/request.wav")
		testString := "Send Email"
		var commands = []string{
			"Get me the weather",
			"Events near me",
			"Send Email",
			"Search dictionary for term",
		}
		flag := 0
		//if string is equal to command
		for i := 0; i < len(commands); i++ {
			if strings.Compare(testString, commands[i]) == 0 {
				log.Println(commands[i])
				flag = 1
				if (i == 0) {
					log.Println(com.GetWeather("fremont"))
					fmt.Fprintln(res, com.GetWeather("fremont"))
				} else if (i == 1) {
					log.Println(com.GetEvents("fremont"))
					fmt.Fprintln(res, com.GetEvents("fremont"))
				} else if (i == 2) {
					message := "HELLO This is from kift"
					who := "stsong42@gmail.com"
					value := com.SendEmail(message, who)
					log.Println(value)
					fmt.Fprintln(res, value)
				} else if (i == 3) {
					log.Println(com.SearchTerm("potato"))
					fmt.Fprintln(res, com.SearchTerm("potato"))
				}
			}
		}
		if (flag == 0) {
			log.Println("Command not found.")
			fmt.Fprintln(res, "Command not found. Please Try Again.")
		}	
		// GoString := C.GoString(Cstring)
		// fmt.Fprintln(res, GoString)
		// fmt.Fprintln(res, "HALLLO")
		// os.Remove("~/Download/require.wav")
	})
	fmt.Println("Server Running...")
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/plain"
		}

		res.Header().Add("Content Type", contentType)
		res.Write(data)
	}

	http.ServeFile(res, req, fmt.Sprintf("%s", "index.html"))
}
