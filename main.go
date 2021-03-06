package main

/*
#include <stdlib.h>
*/
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
	"os/exec"
	"strconv"
	"strings"

	com "./commands"
)

//var serverDirectory string
var serverPort int
var serverHost string
var tmpl = template.Must(template.New("tmpl").ParseFiles("./static/index.html"))
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
	//http.HandleFunc("/", handler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/command", commandhandler)
	http.HandleFunc("/response", responseHandler)
	http.HandleFunc("/oauth", oauthTokenHandler)
	http.HandleFunc("/users", usersHandler)

	//Listen and serve connections
	fmt.Println("Server Running...")
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), nil)
}
func runSphinx() string {
	out, err := exec.Command("./psphinx").Output()
	if err != nil {
		log.Printf("%s\n", err)
	}
	check := string(out)
	log.Printf("Command: %s", check)
	return check
}
func commandhandler(res http.ResponseWriter, req *http.Request) {

	testString := runSphinx()
	var commands = []string{
		"get me the weather\n",
		"events near me\n",
		"send email\n",
		"set alarm\n",
		"play music\n",
		//"who is connected\n",
		//"where is\n",
		"turn on light\n",
		"turn off light\n",
		"list commands\n",
	}

	//add string from user to a history buffer
	history = history + fmt.Sprintf("%d. %s\n", HistoryCounter, testString)
	HistoryCounter += 1
	flag := 0
	//if string is equal to command
	for i := 0; i < len(commands); i++ {
		if strings.Compare(testString, commands[i]) == 0 {
			log.Println(commands[i])
			flag = 1
			if i == 0 {
				log.Println("What city?")
				fmt.Fprintln(res, "weather; Please input the city")
			} else if i == 1 {
				flag = 1
				log.Println("What city?")
				fmt.Fprintln(res, "event; Please input the city")
			} else if i == 2 {
				flag = 1
				log.Println("Type Email Address")
				fmt.Fprintln(res, "email; Please type Email Address")
				/*} else if i == 3 {
				flag = 1
				log.Println("What word would you like to search?")
				fmt.Fprintln(res, "define; What word would you like to search?")*/
			} else if i == 3 {
				flag = 1
				log.Println("Alarm; Enter Alarm Time in minutes (MAX 59)")
				fmt.Fprintln(res, ("alarm; Enter Alarm Time in minutes (MAX 59)"))
			} else if i == 4 {
				flag = 1
				log.Println(com.PlayMusic(AccessToken))
				fmt.Fprintln(res, com.PlayMusic(AccessToken))
				/*} else if i == 5 {
				flag = 1
				log.Println(PrintConnected(UsersConnected))
				fmt.Fprintln(res, PrintConnected(UsersConnected))*/
			} else if i == 5 {
				flag = 1
				log.Println("Turn on lights")
				fmt.Fprintln(res, "lights; Turning on lights")
			} else if i == 6 {
				flag = 1
				log.Println("Turn off lights")
				fmt.Fprintln(res, "lights; Turning off lights")
			} else if i == 7 {
				flag = 1
				log.Println((commands))
				i := 0
				fmt.Fprintln(res, "commands; List Commands;")
				for i < len(commands) {
					fmt.Fprintln(res, (commands[i]))
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
func responseHandler(res http.ResponseWriter, req *http.Request) {

	sp := req.FormValue("value")
	speech, err := strconv.Atoi(sp)
	if err == nil {
		speech = speech
	}
	duty := req.FormValue("duty")
	var whom = "test@gmail.com"
	var answer = "a"
	if duty != "email" {
		answer = req.FormValue("value")
		answer = strings.TrimSuffix(answer, "\n")
	} else {
		whom = req.FormValue("email")
	}
	who := string([]byte(whom))

	if duty == "weather" {
		log.Println(com.GetWeather(answer))
		fmt.Fprintln(res, com.GetWeather(answer))
	} else if duty == "event" {
		log.Println(com.GetEvents(answer))
		fmt.Fprintln(res, com.GetEvents(answer))
	} else if duty == "email" {
		message := "HELLO This is from kift"
		value := com.SendEmail(message, who)
		log.Println(value)
		fmt.Fprintln(res, value)
	} else if duty == "define" {
		log.Println(com.SearchTerm(answer))
		fmt.Fprintln(res, com.SearchTerm(answer))
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
	var location = "Location; No user found"
	for i := 0; i < len(UsersConnected); i++ {
		if UsersConnected[i].Username == username {
			location = fmt.Sprintf("Location; Location is %s\n", UsersConnected[i].Location)
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
	return fmt.Sprintf("Who; %s\n", who)
}
