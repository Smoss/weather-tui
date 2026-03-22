package models

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/smoss/weather-tui/api"
	"github.com/smoss/weather-tui/models"
)

type State struct {
	// APIKey  string
	Zipcode        string
	ZipcodeCoord   *models.ZipcodeCoord
	CurrentWeather *models.WeatherProps
	cursor         int
}

func InitialState() State {
	return State{
		Zipcode: "",
		cursor:  -1,
	}
}

func updateZipcode(character string, zipcode string) string {
	if len(zipcode) >= 5 {
		return zipcode
	}
	return zipcode + character
}

func (m State) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m *State) getWeather() {
	zipcodeCoord, err := api.GetZipcode(m.Zipcode)
	if err != nil {
		return
	}
	m.ZipcodeCoord = zipcodeCoord
	weather, err := api.GetWeather(*zipcodeCoord)
	if err != nil {
		return
	}
	m.CurrentWeather = weather
	m.cursor = 0
}

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "right" key moves to the next weather period
		case "right":
			if m.cursor < 0 {
				break
			}
			m.cursor++
			if m.CurrentWeather != nil && m.cursor >= len(m.CurrentWeather.Periods) {
				m.cursor = 0
			}
		// The "left" key moves to the previous weather period
		case "left":
			if m.cursor < 0 {
				break
			}
			m.cursor--
			if m.CurrentWeather != nil && m.cursor < 0 {
				m.cursor = len(m.CurrentWeather.Periods) - 1
			}

		// The "enter" key refreshes the weather
		case "enter":
			m.getWeather()

		// The "0-9" keys update the zipcode
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.Zipcode = updateZipcode(msg.String(), m.Zipcode)
			if len(m.Zipcode) == 5 {
				m.getWeather()
			}

		// The "backspace" key deletes the last character of the zipcode
		case "backspace":
			if len(m.Zipcode) > 0 {
				zipRunes := []rune(m.Zipcode)
				m.Zipcode = string(zipRunes[0 : len(zipRunes)-1])
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m State) View() tea.View {
	s := "What is your zipcode?\n\n"

	s += fmt.Sprintf("User Zip: %s\n", m.Zipcode)

	if m.CurrentWeather != nil {
		s += "\n"
		currPeriod := m.CurrentWeather.Periods[m.cursor]

		// Display the current index and the total periods in [ ] boxes
		periods := ""
		for i := 0; i < len(m.CurrentWeather.Periods); i++ {
			if i == m.cursor {
				periods += "[X] "
			} else {
				periods += "[ ] "
			}
		}
		s += periods + "\n"

		altitudeUnit := string([]rune(m.CurrentWeather.ElevationBlock.ElevationUnit)[len(m.CurrentWeather.ElevationBlock.ElevationUnit)-1])
		s += fmt.Sprintf("Current Altitude: %f %s\n", m.CurrentWeather.ElevationBlock.Value, altitudeUnit)
		s += fmt.Sprintf("Times: %s - %s\n", currPeriod.StartTime, currPeriod.EndTime)
		s += fmt.Sprintf("Current Temp: %d%s\n", currPeriod.Temperature, currPeriod.TemperatureUnit)
		s += fmt.Sprintf("Risk of Rain: %d%%\n", currPeriod.PrecipitationChance.Value)
		s += fmt.Sprintf("Details: %s\n", currPeriod.Details)
	}

	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
