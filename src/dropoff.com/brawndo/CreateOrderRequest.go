package brawndo

type CreateOrderAddress struct {
	CompanyName  				string  `json:"company_name"`
	Email        				string  `json:"email"`
	Phone        				string  `json:"phone"`
	FirstName    				string  `json:"first_name"`
	LastName     				string  `json:"last_name"`
	AddressLine1 				string  `json:"address_line_1"`
	AddressLine2 				string  `json:"address_line_2,omitempty"`
	City         				string  `json:"city"`
	State        				string  `json:"state"`
	Zip          				string  `json:"zip"`
	Remarks      				string  `json:"remarks,omitempty"`
	Lat          				float64 `json:"lat"`
	Lng          				float64 `json:"lng"`
	EmailNotifications 			*bool 	`json:"email_notifications,omitempty"`
	SmsNotifications 			*bool 	`json:"sms_notifications,omitempty"`
}

type CreateOrderDetails struct {
	Quantity      int64  `json:"quantity"`
	Weight        int64  `json:"weight"`
	ETA           string `json:"eta"`
	Distance      string `json:"distance"`
	Price         string `json:"price"`
	ReadyDate     int64  `json:"ready_date"`
	Type          string `json:"type"`
	ReferenceCode string `json:"reference_code,omitempty"`
	ReferenceName string `json:"reference_name,omitempty"`
}

type CreateOrderItem struct {
	Container   int64  `json:"container,omitempty"`
	Description string `json:"description,omitempty"`
	Width       string `json:"width,omitempty"`
	Height      string `json:"height,omitempty"`
	Depth       string `json:"depth,omitempty"`
	PersonName  string `json:"person_name,omitempty"`
	Price       string `json:"price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Sku         string `json:"sku,omitempty"`
	Temperature int64  `json:"temperature,omitempty"`
	Weight      string `json:"weight,omitempty"`
	Unit        string `json:"unit,omitempty"`
}

type CreateOrderRequest struct {
	Details     *CreateOrderDetails `json:"details"`
	Origin      *CreateOrderAddress `json:"origin"`
	Destination *CreateOrderAddress `json:"destination"`
	Properties  []int64             `json:"properties,omitempty"`
	Items       []CreateOrderItem   `json:"items",omitempty"`
	CompanyId   string              `json:"company_id,omitempty"`
}
