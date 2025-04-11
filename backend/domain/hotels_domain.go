package domain

type Hotel struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Location   Location `json:"location"`
	RoomsCount int      `json:"rooms_count"`
	Rating     float32  `json:"rating"`
}

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  int    `json:"number"`
	ZipCode int    `json:"zip_code"`
}
