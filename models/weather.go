package models

import "time"

type TemperatureScale string

const (
	Farenheit TemperatureScale = "F"
	Celsius   TemperatureScale = "C"
)

type ElevationUnit string

const (
	Meters ElevationUnit = "wmoUnit:m"
	Feet   ElevationUnit = "wmoUnit:f"
)

type SubBlock struct {
	Value int64 `json:"value"`
}

type ElevationBlock struct {
	Value         float64       `json:"value"`
	ElevationUnit ElevationUnit `json:"unitCode"`
}

type LargeWeatherBlock struct {
	Name                string           `json:"name"`
	Temperature         int              `json:"temperature"`
	TemperatureUnit     TemperatureScale `json:"temperatureUnit"`
	StartTime           string           `json:"startTime"`
	EndTime             string           `json:"endTime"`
	PrecipitationChance SubBlock         `json:"probabilityOfPrecipitation"`
	Details             string           `json:"detailedForecast"`
	WindSpeed           string           `json:"windSpeed"`
	WindDir             string           `json:"windDirection"`
	ShortDescription    string           `json:"shortForecast"`
}

type WeatherBlock struct {
	Name                string           `json:"name"`
	Temperature         int              `json:"temperature"`
	TemperatureUnit     TemperatureScale `json:"temperatureUnit"`
	StartTime           time.Time        `json:"startTime"`
	EndTime             time.Time        `json:"endTime"`
	PrecipitationChance SubBlock         `json:"probabilityOfPrecipitation"`
	Details             string           `json:"detailedForecast"`
	WindSpeed           string           `json:"windSpeed"`
	WindDir             string           `json:"windDirection"`
	ShortDescription    string           `json:"shortForecast"`
	IsDaytime           bool             `json:"isDaytime"`
	RelativeHumidity    SubBlock         `json:"relativeHumidity"`
}

type WeatherProps struct {
	Periods        []WeatherBlock `json:"periods"`
	ElevationBlock ElevationBlock `json:"elevation"`
}

type WeatherResult struct {
	WeatherProps WeatherProps `json:"properties"`
}
