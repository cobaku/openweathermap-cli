package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"encoding/json"
)

type Weather struct {
	Main WeatherMain `json:"main"`
	Name string      `json:"name"`
}

type WeatherMain struct {
	Temp int `json:"temp"`
}

func (w *Weather) TempString() string {
	return w.Name + " " + strconv.Itoa(w.Main.Temp) + " °C"
}

var URL = "http://api.openweathermap.org/data/2.5/weather?q="

func main() {
	owmKey, success := os.LookupEnv("OWM_KEY")

	if !success {
		panic("Export OWM_KEY into your env and try again")
	}

	owmCity, success := os.LookupEnv("OWM_CITY")

	if !success {
		panic("Export City,country as OWM_CITY into env. Ex: Obninsk,ru")
	}

	resp, err := http.Get(URL + owmCity + "&units=metric&APPID=" + owmKey)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	var currentWeather Weather
	json.Unmarshal(data, &currentWeather)

	fmt.Print(currentWeather.TempString())
}
