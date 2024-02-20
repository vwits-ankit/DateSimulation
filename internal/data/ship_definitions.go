package data

import "time"

type ShipDefinitions struct {
	ChangeVersion float64
	ShipNo        string
	IMO           float64
	Username      string
	LastChange    time.Time
	NameOfShip    string
}
