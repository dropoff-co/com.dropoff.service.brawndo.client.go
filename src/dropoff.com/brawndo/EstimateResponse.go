package brawndo

type EstimateLocation struct {
	Lat				float64 `json:"lat"`
	Lng				float64 `json:"lng"`
	LocationType	string  `json:"type"`
}

type EstimateCoordinates struct {
	Destination		*EstimateLocation			`json:"destination"`
	Pickup			*EstimateLocation			`json:"pickup"`
}

type EstimateServiceType struct {
	ETA 		string			`json:"ETA"`
	Distance 	string			`json:"Distance"`
	Price 		string			`json:"Price"`
}

type EstimateData struct {
	ETA 		string			`json:"ETA"`
	Distance 	string			`json:"Distance"`
	ServiceType string			`json:"service_type"`
	Asap  		*EstimateServiceType	`json:"asap"`
	TwoHr  		*EstimateServiceType	`json:"two_hr"`
	FourHr 		*EstimateServiceType	`json:"four_hr"`
	AllDay 		*EstimateServiceType	`json:"all_day"`
	Coordinates *EstimateCoordinates	`json:"coordinates"`
}

type EstimateResponse struct {
	Data      	*EstimateData		`json:"data"`
	Success   	bool			`json:"success"`
	Timestamp 	string			`json:"timestamp"`
}

