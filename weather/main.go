package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct for storing API configuration
type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

// Struct for storing weather data
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin   float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

// Initialize the API key
var apiConfig = apiConfigData{
	OpenWeatherMapApiKey: "753fe873ca756fa518055d62000c65a8",
}

// Handler for the /hello endpoint
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))
}

// Convert Kelvin to Celsius
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// Convert Kelvin to Fahrenheit
func kelvinToFahrenheit(k float64) float64 {
	return (k-273.15)*9/5 + 32
}

// Function to query weather data from the API
func query(city string) (weatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s", apiConfig.OpenWeatherMapApiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	// Convert temperature to Celsius and Fahrenheit
	d.Main.TempMin = kelvinToCelsius(d.Main.TempMin)
	d.Main.TempMax = kelvinToCelsius(d.Main.TempMax)

	return d, nil
}

// Function to enable CORS
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Main function
func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		// Extract city from query parameter
		queryValues := r.URL.Query()
		city := queryValues.Get("city")

		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		data, err := query(city)
		if err != nil {
			log.Printf("query error: %v", err)
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("json encoding error: %v", err)
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	// Serve static files from the 'static' directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
