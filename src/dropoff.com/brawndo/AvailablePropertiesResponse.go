package brawndo

type AvailablePropertiesData struct {
	Id					int64		`json:"id"`
	Label				string		`json:"label"`
	Description			string		`json:"description"`
	PriceAdjustment		float64		`json:"price_adjustment"`
	Requires			[]int64		`json:"requires"`
	Conflicts			[]int64		`json:"conflicts"`
}

type AvailablePropertiesResponse struct {
	Total		int64							`json:"total"`
	Count		int64							`json:"count"`
	LastKey		string							`json:"last_key"`
	Data      	[]*AvailablePropertiesData		`json:"data"`
	Success   	bool							`json:"success"`
	Timestamp 	string							`json:"timestamp"`
}