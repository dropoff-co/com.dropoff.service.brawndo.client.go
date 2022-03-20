package brawndo

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	//"net/http"
	//"crypto"
	//"crypto/hmac"
)

//Line item constants
const LINE_ITEM_DISABLED = 0
const LINE_ITEM_OPTIONAL = 1
const LINE_ITEM_REQUIRED = 2

const TEMP_NA = 0
const TEMP_AMBIENT = 100
const TEMP_REFRIGERATED = 200
const TEMP_FROZEN = 300

const CONTAINER_NA = 0
const CONTAINER_BAG = 100
const CONTAINER_BOX = 200
const CONTAINER_TRAY = 300
const CONTAINER_PALLET = 400
const CONTAINER_BARREL = 500
const CONTAINER_BASKET = 600
const CONTAINER_BUCKET = 700
const CONTAINER_CARTON = 800
const CONTAINER_CASE = 900
const CONTAINER_COOLER = 1000
const CONTAINER_CRATE = 1100
const CONTAINER_TOTE = 1200

type Client struct {
	Transport *Transport
}

func (b Client) Info() (InfoResponse, error) {
	var ir InfoResponse
	resp, err := b.Transport.MakeRequest("GET", "/info", "info", "", nil, "")

	if err != nil {
		return ir, err
	}

	err = json.Unmarshal([]byte(resp), &ir)

	return ir, nil
}

type AvailablePropertiesRequest struct {
	CompanyId string
}

func (b Client) AvailableProperties(req *AvailablePropertiesRequest) (AvailablePropertiesResponse, error) {
	var apgr AvailablePropertiesResponse
	var queryString string

	var companyId = req.CompanyId

	req.CompanyId = ""

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return apgr, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order/properties", "order", queryString, nil, "")

	if err != nil {
		return apgr, err
	}

	err = json.Unmarshal([]byte(resp), &apgr)

	return apgr, nil
}

type DriverActionsMetaRequest struct {
	CompanyId string
}

func (b Client) DriverActionsMeta(req *DriverActionsMetaRequest) (DriverActionsMetaResponse, error) {
	var damr DriverActionsMetaResponse
	var queryString string

	var companyId = req.CompanyId

	req.CompanyId = ""

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return damr, err
		}
		query.Add("company_id", companyId)
		query.Add("go", "true")
		queryString = "?" + query.Encode()
	} else {
		query, err := url.ParseQuery("")
		if err != nil {
			return damr, err
		}
		query.Add("go", "true")
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order/driver_actions_meta", "order", queryString, nil, "")

	fmt.Println("Response: ", resp)
	if err != nil {
		return damr, err
	}

	err = json.Unmarshal([]byte(resp), &damr)

	return damr, nil
}

type AvailableItemsRequest struct {
	CompanyId string
}

func (b Client) AvailableItems(req *AvailableItemsRequest) (AvailableItemsResponse, error) {
	var availableItemsResponse AvailableItemsResponse
	var queryString string

	var companyId = req.CompanyId

	req.CompanyId = ""

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return availableItemsResponse, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order/items", "order", queryString, nil, "")

	if err != nil {
		return availableItemsResponse, err
	}

	err = json.Unmarshal([]byte(resp), &availableItemsResponse)

	return availableItemsResponse, nil
}

type GetSignatureRequest struct {
	CompanyId string
	OrderId   string
}

type GetSignatureResponse struct {
	Url     string `json:"url"`
	Success bool   `json:"success"`
}

func (b Client) GetSignature(req *GetSignatureRequest) (GetSignatureResponse, error) {
	var gsr GetSignatureResponse
	var queryString string

	var companyId = req.CompanyId

	req.CompanyId = ""

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return gsr, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order/signature/"+req.OrderId, "order", queryString, nil, "")

	if err != nil {
		return gsr, err
	}

	err = json.Unmarshal([]byte(resp), &gsr)

	return gsr, nil
}

