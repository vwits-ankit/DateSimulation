package data

import "time"

type F000PlanAndF480_ETACurrent struct {
	ProductBrand  string
	Active        bool
	Plant         string
	ChangeVersion float64
	Model         string
	Username      string
	LastChange    time.Time
}
