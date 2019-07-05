package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func GetWeather(city string) string {
	Apikey := "a84567876c635d5929647ab1879c3122"
	//define url for get request
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&APPID=%s", city, Apikey)
	fmt.Printf("Performed Get on %s\n", url)
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
		Weather []struct {
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

func GetEvents(city string) string {
	//use api.eventful.com for this one
	url := fmt.Sprintf("http://api.eventful.com/json/events/search?...&keywords=fun&location=%s&date=Future&within=10&app_key=nDJFtjJmt4pt4WjP", city)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//Grab the body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type EventInfo struct {
		Events struct {
			Event []struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"event"`
		} `json:"events"`
	}
	var m EventInfo
	err = json.Unmarshal(body, &m)
	// maybe go through a few of the events on the list and concatenate to the string
	var response = "Here are 3 events nearby.\n"
	var i = 0
	for i < 3 {
		response += m.Events.Event[i].Title + "\n"
		response += "------------- \n"
		i++
	}
	return response
}

func SearchTerm(term string) string {
	//Using oxford dictionary api
	language := "en"

	client := &http.Client{}

	url := fmt.Sprintf("https://od-api.oxforddictionaries.com:443/api/v1/entries/%s/%s", language, strings.ToLower(term))
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("app_id", "3667cdeb")
	req.Header.Set("app_key", "0b5ead51722063464f81ec4b748eb952")
	resp, err := client.Do(req)

	if err != nil {
		return ("Error during search. Try again")
	}
	//Grab the body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ("Error during search. Try again")
	}

	type Definition struct {
		Metadata struct {
			Provider string `json:"provider"`
		} `json:"metadata"`
		Results []struct {
			ID             string `json:"id"`
			Language       string `json:"language"`
			LexicalEntries []struct {
				Entries []struct {
					Etymologies         []string `json:"etymologies"`
					GrammaticalFeatures []struct {
						Text string `json:"text"`
						Type string `json:"type"`
					} `json:"grammaticalFeatures"`
					HomographNumber string `json:"homographNumber"`
					Senses          []struct {
						Definitions []string `json:"definitions"`
						Domains     []string `json:"domains,omitempty"`
						Examples    []struct {
							Text string `json:"text"`
						} `json:"examples,omitempty"`
						ID               string   `json:"id"`
						ShortDefinitions []string `json:"short_definitions"`
						Notes            []struct {
							Text string `json:"text"`
							Type string `json:"type"`
						} `json:"notes,omitempty"`
						Regions   []string `json:"regions,omitempty"`
						Registers []string `json:"registers,omitempty"`
					} `json:"senses"`
				} `json:"entries"`
				Language        string `json:"language"`
				LexicalCategory string `json:"lexicalCategory"`
				Pronunciations  []struct {
					AudioFile        string   `json:"audioFile"`
					Dialects         []string `json:"dialects"`
					PhoneticNotation string   `json:"phoneticNotation"`
					PhoneticSpelling string   `json:"phoneticSpelling"`
				} `json:"pronunciations"`
				Text string `json:"text"`
			} `json:"lexicalEntries"`
			Type string `json:"type"`
			Word string `json:"word"`
		} `json:"results"`
	}

	var m Definition
	err = json.Unmarshal(body, &m)

	var response = fmt.Sprintf("A %s is %s\n", term, m.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0])
	return response
}

func SendEmail(message string, who string) string {

	username := "kiftkift42@gmail.com"
	passwd := "shotgun."

	cmd := exec.Command("python", "./PyCommands/command_email.py", message, who, username, passwd)
	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("Command finished with error: %v", err)
	}
	return fmt.Sprintf("Message Sent")
}

func PlayMusic(accessToken string) string {
	if accessToken == "" {
		return "Log in first. Repeat command."
	}

	//get available devices
	client := &http.Client{}
	req, nil := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/devices", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := client.Do(req)
	if err != nil {
		return ("Error during music play. Try again")
	}

	//Grab the body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ("Error during music play. Try again")
	}

	type connectedDevices struct {
		Devices []struct {
			ID               string `json:"id"`
			IsActive         bool   `json:"is_active"`
			IsPrivateSession bool   `json:"is_private_session"`
			IsRestricted     bool   `json:"is_restricted"`
			Name             string `json:"name"`
			Type             string `json:"type"`
			VolumePercent    int    `json:"volume_percent"`
		} `json:"devices"`
	}

	//parse json
	var m connectedDevices
	err = json.Unmarshal(body, &m)

	//if length of devices equal to zero
	if len(m.Devices) == 0 {
		return "Please open a device"
	}

	//construct url
	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/play?device_id=%s", m.Devices[0].ID)
	fmt.Println(url)
	if PlayMusicHelper(url, accessToken) == 0 {
		return "Error playing music"
	}
	return "Music now playing"
}

func PlayMusicHelper(url string, accessToken string) int {

	client := &http.Client{}

	//@PUT and play music with device id
	req, err := http.NewRequest("PUT", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	_, err = client.Do(req)
	if err != nil {
		return 0
	}
	return 1
}
