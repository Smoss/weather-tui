package tui

import (
	"strings"

	"github.com/smoss/weather-tui/models"
)

var sun string = `
        .
      \ | /
    '-.;;;.-'
   -==;;;;;==-
    .-';;;'-.
      / | \
jgs     '
`
var moon string = `
       _..._     
     .' .::::.    
    :  ::::::::
    :  ::::::::  
	    \. '::::::'  
      \-.::'' 
jgs
`
var cloud string = `
     .--.-.
    ( (    )__ 
   (_,  \ ) ,_)
     '------'


jgs
`

var rain string = `
     .--.-.
    ( (    )__ 
   (_,  \ ) ,_)
     '------'
     \  \  \
     \  \  \ \     
jgs+smoss
`

func getSymbol(weather models.WeatherBlock) string {
	if weather.PrecipitationChance.Value >= 50 {
		return rain
	} else if strings.Contains(weather.ShortDescription, "Cloudy") ||
		strings.Contains(weather.ShortDescription, "Rain") ||
		strings.Contains(weather.ShortDescription, "Snow") {
		return cloud
	} else if strings.Contains(weather.Name, "Night") || strings.Contains(weather.Name, "Tonight") || !weather.IsDaytime {
		return moon
	}
	return sun
}
