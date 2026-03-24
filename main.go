package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/smoss/weather-tui/tui"
)

func main() {
	// result, err := api.GetZipcode("01803")
	// fmt.Println(result)
	// fmt.Println(err)

	// if err != nil {
	// 	panic("Didn't work")
	// }

	// weatherResult, err := api.GetWeather(*result)
	// fmt.Println(weatherResult)
	// fmt.Println(err)

	p := tea.NewProgram(tui.InitialState())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
