package brawndo

type AvailableItemsResponseData struct {
	OrderItemEnabled         int    `json:"order_item_enabled"`
	CompanyId                string `json:"company_id"`
	OrderItemAllowSku        int    `json:"order_item_allow_sku"`
	OrderItemTemperatureUnit string `json:"order_item_temp_unit"`
	OrderItemPersonNameLabel string `json:"order_item_person_name_label"`

	OrderItemAllowWeight      int `json:"order_item_allow_weight"`
	OrderItemAllowPersonName  int `json:"order_item_allow_person_name"`
	OrderItemAllowQuantity    int `json:"order_item_allow_quantity"`
	OrderItemAllowDescription int `json:"order_item_allow_description"`
	OderItemAllowDimensions   int `json:"order_item_allow_dimensions"`
	OrderItemAllowContainer   int `json:"order_item_allow_container"`
	OrderItemAllowTemperature int `json:"order_item_allow_temperature"`
	OrderItemAllowPrice       int `json:"order_item_allow_price"`
}

type AvailableItemsResponse struct {
	Data      *AvailableItemsResponseData `json:"data"`
	Success   bool                        `json:"success"`
	Timestamp string                      `json:"timestamp"`
}
