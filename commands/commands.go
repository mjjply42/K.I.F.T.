package commands

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

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


func GetEvents(city string) string{
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
		Events     struct {
			Event []struct {
				URL           string      `json:"url"`
			} `json:"event"`
		} `json:"events"`
	}
	var m EventInfo
	err = json.Unmarshal(body, &m)
	// maybe go through a few of the events on the list and concatenate to the string
	var response = "Here are 3 events nearby.\n"
	var i = 0
	for i < 3 {
		response += m.Events.Event[i].URL + "\n"
		i++
	}
	return response
}