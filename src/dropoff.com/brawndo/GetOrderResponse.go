package brawndo

type GetOrderDetails struct {
	OrderId            string `json:"order_id"`
	CustomerName       string `json:"customer_name"`
	Price              string `json:"price"`
	Distance           string `json:"distance"`
	Quantity           int64  `json:"quantity"`
	Weight             int64  `json:"weight"`
	Market             string `json:"market"`
	ServiceType        string `json:"service_type"`
	TimeFrame          string `json:"time_frame"`
	Timezone           string `json:"timezone"`
	UTCOffsetMinutes   int64  `json:"utc_offset_minutes"`
	CreateDate         int64  `json:"createdate"`
	UpdateDate         int64  `json:"updatedate"`
	ReadyForPickupDate int64  `json:"readyforpickupdate"`
	OrderStatusCode    int64  `json:"order_status_code"`
	OrderStatusName    string `json:"order_status_name"`
	ReferenceCode      string `json:"reference_code,omitempty"`
	ReferenceName      string `json:"reference_name,omitempty"`
}

type GetOrderAddress struct {
	CompanyName   string  `json:"company_name"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AddressLine1  string  `json:"address_line_1"`
	AddressLine2  string  `json:"address_line_2"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	Zip           string  `json:"zip"`
	Lng           float64 `json:"lat"`
	Lat           float64 `json:"lng"`
	Email         string  `json:"email_address"`
	Phone         string  `json:"phone_number"`
	CreateDate    int64   `json:"createdate"`
	UpdateDate    int64   `json:"updatedate"`
	DriverActions string  `json:"workflow_sequence"`
}

type GetOrderProperty struct {
	Id              int64   `json:"id"`
	Label           string  `json:"label"`
	Description     string  `json:"description"`
	PriceAdjustment float64 `json:"price_adjustment"`
}

type GetOrderItem struct {
	Container   int64   `json:"container"`
	Quantity    int64   `json:"quantity"`
	Weight      int64   `json:"weight"`
	Description string  `json:"description"`
	Createdate  int64   `json:"createdate"`
	PersonName  string  `json:"person_name"`
	OrderItemId string  `json:"order_item_id"`
	Unit        string  `json:"unit"`
	Depth       int64   `json:"depth"`
	Updatedate  int64   `json:"updatedate"`
	Price       float64 `json:"price"`
	Temperature int64   `json:"temperature"`
	Width       int64   `json:"width"`
	Sku         string  `json:"sku"`
	OrderId     string  `json:"order_id"`
	Height      int64   `json:"height"`
}

type GetOrderWorkflow struct {
	Result          int64  `json:"result"`
	Updatedate      int64  `json:"updatedate"`
	Action          int64  `json:"action"`
	Createdate      int64  `json:"createdate"`
	OrderWorkflowId string `json:"order_workflow_id"`
	OrderId         string `json:"order_id"`
	WorkerId        string `json:"worker_id"`
}

type GetOrderData struct {
	Url			string				`json:"url"`
	LabelUrl	string				`json:"label_url"`
	Details     *GetOrderDetails    `json:"details"`
	Origin      *GetOrderAddress    `json:"origin"`
	Destination *GetOrderAddress    `json:"destination"`
	Properties  []*GetOrderProperty `json:"properties,omitempty"`
	Items       []*GetOrderItem     `json:"items,omitempty"`
	Workflow    []*GetOrderWorkflow `json:"workflow,omitempty"`
}

type GetOrderResponse struct {
	Data      *GetOrderData `json:"data"`
	Success   bool          `json:"success"`
	Timestamp string        `json:"timestamp"`
}

type GetOrdersResponse struct {
	Total     int64           `json:"total"`
	Count     int64           `json:"count"`
	LastKey   string          `json:"last_key"`
	Data      []*GetOrderData `json:"data"`
	Success   bool            `json:"success"`
	Timestamp string          `json:"timestamp"`
}
