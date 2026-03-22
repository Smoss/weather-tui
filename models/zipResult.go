package models

type ZipcodeCoord struct {
	PlaceName   string `json:"place name"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	State       string `json:"state"`
	StateAbbrev string `json:"state abbreviation"`
}

type ZippoResult struct {
	Country       string         `json:"country"`
	CountryAbbrev string         `json:"country abbreviation"`
	PostCode      string         `json:"post code"`
	Places        []ZipcodeCoord `json:"places"`
}
