package models

type Machine struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Status   string  `json:"status"` // green|blue|red
	TimeLeft int     `json:"timeLeft"`
	Price    float64 `json:"price"`
}
