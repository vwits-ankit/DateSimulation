package data

type CommunicationETADates struct {
	ProductBrand                                  string  `json:"product_brand,omitempty"`
	ImporterID                                    string  `json:"importer_id,omitempty"`
	DeliveryCheckPoint                            string  `json:"delivery_check_point,omitempty"`
	Active                                        bool    `json:"active,omitempty"`
	CommunicationOnlyWhenChangesOfDeliveryCPOccur bool    `json:"communication_only_when_changes_of_delivery_cp_occur,omitempty"`
	ChangeVersion                                 float64 `json:"change_version,omitempty"`
}
