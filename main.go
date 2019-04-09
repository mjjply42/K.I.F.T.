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
	"encoding/json"
	"net/url"
	"strings"
	"os"
	b64 "encoding/base64"
	com "./commands"
)

//var serverDirectory string
var serverPort int
var serverHost string
var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))
var history string
var HistoryCounter int = 1
var SpotifyClientID = "f137e890f3734bd38fdcf1d980158139"
var SpotifyClientSecrt = "119bb13ee90043919bb53c48baf17fa9"
var AccessToken = ""

func main() {

	flag.StringVar(&serverHost, "host", "0.0.0.0", "Server Host")
	flag.IntVar(&serverPort, "port", 3000, "Server Port")
	//flag.StringVar(&serverDirectory, "dir", "index.html", "Server start file")
	flag.Parse()


	//define endpoints
	http.HandleFunc("/", handler)
	http.HandleFunc("/callme", commandhandler)
	http.HandleFunc("/oauth", oauthTokenHandler)
	// http.Handle

	//Listen and serve connections
	fmt.Println("Server Running...")
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), nil)
}

func commandhandler(res http.ResponseWriter, req *http.Request) {
	// Cstring := C.printNumber()
	//Cstring:= C.pocketsphinx_continuous("~/Downloads/request.wav")
	testString := "Play Music"
	var commands = []string{
		"Get me the weather",
		"Events near me",
		"Send Email",
		"Search dictionary for term",
		"Check History",
		"Play Music",
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
			} else if (i == 5) {
					log.Println(com.PlayMusic(AccessToken))
					fmt.Fprintln(res, com.PlayMusic(AccessToken))
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


func GetWeather(city string) string{
	Apikey := "a84567876c635d5929647ab1879c3122"
	//define url for get request
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&APPID=%s", city, Apikey)
	fmt.Printf("Performed Get on %s\n", url )
	//perform get request 
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//Grab the body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a place where the json data will be stored
	type WeatherData struct {
		Main struct {
			Humidity int     `json:"humidity"`
			Pressure int     `json:"pressure"`
			Temp     float64 `json:"temp"`
			TempMax  int     `json:"temp_max"`
			TempMin  float64 `json:"temp_min"`
		} `json:"main"`
		Weather    []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
			ID          int    `json:"id"`
			Main        string `json:"main"`
		} `json:"weather"`
		Wind struct {
			Deg   int     `json:"deg"`
			Speed float64 `json:"speed"`
		} `json:"wind"`
	}
	
	var m WeatherData
	err = json.Unmarshal(body, &m)
	return fmt.Sprintf("The temperature in %s is %.2f degrees farenheit and it is %s.\n", city, m.Main.Temp, m.Weather[0].Description)
}

func oauthTokenHandler(w http.ResponseWriter, r *http.Request) {

	//grab the code which is used to grab auth token
	code := r.URL.Query().Get("code")

	//Create form
	form := url.Values{}
	form.Add("code", code)
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", "http://localhost:3000/oauth")
	
	// Construct post request
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))
    if err != nil {
		log.Print(err)
        os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	encoded := fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", SpotifyClientID, SpotifyClientSecrt))))
	req.Header.Set("Authorization", encoded)
	
	//Using Token grab auth token
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        log.Print(err)
        os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	type AccessResp struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
	}

	var m AccessResp
	err = json.Unmarshal(body, &m)

	// Save accessToken for later use
	AccessToken = m.AccessToken
	fmt.Println(AccessToken)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}