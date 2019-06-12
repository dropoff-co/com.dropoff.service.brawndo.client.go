package main

import (
	"fmt"
	"time"

	"dropoff.com/brawndo"
	"github.com/davecgh/go-spew/spew"
)

func testAvailableItems(b *brawndo.Client) {

	var availableItemsRequest brawndo.AvailableItemsRequest
	availableItemsRequest.CompanyId = ""
	availableItemsResponse, err := b.AvailableItems(&availableItemsRequest)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(availableItemsResponse)
	}

}
func testCreateNewOrder(b *brawndo.Client) string {
	return testCreateNewOrderForManagedClient(b, "")
}

func testCreateNewOrderForManagedClient(b *brawndo.Client, company_id string) string {
	var cor brawndo.CreateOrderRequest
	var cor_det brawndo.CreateOrderDetails
	var cor_o, cor_d brawndo.CreateOrderAddress
	var cor_item1, cor_item2 brawndo.CreateOrderItem

	cor_item1.Container = brawndo.CONTAINER_BOX
	cor_item1.Description = "Please handle gently"
	cor_item1.Width = "5"
	cor_item1.Height = "5"
	cor_item1.Depth = "5"
	cor_item1.PersonName = "John Locke"
	cor_item1.Price = "15.99"
	cor_item1.Quantity = 2
	cor_item1.Sku = "123456123456"
	cor_item1.Temperature = brawndo.TEMP_AMBIENT
	cor_item1.Weight = "10"
	cor_item1.Unit = "in"

	cor_item2.Container = brawndo.CONTAINER_BAG
	cor_item2.Description = "This one must really be handled carefully"
	cor_item2.Width = "5"
	cor_item2.Height = "5"
	cor_item2.Depth = "5"
	cor_item2.PersonName = "John Locke"
	cor_item2.Price = "15.99"
	cor_item2.Quantity = 2
	cor_item2.Sku = "123456123456"
	cor_item2.Temperature = brawndo.TEMP_FROZEN
	cor_item2.Weight = "10"
	cor_item2.Unit = "ft"

	cor_det.Quantity = 1
	cor_det.Weight = 5
	cor_det.ETA = "448.5"
	cor_det.Distance = "0.64"
	cor_det.Price = "13.99"
	cor_det.ReadyDate = time.Now().Unix()
	cor_det.Type = "two_hr"
	cor_det.ReferenceCode = "reference code 0001"
	cor_det.ReferenceName = "reference name"

	cor_o.CompanyName = "Dropoff GO Origin"
	cor_o.Email = "noreply+origin@dropoff.com"
	cor_o.Phone = "5125555555"
	cor_o.FirstName = "Napoleon"
	cor_o.LastName = "Bonner"
	cor_o.AddressLine1 = "117 San Jacinto Blvd"
	//cor_o.AddressLine2 = ""
	cor_o.City = "Austin"
	cor_o.State = "TX"
	cor_o.Zip = "78701"
	cor_o.Lat = 30.263706
	cor_o.Lng = -97.741703
	cor_o.Remarks = "Be nice to napoleon"

	cor_d.CompanyName = "Dropoff GO Destination"
	cor_d.Email = "noreply+destination@dropoff.com"
	cor_d.Phone = "5125555555"
	cor_d.FirstName = "Del"
	cor_d.LastName = "Fitzgitibit"
	cor_d.AddressLine1 = "1601 S MoPac Expy"
	cor_d.AddressLine2 = "C301"
	cor_d.City = "Austin"
	cor_d.State = "TX"
	cor_d.Zip = "78746"
	cor_d.Lat = 30.260228
	cor_d.Lng = -97.793359
	//cor_d.Remarks = "Be nice to napoleon";

	cor.Details = &cor_det
	cor.Destination = &cor_d
	cor.Origin = &cor_o

	items := []brawndo.CreateOrderItem{cor_item1, cor_item2}
	cor.Items = items

	if company_id != "" {
		cor.CompanyId = company_id
	}

	res, err := b.CreateOrder(&cor)

	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		spew.Dump(res)
		return res.Data.OrderId
	}
}

func testEstimate(b *brawndo.Client) {
	testEstimateForManagedClient(b, "")
}

func testEstimateForManagedClient(b *brawndo.Client, company_id string) {
	var req brawndo.EstimateRequest

	_, now := time.Now().Zone()

	req.Origin = "117 San Jacinto Blvd, Austin, TX 78701"
	req.Destination = "1601 S MoPac Expy, Austin, TX 78746"
	req.UTCOffset = now
	req.ReadyTimestamp = -1
	req.CompanyId = company_id

	est_res, est_err := b.Estimate(&req)

	if est_err != nil {
		fmt.Println(est_err)
	} else {
		spew.Dump(est_res)
	}
}

func testGetOrder(b *brawndo.Client, order_id string) {
	testGetOrderForManagedClient(b, order_id, "")
}

