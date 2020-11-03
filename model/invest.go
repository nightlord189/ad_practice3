package model

type Invest struct {
	Stoimost float64 `json:"stoimost"`
	Period   string  `json:"period"`
	Product  string  `json:"product"`
	Iniciatr string  `json:"iniciatr"`
	ID       string  `json:"id"`
}
