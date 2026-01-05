package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	//"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
)

type myJson struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`

	Visibility int    `json:"visibility"`
	Name       string `json:"name"`
	Timezone   int    `json:"timezone"`
}

// Go Weather App
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	BASE_URL := os.Getenv("BASE_URL")
	API_KEY := os.Getenv("API_KEY")

	fmt.Println("Where do you want to check the weather:")
	var city string
	fmt.Scanln(&city)

	SEARCH_URL := fmt.Sprintf("%s?q=%s&APPID=%s", BASE_URL, city, API_KEY)

	resp, err := http.Get(SEARCH_URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		weatherBytes, _ := io.ReadAll(resp.Body)
		weather := myJson{}
		json.Unmarshal(weatherBytes, &weather)

		fmt.Printf("Today in %v the weather is %v. The Temperature is %2.2v °C and the humidity is %v.", city, weather.Weather[0].Description, weather.Main.Temp-273.15, weather.Main.Humidity)
	}
}
