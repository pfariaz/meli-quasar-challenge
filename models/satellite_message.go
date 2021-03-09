package models

type SatelliteMessage struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
	Message  string  `json:"message"`
}
