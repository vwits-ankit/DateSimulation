package data

import "time"

type FrequencyofShippingRoutes struct {
	Frequency            float64
	TransitTime          float64
	ChangeVersion        float64
	Username             string
	ProductOfDepartureId string
	PortOfArrivalID      string
	LastChange           time.Time
}
