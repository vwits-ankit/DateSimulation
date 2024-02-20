package data

import "time"

type CheckPoints struct {
	Type          string
	ETARelevant   string
	changeVersion float64
	CPNo          float64
	S             float64
	CheckPoint    string
	Username      string
	LastChange    time.Time
	Description   string
}
