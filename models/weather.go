package models

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

type ProbPrecip struct {
	Value int64 `json:"value"`
}

type ElevationBlock struct {
	Value         float64       `json:"value"`
	ElevationUnit ElevationUnit `json:"unitCode"`
}

type WeatherBlock struct {
	Name            string           `json:"name"`
	Temperature     int              `json:"temperature"`
	TemperatureUnit TemperatureScale `json:"temperatureUnit"`
	StartTime       string           `json:"startTime"`
	EndTime         string           `json:"endTime"`
	// CloudCover      float64
	PrecipitationChance ProbPrecip `json:"probabilityOfPrecipitation"`
	Details             string     `json:"detailedForecast"`
	WindSpeed           string     `json:"windSpeed"`
	WindDir             string     `json:"windDirection"`
	ShortDescription    string     `json:"shortForecast"`
}

type WeatherProps struct {
	Periods        []WeatherBlock `json:"periods"`
	ElevationBlock ElevationBlock `json:"elevation"`
}

type WeatherResult struct {
	WeatherProps WeatherProps `json:"properties"`
}