func testGetOrderForManagedClient(b *brawndo.Client, order_id string, company_id string) {
	var req brawndo.OrderRequest

	req.OrderId = order_id
	req.CompanyId = company_id

	res, err := b.GetOrder(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testGetOrderPageWithLastKey(b *brawndo.Client, last_key string) {
	testGetOrderPageWithLastKeyForManagedClient(b, last_key, "")
}

func testGetOrderPageWithLastKeyForManagedClient(b *brawndo.Client, last_key string, company_id string) {
	var req brawndo.OrderRequest

	req.LastKey = last_key
	req.CompanyId = company_id

	res, err := b.GetOrderPage(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func testGetOrderPage(b *brawndo.Client) {
	testGetOrderPageForManagedClient(b, "")
}

func testGetOrderPageForManagedClient(b *brawndo.Client, company_id string) {
	var req brawndo.OrderRequest
	req.CompanyId = company_id

	res, err := b.GetOrderPage(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testCancelOrder(b *brawndo.Client, order_id string) {
	testCancelOrderForManagedClient(b, order_id, "")
}

func testCancelOrderForManagedClient(b *brawndo.Client, order_id string, company_id string) {
	var req brawndo.OrderRequest

	req.OrderId = order_id
	req.CompanyId = company_id

	res, err := b.CancelOrder(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

// func testSimulateOrder(b *brawndo.Client, market string) {
// 	res, err := b.SimulateOrder(market)

// 	if (err != nil) {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 	}
// }

func testCreateTip(b *brawndo.Client, order_id string, amount string) {
	testCreateTipForManagedClient(b, order_id, amount, "")
}

func testCreateTipForManagedClient(b *brawndo.Client, order_id string, amount string, company_id string) {
	var req brawndo.OrderTipRequest
	req.OrderId = order_id
	req.Amount = amount
	req.CompanyId = company_id

	res, err := b.CreateOrderTip(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testGetTip(b *brawndo.Client, order_id string) {
	testGetTipForManagedClient(b, order_id, "")
}

func testGetTipForManagedClient(b *brawndo.Client, order_id string, company_id string) {
	var req brawndo.OrderTipRequest
	req.OrderId = order_id
	req.CompanyId = company_id

	res, err := b.GetOrderTip(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testDeleteTip(b *brawndo.Client, order_id string) {
	testDeleteTipForManagedClient(b, order_id, "")
}

func testDeleteTipForManagedClient(b *brawndo.Client, order_id string, company_id string) {
	var req brawndo.OrderTipRequest
	req.OrderId = order_id
	req.CompanyId = company_id

	res, err := b.DeleteOrderTip(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testInfo(b *brawndo.Client) {
	res, err := b.Info()

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testAvailableProperties(b *brawndo.Client) {
	var req brawndo.AvailablePropertiesRequest
	req.CompanyId = ""

	res, err := b.AvailableProperties(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testGetSignature(b *brawndo.Client) {
	var req brawndo.GetSignatureRequest
	req.CompanyId = ""
	req.OrderId = ""

	res, err := b.GetSignature(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testGetPickupSignature(b *brawndo.Client) {
	var req brawndo.GetPickupSignatureRequest
	req.CompanyId = ""
	req.OrderId = ""

	res, err := b.GetPickupSignature(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func testDriverActionsMeta(b *brawndo.Client) {
	var req brawndo.DriverActionsMetaRequest
	// req.CompanyId = ""

	res, err := b.DriverActionsMeta(&req)

	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(res)
	}
}

func main() {
	var t brawndo.Transport

	//var managed_company = "111111111111111"
	//var order_id = "22222222222222"

	t.Host = "sandbox-brawndo.dropoff.com"
	t.ApiURL = "https://sandbox-brawndo.dropoff.com/v1"
	t.PublicKey = ""
	t.SecretKey = ""

	var b brawndo.Client
	b.Transport = &t

	testInfo(&b)
	testDriverActionsMeta(&b)
	// testCreateNewOrder(&b)

	// testAvailableItems(&b)
	//testInfo(&b)
	//testAvailableProperties(&b)
	testGetSignature(&b)
	testGetPickupSignature(&b)

	//testEstimate(&b)
	//testEstimateForManagedClient(&b, managed_company)
	//
	//testGetOrder(&b, "")
	//testGetOrderForManagedClient(&b, order_id, managed_company)
	//
	//testGetOrderPage(&b)
	//testGetOrderPageForManagedClient(&b, managed_company)
	//
	//testGetOrderPage(&b)
	//testCreateNewOrderForManagedClient(&b, "")
	//testGetOrderPageForClient(&b, "")
	// testGetOrder(&b, "")
	//testGetOrderPage(&b)
	//testGetOrderPageWithLastKey(&b)
	//new_order_id := testCreateNewOrder(&b)
	//
	//if (new_order_id != "") {
	//	testCancelOrder(&b, new_order_id);
	//} else {
	//	fmt.Println("No order to cancel");
	//}
	//testCancelOrder(&b, "bogus");
	//
	//testSimulateOrder(&b, "austin");

	//testCreateTipForManagedClient(&b, order_id, "5.55", managed_company);
	//testGetTipForManagedClient(&b, order_id, managed_company);
	//testDeleteTipForManagedClient(&b, order_id, managed_company);
	//testGetTipForManagedClient(&b, order_id, managed_company);
}
