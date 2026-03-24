package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smoss/weather-tui/models"
)

func getGridCoords(zipcode models.ZipcodeCoord) (*models.GridProps, error) {
	gridUrl := fmt.Sprintf("https://api.weather.gov/points/%s,%s", zipcode.Latitude, zipcode.Longitude)
	resp, err := http.Get(gridUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var gridResult models.GridResults
	if err = json.NewDecoder(resp.Body).Decode(&gridResult); err != nil {
		return nil, err
	}
	return &gridResult.Props, nil
}

func GetWeather(zipcode models.ZipcodeCoord) (*models.WeatherProps, error) {
	grid, err := getGridCoords(zipcode)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	weatherUrl := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast/hourly", grid.GridId, grid.GridX, grid.GridY)
	req, err := http.NewRequest("GET", weatherUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var weatherResult models.WeatherResult
	if err = json.NewDecoder(resp.Body).Decode(&weatherResult); err != nil {
		return nil, err
	}
	return &weatherResult.WeatherProps, nil
}
