package types

type Point struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Author string `json:"author"`
}

