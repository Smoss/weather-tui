package models

type GridProps struct {
	GridX  int    `json:"gridX"`
	GridY  int    `json:"gridY"`
	GridId string `json:"gridId"`
}

type GridResults struct {
	Props GridProps `json:"properties"`
}
