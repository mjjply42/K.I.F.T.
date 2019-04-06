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
	"os"
	"log"
	"net/http"
	"strings"
	"encoding/json"
	com "./commands"
)

//var serverDirectory string
var serverPort int
var serverHost string
var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))
var history string
var HistoryCounter int = 1

func main() {

	flag.StringVar(&serverHost, "host", "0.0.0.0", "Server Host")
	flag.IntVar(&serverPort, "port", 3000, "Server Port")
	//flag.StringVar(&serverDirectory, "dir", "index.html", "Server start file")
	flag.Parse()


	//define endpoints
	http.HandleFunc("/", handler)
	http.HandleFunc("/callme", commandhandler)
	http.HandleFunc("/oauth", oauthHandler)
	// http.Handle

	//Listen and serve connections
	fmt.Println("Server Running...")
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), nil)
}

func commandhandler(res http.ResponseWriter, req *http.Request) {
	// Cstring := C.printNumber()
	//Cstring:= C.pocketsphinx_continuous("~/Downloads/request.wav")
	testString := "Check History"
	var commands = []string{
		"Get me the weather",
		"Events near me",
		"Send Email",
		"Search dictionary for term",
		"Check History",
	}
	flag := 0

	//add string from user to a history buffer
	history = history + fmt.Sprintf("%d. %s\n", HistoryCounter, testString)
	HistoryCounter += 1
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
			} else if (i == 4) {
				log.Println(history)
				fmt.Fprintln(res, history)
			}
		}
	}
	if (flag == 0) {
		log.Println("Command not found.")
		fmt.Fprintln(res, "Command not found. Please Try Again.")
	} 
}

func handler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))
	userIP := req.RemoteAddr
	log.Println(userIP);
	
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

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	// We will be using `httpClient` to make external HTTP requests later in our code
	httpClient := http.Client{}

	clientID := "f137e890f3734bd38fdcf1d980158139"
	clientSecret := "119bb13ee90043919bb53c48baf17fa9"
	redirect_uri2 := "http://localhost:3000/"

	// First, we need to get the value of the `code` query param
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")

	r.Header.Set("Authorization", "Basic " + clientSecret)

	// Next, lets for the HTTP request to call the github oauth enpoint
	// to get our access token
	reqURL := fmt.Sprintf("https://accounts.spotify.com/authorize?grant_type=authorization_code&client_id=%s&code=%s&redirect_uri=%s", clientID, code, redirect_uri2)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// We set this header since we want the response
	// as JSON
	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	log.Print(string(res.Body));

	// Parse the request body into the `OAuthAccessResponse` struct
	var t OAuthAccessResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// Finally, send a response to redirect the user to the "welcome" page
	// with the access token
	w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
	w.WriteHeader(http.StatusFound)
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}