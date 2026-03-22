# weather-tui

A terminal weather forecast viewer for US ZIP codes, built with Go and [Bubble Tea](https://charm.land/bubbletea).

Type a 5-digit ZIP code and get the current National Weather Service forecast right in your terminal. Navigate through forecast periods with the arrow keys.

## Features

- Instant lookup -- forecast fetches automatically once you type 5 digits
- Browse forecast periods (today, tonight, tomorrow, etc.) with left/right arrows
- Displays temperature, precipitation chance, altitude, and time range
- No API keys required

## Prerequisites

- [Go](https://go.dev/) 1.25+

## Build & Install

```bash
# Run directly
go run .

# Or build a binary
go build -o weather-tui .
./weather-tui
```

## Usage

Launch the app and type a US ZIP code. The forecast loads automatically after the fifth digit.

```
What is your zipcode?

User Zip: 01803
Current Index: 1/14
Current Altitude: 47.000000 m
Times: 2026-03-22T13:00:00-04:00 - 2026-03-22T18:00:00-04:00
Current Temp: 52F
Risk of Rain: 10%

Press q to quit.
```

## Controls

| Key | Action |
|-----|--------|
| `0`-`9` | Enter ZIP code digits |
| `Enter` | Re-fetch weather for the current ZIP |
| `Left` / `Right` | Navigate forecast periods (wraps around) |
| `Backspace` | Delete last ZIP digit |
| `q` / `Ctrl+C` | Quit |

## How It Works

```
ZIP code ─→ Zippopotam.us ─→ lat/lon ─→ NWS /points ─→ grid coords ─→ NWS /forecast ─→ terminal
```

1. The ZIP code is sent to [Zippopotam.us](https://api.zippopotam.us) to resolve latitude and longitude.
2. Those coordinates hit the [National Weather Service API](https://api.weather.gov) `/points` endpoint to get a forecast grid location.
3. The grid location is used to fetch the full forecast from the NWS `/gridpoints/.../forecast` endpoint.
4. The forecast periods are displayed in the terminal.

## Project Structure

```
weather-tui/
├── main.go              # Entry point -- starts the Bubble Tea program
├── api/
│   ├── nws.go           # NWS API client (grid lookup + forecast)
│   └── zippo.go         # Zippopotam.us client (ZIP → coordinates)
├── models/
│   ├── grid.go          # NWS grid response types
│   ├── weather.go       # Forecast response types
│   └── zipResult.go     # ZIP lookup response types
├── tui/
│   └── state.go         # Bubble Tea model, update loop, and view
├── go.mod
└── go.sum
```

## APIs

| API | URL | Purpose |
|-----|-----|---------|
| National Weather Service | `https://api.weather.gov` | Grid coordinates and forecast data |
| Zippopotam.us | `https://api.zippopotam.us` | ZIP code to lat/lon geocoding |

Both APIs are free and require no authentication.
