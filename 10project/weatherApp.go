package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const apiKey = "e4c16c9f1124e4e4f37093e22dfced61" 

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to the weather API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("city not found or invalid request. HTTP status: %v", resp.Status)
	}

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %v", err)
	}

	return &weatherData, nil
}

func main() {
	fmt.Print("Enter the city name: ")
	var city string
	_, err := fmt.Scanln(&city)
	if err != nil || city == "" {
		log.Fatal("Invalid input. Please enter a valid city name.")
	}

	city = strings.TrimSpace(city)

	weatherData, err := getWeather(city)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v\n", err)
	}

	fmt.Printf("Weather in %s:\n", city)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp)
	fmt.Printf("Description: %s\n", weatherData.Weather[0].Description)
}
