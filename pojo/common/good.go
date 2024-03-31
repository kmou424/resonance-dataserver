package common

type GoodBase struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Station   string `json:"station"`
	Stock     int    `json:"stock"`
	Type      string `json:"type"`
	BasePrice int    `json:"base_price"`
}
