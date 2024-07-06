# Weather App

A simple weather application built using Go and JavaScript that fetches and displays weather information including temperature, humidity, pressure, and more for a specified city. The application also dynamically changes the background based on the temperature.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [API Configuration](#api-configuration)
- [Demo](#demo)
- [Acknowledgements](#acknowledgements)

## Overview

This project utilizes the [OpenWeatherMap API](https://openweathermap.org/api) to retrieve weather data for different cities. The weather details are displayed on a simple web interface, which changes its background image dynamically based on the temperature.

## Features

- Retrieve current weather data for any city.
- Display temperature in Kelvin, Celsius, and Fahrenheit.
- Show additional weather details like humidity, pressure, weather description, and wind speed.
- Dynamic background changes based on temperature ranges.
- Responsive and visually appealing user interface.

## Project Structure

```
weather/
├── main.go               # Go server to handle API requests and serve static files
├── static/
│   ├── index.html        # HTML for the weather app
│   ├── script.js         # JavaScript for fetching weather data and updating the UI
│   ├── styles.css        # CSS for styling the app
│   └── src/              # Directory for background images
│       ├── cold.jpg
│       ├── cool.JPG
│       ├── freeze.jpg
│       ├── hot.jpg
│       ├── mild.jpg
│       ├── vhot.jpg
│       └── warm.jpg
└── apiConfig.json       
```

## Installation

To set up the project locally, follow these steps:

1. **Clone the repository**:
    ```bash
    git clone https://github.com/yourusername/weather-app.git
    cd weather-app
    ```

2. **Ensure you have Go installed**: Download and install Go from [golang.org](https://golang.org/).

3. **Install dependencies**: No external Go dependencies are required for this project.

4. **Obtain an API key from OpenWeatherMap**:
    - Sign up at [OpenWeatherMap](https://openweathermap.org/) and get your API key.

5. **Configure the API key**:
    - Open `apiConfig.json` and replace the placeholder with your actual API key:
    ```json
    {
        "OpenWeatherMapApiKey": "your_actual_api_key_here"
    }
    ```

6. **Run the server**:
    ```bash
    go run main.go
    ```

7. **Open the app in your browser**:
    - Navigate to `http://localhost:8082/`.

## Usage

1. **Enter the city name**: In the text box, input the name of the city you want to check the weather for.
2. **Submit the form**: Click on "Get Weather" to retrieve and display the weather data.
3. **View the details**: The app will show temperature, humidity, pressure, weather description, and wind speed.

## API Configuration

- **apiConfig.json**:
  ```json
  {
      "OpenWeatherMapApiKey": "your_actual_api_key_here"
  }
  ```

- Replace `"your_actual_api_key_here"` with your own API key from OpenWeatherMap.

## Demo
![default](https://github.com/poornima2330/Weather/assets/113591982/3335a616-11a2-4c1e-a525-92e3635c8d09)
![Madurai](https://github.com/poornima2330/Weather/assets/113591982/909396b6-909f-4847-bf3a-15e3ec3db940)
![London](https://github.com/poornima2330/Weather/assets/113591982/9e33477f-fe11-4ff2-a7de-70d24f2f535a)
![kuwait](https://github.com/poornima2330/Weather/assets/113591982/d3afc57a-56ed-450f-9822-5a3889096fac)


## Acknowledgements

- [OpenWeatherMap](https://openweathermap.org/) for providing the weather API.
- [Go Programming Language](https://golang.org/) for the server-side logic.
