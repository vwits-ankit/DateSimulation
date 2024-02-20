package data

import "time"

type DistributionCalendar struct {
	Plant             string    `json:"plant,omitempty"`
	Model             string    `json:"model,omitempty"`
	Location          string    `json:"location,omitempty"`
	DealerNo          string    `json:"dealer_no,omitempty"`
	ProductBrand      string    `json:"product_brand,omitempty"`
	ModelGroup        string    `json:"model_group,omitempty"`
	Date              string    `json:"date,omitempty"`
	WildCard          string    `json:"wild_card,omitempty"`
	OperationTimeFrom string    `json:"operation_time_from,omitempty"`
	OperationTimeTo   string    `json:"operation_time_to,omitempty"`
	Duration          string    `json:"duration,omitempty"`
	Remark            string    `json:"remark,omitempty"`
	Username          string    `json:"username,omitempty"`
	LastChange        time.Time `json:"last_change,omitempty"`
	ChangeVersion     float64   `json:"change_version,omitempty"`
}
