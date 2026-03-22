package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/smoss/weather-tui/models"
)

func GetZipcode(zipcode string) (*models.ZipcodeCoord, error) {

	if len(zipcode) != 5 {
		return nil, errors.New("Invalid Zipcode")
	}
	zipcodeUrl := fmt.Sprintf("https://api.zippopotam.us/us/%s", zipcode)
	resp, err := http.Get(zipcodeUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var zipResult models.ZippoResult
	if err = json.NewDecoder(resp.Body).Decode(&zipResult); err != nil {
		return nil, err
	}

	if len(zipResult.Places) == 0 {
		return nil, nil
	}
	return &zipResult.Places[0], nil
}
