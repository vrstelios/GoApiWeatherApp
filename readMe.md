## Go Weather App
Small Go CLI that tasks for a city, calls the OpenWeather API, and shows description, temperature, and humidity.

### Features
- City input from the terminal.
- Loads credentials from `.env` using `github.com/joho/godotenv`.
- Displays temperature in °C plus basic weather details.

### Requirements
- Go 1.24+
- OpenWeather account/API key.

### Setup
1) Clone:
```bash
git clone https://github.com/vrstelios/GoApiWeatherApp.git
cd GoApiWeatherApp
```
2) Environment:
The site [OpenWeatherMap](https://openweathermap.org/current) takes a correct API key.
Create a `.env` in the project root:
```bash
BASE_URL=https://api.openweathermap.org/data/2.5/weather
API_KEY=your_api_key_here
```
3) Dependencies (if needed):
```bash
go mod tidy
```

### Run
```bash
go run .
```
You will be prompted for a city, then the current weather data prints out.

### Example output
```
Where do you want to check the weather:
Thessaloniki
Today in thessaloniki the weather is scattered clouds. The Temperature is 16 °C and the humidity is 61.
```

### Notes
- Temperature is converted from Kelvin to °C.
- If the API responds with an error/invalid status, the program exits with a log.
