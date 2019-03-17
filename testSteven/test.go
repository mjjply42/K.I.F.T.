package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func getWeather(city string) string{
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
	//Parse the json body

	// Create a place where the json data will be stored
	type WeatherData struct {
		Base   string `json:"base"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Cod   int `json:"cod"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Dt   int `json:"dt"`
		ID   int `json:"id"`
		Main struct {
			Humidity int     `json:"humidity"`
			Pressure int     `json:"pressure"`
			Temp     float64 `json:"temp"`
			TempMax  int     `json:"temp_max"`
			TempMin  float64 `json:"temp_min"`
		} `json:"main"`
		Name string `json:"name"`
		Sys  struct {
			Country string  `json:"country"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Sunrise int     `json:"sunrise"`
			Sunset  int     `json:"sunset"`
			Type    int     `json:"type"`
		} `json:"sys"`
		Visibility int `json:"visibility"`
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

// func getWeather(city string) string{
// 	Apikey := "a84567876c635d5929647ab1879c3122"
// 	//define url for get request
// 	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&APPID=%s", city, Apikey)
// 	fmt.Printf("Performed Get on %s\n", url )
// 	//perform get request 
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	//Grab the body from response
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	//Parse the json body
// 	//	Create a place where the json data will be stored
// 	type WeatherData struct {
// 		Base   string `json:"base"`
// 		Clouds struct {
// 			All int `json:"all"`
// 		} `json:"clouds"`
// 		Cod   int `json:"cod"`
// 		Coord struct {
// 			Lat float64 `json:"lat"`
// 			Lon float64 `json:"lon"`
// 		} `json:"coord"`
// 		Dt   int `json:"dt"`
// 		ID   int `json:"id"`
// 		Main struct {
// 			Humidity int     `json:"humidity"`
// 			Pressure int     `json:"pressure"`
// 			Temp     float64 `json:"temp"`
// 			TempMax  int     `json:"temp_max"`
// 			TempMin  float64 `json:"temp_min"`
// 		} `json:"main"`
// 		Name string `json:"name"`
// 		Sys  struct {
// 			Country string  `json:"country"`
// 			ID      int     `json:"id"`
// 			Message float64 `json:"message"`
// 			Sunrise int     `json:"sunrise"`
// 			Sunset  int     `json:"sunset"`
// 			Type    int     `json:"type"`
// 		} `json:"sys"`
// 		Visibility int `json:"visibility"`
// 		Weather    []struct {
// 			Description string `json:"description"`
// 			Icon        string `json:"icon"`
// 			ID          int    `json:"id"`
// 			Main        string `json:"main"`
// 		} `json:"weather"`
// 		Wind struct {
// 			Deg   int     `json:"deg"`
// 			Speed float64 `json:"speed"`
// 		} `json:"wind"`
// 	}
	
// 	var m WeatherData
// 	err = json.Unmarshal(body, &m)
// 	return fmt.Sprintf("The temperature in %s is %.2f degrees farenheit and it is %s.\n", city, m.Main.Temp, m.Weather[0].Description)
// }

func getEvents(city string) string{
	//use api.eventful.com for this one
	url := fmt.Sprintf("http://api.eventful.com/json/events/search?&keywords=books&location=%s&date=Future&within=10&app_key=nDJFtjJmt4pt4WjP", city)
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
		LastItem   interface{} `json:"last_item"`
		TotalItems string      `json:"total_items"`
		FirstItem  interface{} `json:"first_item"`
		PageNumber string      `json:"page_number"`
		PageSize   string      `json:"page_size"`
		PageItems  interface{} `json:"page_items"`
		SearchTime string      `json:"search_time"`
		PageCount  string      `json:"page_count"`
		Events     struct {
			Event []struct {
				WatchingCount interface{} `json:"watching_count"`
				OlsonPath     string      `json:"olson_path"`
				CalendarCount interface{} `json:"calendar_count"`
				CommentCount  interface{} `json:"comment_count"`
				RegionAbbr    string      `json:"region_abbr"`
				PostalCode    string      `json:"postal_code"`
				GoingCount    interface{} `json:"going_count"`
				AllDay        string      `json:"all_day"`
				Latitude      string      `json:"latitude"`
				Groups        interface{} `json:"groups"`
				URL           string      `json:"url"`
				ID            string      `json:"id"`
				Privacy       string      `json:"privacy"`
				CityName      string      `json:"city_name"`
				LinkCount     interface{} `json:"link_count"`
				Longitude     string      `json:"longitude"`
				CountryName   string      `json:"country_name"`
				CountryAbbr   string      `json:"country_abbr"`
				RegionName    string      `json:"region_name"`
				StartTime     string      `json:"start_time"`
				TzID          interface{} `json:"tz_id"`
				Description   string      `json:"description"`
				Modified      string      `json:"modified"`
				VenueDisplay  string      `json:"venue_display"`
				TzCountry     interface{} `json:"tz_country"`
				Performers    interface{} `json:"performers"`
				Title         string      `json:"title"`
				VenueAddress  string      `json:"venue_address"`
				GeocodeType   string      `json:"geocode_type"`
				TzOlsonPath   string      `json:"tz_olson_path"`
				RecurString   interface{} `json:"recur_string"`
				Calendars     interface{} `json:"calendars"`
				Owner         string      `json:"owner"`
				Going         interface{} `json:"going"`
				CountryAbbr2  string      `json:"country_abbr2"`
				Image         struct {
					Small struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"small"`
					Width   string      `json:"width"`
					Caption interface{} `json:"caption"`
					Medium  struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"medium"`
					URL   string `json:"url"`
					Thumb struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"thumb"`
					Height string `json:"height"`
				} `json:"image"`
				Created   string      `json:"created"`
				VenueID   string      `json:"venue_id"`
				TzCity    interface{} `json:"tz_city"`
				StopTime  interface{} `json:"stop_time"`
				VenueName string      `json:"venue_name"`
				VenueURL  string      `json:"venue_url"`
			} `json:"event"`
		} `json:"events"`
	}
	var m EventInfo
	err = json.Unmarshal(body, &m)
	// maybe go through a few of the events on the list and concatenate to the string
	//
	return fmt.Sprintf("The temperature in %s is %.2f degrees farenheit and it is %s.\n", city, m.Main.Temp, m.Weather[0].Description)
}

func main() {
	// var ret = getWeather("fremont")
	var ret = getEvents("fremont")
	log.Println(ret)
}