type GetPickupSignatureRequest struct {
	CompanyId string
	OrderId   string
}

type GetPickupSignatureResponse struct {
	Url     string `json:"url"`
	Success bool   `json:"success"`
}

func (b Client) GetPickupSignature(req *GetPickupSignatureRequest) (GetPickupSignatureResponse, error) {
	var gsr GetPickupSignatureResponse
	var queryString string

	var companyId = req.CompanyId

	req.CompanyId = ""

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return gsr, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order/pickup_signature/"+req.OrderId, "order", queryString, nil, "")

	if err != nil {
		return gsr, err
	}

	err = json.Unmarshal([]byte(resp), &gsr)

	return gsr, nil
}

type EstimateRequest struct {
	Origin         string
	Destination    string
	UTCOffset      int
	ReadyTimestamp int64
	CompanyId      string
}

func (b Client) Estimate(req *EstimateRequest) (EstimateResponse, error) {
	var origin = req.Origin
	var destination = req.Destination
	var readyTimestamp = req.ReadyTimestamp
	var utcOffest = req.UTCOffset
	var companyId = req.CompanyId

	var er EstimateResponse
	query, err := url.ParseQuery("")

	if err != nil {
		return er, err
	}

	query.Add("origin", origin)
	query.Add("destination", destination)

	if companyId != "" {
		query.Add("company_id", companyId)
	}

	var utcOffsetString string

	var utcOffsetHours = utcOffest / 3600

	if utcOffsetHours < 0 {
		utcOffsetString = "-"
		utcOffsetHours = utcOffsetHours * -1
	}

	if utcOffsetHours < 10 {
		utcOffsetString += "0"
	}

	utcOffsetString += strconv.Itoa(utcOffsetHours) + ":00"

	query.Add("utc_offset", utcOffsetString)

	if readyTimestamp > 0 {
		query.Add("ready_timestamp", strconv.FormatInt(readyTimestamp, 10))
	}

	resp, err := b.Transport.MakeRequest("GET", "/estimate", "estimate", "?"+query.Encode(), nil, "")

	if err != nil {
		return er, err
	}

	err = json.Unmarshal([]byte(resp), &er)

	return er, nil
}

func (b Client) CreateOrder(req *CreateOrderRequest) (CreateOrderResponse, error) {
	var cor CreateOrderResponse
	var queryString string
	var companyId = req.CompanyId

	req.CompanyId = ""

	ba, err := json.Marshal(req)

	if err != nil {
		return cor, err
	}

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return cor, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("POST", "/order", "order", queryString, ba, "")

	if err != nil {
		return cor, err
	}

	err = json.Unmarshal([]byte(resp), &cor)

	if err != nil {
		return cor, err
	}

	return cor, nil
}

type OrderRequest struct {
	OrderId   string
	LastKey   string
	CompanyId string
}

func (b Client) GetOrder(req *OrderRequest) (GetOrderResponse, error) {
	var gor GetOrderResponse

	var orderId = req.OrderId
	var companyId = req.CompanyId

	var queryString string

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return gor, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order/"+orderId, "order", queryString, nil, "")

	if err != nil {
		return gor, err
	}

	err = json.Unmarshal([]byte(resp), &gor)

	if err != nil {
		return gor, err
	}

	return gor, nil
}

func (b Client) GetOrderPage(req *OrderRequest) (GetOrdersResponse, error) {
	var gor GetOrdersResponse
	var queryString string
	var lastKey = req.LastKey
	var companyId = req.CompanyId

	if lastKey != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return gor, err
		}
		query.Add("last_key", lastKey)
		if companyId != "" {
			query.Add("company_id", companyId)
		}
		queryString = "?" + query.Encode()
	} else if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return gor, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order", "order", queryString, nil, "")

	if err != nil {
		return gor, err
	}

	err = json.Unmarshal([]byte(resp), &gor)

	if err != nil {
		return gor, err
	}

	return gor, nil
}

