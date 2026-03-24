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
	if strings.Contains(weather.ShortDescription, "Rain") {
		return rain
	} else if strings.Contains(weather.ShortDescription, "Cloudy") {
		return cloud
	} else if strings.Contains(weather.Name, "Night") || strings.Contains(weather.Name, "Tonight") {
		return moon
	}
	return sun
}
