package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	sehir := "Antalya"
	apikey := "bc2171be2757ce326dea89093b83c9bc"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=tr", sehir, apikey)

	yanit, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer yanit.Body.Close()

	if yanit.StatusCode != 200 {
		fmt.Printf("API çağrısı başarısız, durum kodu: %d\n", yanit.StatusCode)
		return
	}

	body, err := io.ReadAll(yanit.Body)
	if err != nil {
		panic(err)
	}

	var data WeatherData
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	if len(data.Weather) == 0 {
		fmt.Println("Hava durumu bilgisi bulunamadı.")
		return
	}

	fmt.Printf("Şehir: %s\nSıcaklık: %.2f °C\nNem: %d%%\nHava Durumu: %s\n",
		data.Name, data.Main.Temp, data.Main.Humidity, data.Weather[0].Description)
}
