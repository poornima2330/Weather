document.getElementById('weatherForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const city = document.getElementById('city').value;
    fetchWeather(city);
});

async function fetchWeather(city) {
    const resultElement = document.getElementById('weatherResult');
    resultElement.textContent = 'Loading...';

    try {
        const response = await fetch(`/weather/?city=${city}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();

        if (!data || !data.name || !data.main) {
            throw new Error(`Invalid response format`);
        }

        const kelvinTemp = data.main.temp;
        const celsiusTemp = kelvinTemp - 273.15;
        const fahrenheitTemp = (kelvinTemp - 273.15) * 9 / 5 + 32;
        const humidity = data.main.humidity;
        const pressure = data.main.pressure;
        const weather = data.weather && data.weather.length > 0 ? data.weather[0].description : "N/A";
        const windSpeed = data.wind.speed;

        resultElement.innerHTML = `
            <p>The temperature in ${data.name} is ${kelvinTemp.toFixed(2)} K (${celsiusTemp.toFixed(2)} °C, ${fahrenheitTemp.toFixed(2)} °F).</p>
            <p>Humidity: ${humidity}%</p>
            <p>Pressure: ${pressure} hPa</p>
            <p>Weather: ${weather}</p>
            <p>Wind Speed: ${windSpeed} m/s</p>
        `;
        changeBackground(kelvinTemp);
    } catch (error) {
        console.error('Fetch error:', error);
        resultElement.textContent = `Failed to fetch weather data. Please try again later.`;
    }
}

function changeBackground(temp) {
    let imageUrl;
    const basePath = 'src/';

    if (temp <= 273) {
        imageUrl = `${basePath}freeze.jpg`;
    } else if (temp > 273 && temp <= 283) {
        imageUrl = `${basePath}cold.jpg`;
    } else if (temp > 283 && temp <= 293) {
        imageUrl = `${basePath}cool.JPG`;
    } else if (temp > 293 && temp <= 298) {
        imageUrl = `${basePath}mild.jpg`;
    } else if (temp > 298 && temp <= 308) {
        imageUrl = `${basePath}warm.jpg`;
    } else if (temp > 308 && temp <= 318) {
        imageUrl = `${basePath}hot.jpg`;
    } else {
        imageUrl = `${basePath}vhot.jpg`;
    }

    document.body.style.background = `linear-gradient(to right, #6a11cb, #2575fc), url('${imageUrl}')`;
    document.body.style.backgroundSize = 'cover';
    document.body.style.backgroundBlendMode = 'overlay';
}
