package main

//#include "test.h"
//#include "printNumber.c"
//#include <stdlib.h>
import "C"
import (
	b64 "encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

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
var UsersConnected []user_struct

func main() {

	flag.StringVar(&serverHost, "host", "0.0.0.0", "Server Host")
	flag.IntVar(&serverPort, "port", 3000, "Server Port")
	//flag.StringVar(&serverDirectory, "dir", "index.html", "Server start file")
	flag.Parse()

	//define endpoints
	http.HandleFunc("/", handler)
	http.HandleFunc("/callme", commandhandler)
	http.HandleFunc("/oauth", oauthTokenHandler)
	http.HandleFunc("/users", usersHandler)

	//Listen and serve connections
	fmt.Println("Server Running...")
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), nil)
}

func commandhandler(res http.ResponseWriter, req *http.Request) {
	// Cstring := C.printNumber()
	//Cstring:= C.pocketsphinx_continuous("~/Downloads/request.wav")
	testString := "Events near me"
	var commands = []string{
		"Get me the weather",
		"Events near me",
		"Send Email",
		"Search dictionary for",
		"Set Alarm",
		"Play Music",
		"Who is connected",
		"Where is",
		"Turn on light",
		"Turn off light",
		"List Commands",
	}

	//add string from user to a history buffer
	history = history + fmt.Sprintf("%d. %s\n", HistoryCounter, testString)
	HistoryCounter += 1
	flag := 0

	//if string is equal to command
	for i := 0; i < len(commands); i++ {
		if i == 3 {
			if len(testString) > 22 && strings.Compare(testString[0:21], commands[i]) == 0 {
				flag = 1
				log.Println(com.SearchTerm(testString[22:]))
				fmt.Fprintln(res, com.SearchTerm(testString[22:]))
			}
		} else if i == 7 {
			if len(testString) > 8 && strings.Compare(testString[0:8], commands[i]) == 0 {
				flag = 1
				//find username with rest of string
				log.Println(testString[9:])
				log.Println(findLocation(testString[9:]))
				fmt.Fprintln(res, findLocation(testString[9:]))
			}
		} else if strings.Compare(testString, commands[i]) == 0 {
			log.Println(commands[i])
			flag = 1
			if i == 0 {
				log.Println(com.GetWeather("Union City"))
				fmt.Fprintln(res, com.GetWeather("Union City"))
			} else if i == 1 {
				log.Println(com.GetEvents("fremont"))
				fmt.Fprintln(res, com.GetEvents("fremont"))
			} else if i == 2 {
				message := "HELLO This is from kift"
				who := "lauracmalagon@gmail.com"
				value := com.SendEmail(message, who)
				log.Println(value)
				fmt.Fprintln(res, value)
			} else if i == 3 {
				log.Println(com.SearchTerm("word"))
				fmt.Fprintln(res, com.SearchTerm("word"))
			} else if i == 4 {
				log.Println("Alarm: Enter Alarm Time in minutes (MAX 59)")
				fmt.Fprintln(res, ("Alarm: Enter Alarm Time in minutes (MAX 59)"))
			} else if i == 5 {
				log.Println(com.PlayMusic(AccessToken))
				fmt.Fprintln(res, com.PlayMusic(AccessToken))
			} else if i == 6 {
				log.Println(PrintConnected(UsersConnected))
				fmt.Fprintln(res, PrintConnected(UsersConnected))
			} else if i == 8 {
				log.Println("Turn on lights")
				fmt.Fprintln(res, "lights: Turning on lights")
			} else if i == 9 {
				log.Println("Turn off lights")
				fmt.Fprintln(res, "lights: Turning off lights")
			} else if i == 10 {
				log.Println((commands))
				i := 0
				for i < len(commands) {
					fmt.Fprintln(res, (commands[i]))
					fmt.Fprintf(res, "-------------")
					i++
				}
			}
		}
	}
	if flag == 0 {
		log.Println("Command not found.")
		fmt.Fprintln(res, "Command not found. Please Try Again.")
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))
	userIP := req.RemoteAddr
	log.Println(userIP)

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

type user_struct struct {
	Username string
	Location string
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		var t user_struct
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}
		log.Println("HERE ", t.Location)
		//ADD THE struct to UsersConnected ARRAY
		UsersConnected = append(UsersConnected, t)
		PrintSlice(UsersConnected)

	default:
		fmt.Fprintf(w, "Only Post request method is supported")
	}

	fmt.Fprintf(w, "Welcome to the users!")
	fmt.Println("Endpoint Hit: users")
}

func PrintSlice(s []user_struct) {
	fmt.Printf("len=%d cap=%d %+v\n", len(s), cap(s), s)
}

func findLocation(username string) string {
	var location = "Location: No user found"
	for i := 0; i < len(UsersConnected); i++ {
		if UsersConnected[i].Username == username {
			location = fmt.Sprintf("Location: Location is %s\n", UsersConnected[i].Location)
		}
	}
	return location
}
func PrintConnected(s []user_struct) string {
	if len(s) == 0 {
		return (fmt.Sprintf("%s\n", "Who: No Users Connected!"))
	}
	var who = fmt.Sprintf("%+v\n", s)
	who = strings.Replace(who, ":", "=>", -1)
	return fmt.Sprintf("Who: %s\n", who)
}