func (b Client) CancelOrder(req *OrderRequest) (CancelOrderResponse, error) {
	var cor CancelOrderResponse
	var queryString string
	var orderId = req.OrderId
	var companyId = req.CompanyId

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return cor, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("POST", "/order/"+orderId+"/cancel", "order", queryString, nil, "")

	if err != nil {
		return cor, err
	}

	err = json.Unmarshal([]byte(resp), &cor)

	if err != nil {
		return cor, err
	}

	return cor, nil
}

type SimulateRequest struct {
	OrderId   string
	Market    string
	CompanyId string
}

func (b Client) SimulateOrder(req *SimulateRequest) (SimulateOrderResponse, error) {
	var sor SimulateOrderResponse
	var queryString string

	if req.CompanyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return sor, err
		}
		query.Add("company_id", req.CompanyId)
		queryString = "?" + query.Encode()
	}

	if req.OrderId != "" {
		resp, err := b.Transport.MakeRequest("GET", "/order/simulate/order/"+req.OrderId, "order", queryString, nil, "")

		if err != nil {
			return sor, err
		}

		err = json.Unmarshal([]byte(resp), &sor)

		if err != nil {
			return sor, err
		}
	} else if req.Market != "" {
		resp, err := b.Transport.MakeRequest("GET", "/order/simulate/"+req.Market, "order", queryString, nil, "")

		if err != nil {
			return sor, err
		}

		err = json.Unmarshal([]byte(resp), &sor)

		if err != nil {
			return sor, err
		}
	}

	return sor, nil
}

type OrderTipRequest struct {
	OrderId   string
	Amount    string
	CompanyId string
}

func (b Client) CreateOrderTip(req *OrderTipRequest) (TipResponse, error) {
	var tr TipResponse

	var queryString string

	var orderId = req.OrderId
	var amount = req.Amount
	var companyId = req.CompanyId

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return tr, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("POST", "/order/"+orderId+"/tip/"+amount, "order", queryString, nil, "")

	if err != nil {
		return tr, err
	}

	err = json.Unmarshal([]byte(resp), &tr)

	if err != nil {
		return tr, err
	}

	return tr, nil
}

func (b Client) GetOrderTip(req *OrderTipRequest) (GetTipResponse, error) {
	var tr GetTipResponse

	var queryString string
	var orderId = req.OrderId
	var companyId = req.CompanyId

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return tr, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("GET", "/order/"+orderId+"/tip", "order", queryString, nil, "")

	if err != nil {
		return tr, err
	}

	err = json.Unmarshal([]byte(resp), &tr)

	if err != nil {
		return tr, err
	}

	return tr, nil
}

func (b Client) DeleteOrderTip(req *OrderTipRequest) (DeleteTipResponse, error) {
	var tr DeleteTipResponse

	var queryString string
	var orderId = req.OrderId
	var companyId = req.CompanyId

	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return tr, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("DELETE", "/order/"+orderId+"/tip", "order", queryString, nil, "")

	if err != nil {
		return tr, err
	}

	err = json.Unmarshal([]byte(resp), &tr)

	if err != nil {
		return tr, err
	}

	return tr, nil
}

func (b Client) CreateBulk(req *CreateBulkRequest) (CreateBulkResponse, error) {
	var cbr CreateBulkResponse
	var queryString string
	var companyId = ""
	var file = "./shortest copy.csv"
	ba, err := json.Marshal(req)
	if err != nil {
		return cbr, err
	}
	if companyId != "" {
		query, err := url.ParseQuery("")
		if err != nil {
			return cbr, err
		}
		query.Add("company_id", companyId)
		queryString = "?" + query.Encode()
	}

	resp, err := b.Transport.MakeRequest("POST", "/bulkupload/nosign", "bulkupload", queryString, ba, file)

	if err != nil {
		return cbr, err
	}
	err = json.Unmarshal([]byte(resp), &cbr)

	if err != nil {
		return cbr, err
	}
	return cbr, nil
}
