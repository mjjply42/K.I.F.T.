package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func SetTimer(seconds time.Duration) string {
	time.Sleep(seconds * time.Second)
	return ("Ring Ring")
}

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
	return fmt.Sprintf("Weather: The temperature in %s is %.2f degrees farenheit and it is %s.\n", city, m.Main.Temp, m.Weather[0].Description)
}

func GetEvents(city string) string {
	//use api.eventful.com for this one
	url := fmt.Sprintf("http://api.eventful.com/json/events/search?...&keywords=books&location=%s&date=Future&within=10&app_key=nDJFtjJmt4pt4WjP", city)
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
				URL string `json:"url"`
			} `json:"event"`
		} `json:"events"`
	}
	var m EventInfo
	err = json.Unmarshal(body, &m)
	// maybe go through a few of the events on the list and concatenate to the string
	var response = "Event: Here are 3 events nearby.\n"
	var i = 0
	for i < 3 {
		response += m.Events.Event[i].URL + "\n"
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
		return ("Error during search")
	}
	//Grab the body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ("Error during search")
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

	var response = fmt.Sprintf("Define: A %s is %s\n", term, m.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0])
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
	return fmt.Sprintf("Mail: Message Sent")
}
