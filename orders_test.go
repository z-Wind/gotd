package gotd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// 暫不測試，因需戶口有錢
func Test_Order(t *testing.T) {
	t.Skip("暫不測試，因需戶口有錢")
	if !onlineTest {
		t.Skipf("online Test setting is %v", onlineTest)
	}

	server := NewOrdersService(tdReal)

	// 建立 order
	fmt.Print("PlaceOrderCall\n=================================================================\n")
	order := &Order{
		Session:    "NORMAL",
		Duration:   "GOOD_TILL_CANCEL",
		OrderType:  "LIMIT",
		CancelTime: time.Now().AddDate(0, 4, 0).UTC().Format("2006-01-02"),
		Price:      1.1,
		OrderLegCollections: []*OrderLegCollection{
			{
				Instrument: &Instrument{
					Symbol:    "BWX",
					AssetType: OrderAssetTypeEQUITY,
				},
				Instruction: OrderInstructionBuy,
				Quantity:    2,
			},
		},
	}

	PlaceOrderCall := server.PlaceOrder(accountID, order)
	resp, err := PlaceOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PlaceOrderCall: %d\n", resp.HTTPStatusCode)
	fmt.Print("=================================================================\n")

	// 獲得建立的 order
	fmt.Print("GetOrdersByQueryCall\n=================================================================\n")
	GetOrdersByQueryCall := server.GetOrdersByQuery(accountID)
	GetOrdersByQueryCall.MaxResults(10)
	GetOrdersByQueryCall.FromEnteredTime(time.Now().AddDate(0, 0, -1))
	GetOrdersByQueryCall.ToEnteredTime(time.Now())
	GetOrdersByQueryCall.Status(OrdersStatusWORKING)
	Orders, err := GetOrdersByQueryCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	content, _ := json.MarshalIndent(Orders.Orders[0], "", "    ")
	t.Log(string(content))
	fmt.Print("=================================================================\n")

	fmt.Print("GetOrdersByPathCall\n=================================================================\n")
	GetOrdersByPathCall := server.GetOrdersByPath(accountID)
	GetOrdersByPathCall.MaxResults(10)
	GetOrdersByPathCall.FromEnteredTime(time.Now().AddDate(0, 0, -1))
	GetOrdersByPathCall.ToEnteredTime(time.Now())
	GetOrdersByPathCall.Status(OrdersStatusWORKING)
	Orders, err = GetOrdersByPathCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	content, _ = json.MarshalIndent(Orders.Orders[0], "", "    ")
	t.Log(string(content))
	fmt.Print("=================================================================\n")

	// 取代建立的 order
	fmt.Print("ReplaceOrderCall\n=================================================================\n")
	order.OrderLegCollections[0].Instrument.Symbol = "VTI"
	ReplaceOrderCall := server.ReplaceOrder(accountID, Orders.Orders[0].OrderID, order)
	resp, err = ReplaceOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ReplaceOrderCall: %d\n", resp.HTTPStatusCode)
	fmt.Print("=================================================================\n")

	// 獲得取代的 order
	fmt.Print("GetOrderCall\n=================================================================\n")
	GetOrderCall := server.GetOrder(accountID, Orders.Orders[0].OrderID)
	order, err = GetOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	content, _ = json.MarshalIndent(order, "", "    ")
	t.Log(string(content))
	fmt.Print("=================================================================\n")

	// 取消取代的 order
	fmt.Print("CancelOrderCall\n=================================================================\n")
	CancelOrderCall := server.CancelOrder(accountID, Orders.Orders[0].OrderID)
	_, err = CancelOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print("=================================================================\n")

	// 獲得取消的 order，狀態應為 cancel
	fmt.Print("GetOrderCall\n=================================================================\n")
	GetOrderCall = server.GetOrder(accountID, Orders.Orders[0].OrderID)
	order, err = GetOrderCall.Do()
	if order.Status != OrdersStatusCANCELED {
		content, _ = json.MarshalIndent(order, "", "    ")
		t.Fatal(string(content))
	}
	fmt.Print("=================================================================\n")
}

func TestOrdersReplaceOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersReplaceOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewOrdersService(tdTest).ReplaceOrder("accountId", 1234, &Order{}), "https://api.tdameritrade.com/v1/accounts/accountId/orders/1234?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersReplaceOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersReplaceOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersReplaceOrderCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersReplaceOrderCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).ReplaceOrder("accountId", 1234, &Order{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersReplaceOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersReplaceOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersGetOrdersByQueryCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersGetOrdersByQueryCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewOrdersService(tdTest).GetOrdersByQuery("accountID"), "https://api.tdameritrade.com/v1/orders?accountId=accountID", false},
		{"maxResults", NewOrdersService(tdTest).GetOrdersByQuery("accountID").MaxResults(100), "https://api.tdameritrade.com/v1/orders?accountId=accountID&maxResults=100", false},
		{"fromEnteredTime", NewOrdersService(tdTest).GetOrdersByQuery("accountID").FromEnteredTime(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "https://api.tdameritrade.com/v1/orders?accountId=accountID&fromEnteredTime=2020-01-01", false},
		{"toEnteredTime", NewOrdersService(tdTest).GetOrdersByQuery("accountID").ToEnteredTime(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "https://api.tdameritrade.com/v1/orders?accountId=accountID&toEnteredTime=2020-01-01", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusAWAITINGPARENTORDER), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=AWAITING_PARENT_ORDER", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusAWAITINGCONDITION), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=AWAITING_CONDITION", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusAWAITINGMANUALREVIEW), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=AWAITING_MANUAL_REVIEW", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusACCEPTED), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=ACCEPTED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusAWAITINGUROUT), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=AWAITING_UR_OUT", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusPENDINGACTIVATION), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=PENDING_ACTIVATION", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusQUEUED), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=QUEUED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusWORKING), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=WORKING", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusREJECTED), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=REJECTED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusPENDINGCANCEL), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=PENDING_CANCEL", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusCANCELED), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=CANCELED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusPENDINGREPLACE), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=PENDING_REPLACE", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusREPLACED), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=REPLACED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusFILLED), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=FILLED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrdersStatusEXPIRED), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=EXPIRED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderInstructionBuy), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=BUY", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderInstructionSELL), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=SELL", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderAssetTypeEQUITY), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=EQUITY", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderAssetTypeOPTION), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=OPTION", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderAssetTypeINDEX), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=INDEX", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderAssetTypeMUTUALFUND), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=MUTUAL_FUND", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderAssetTypeCASHEQUIVALENT), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=CASH_EQUIVALENT", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderAssetTypeFIXEDINCOME), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=FIXED_INCOME", false},
		{"status", NewOrdersService(tdTest).GetOrdersByQuery("accountID").Status(OrderAssetTypeCURRENCYOrderAssetType), "https://api.tdameritrade.com/v1/orders?accountId=accountID&status=CURRENCYOrderAssetType", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersGetOrdersByQueryCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersGetOrdersByQueryCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersGetOrdersByQueryCall_Do(t *testing.T) {
	client := clientTest(`[{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":2,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:26:11+0000","closeTime":"2018-11-03T08:26:11+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."},{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":20,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:26:01+0000","closeTime":"2018-11-03T08:26:01+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"You do not have enough available cash/buying power for this order."}]`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersGetOrdersByQueryCall
		want    *OrderList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).GetOrdersByQuery("accountID"), &OrderList{
			ServerResponse: ServerResponse{200, http.Header{}},
			Orders: []*Order{
				{
					Session:                  "NORMAL",
					Duration:                 "GOOD_TILL_CANCEL",
					OrderType:                "LIMIT",
					CancelTime:               "2019-03-01",
					ComplexOrderStrategyType: "NONE",
					Quantity:                 1,
					FilledQuantity:           0,
					RemainingQuantity:        0,
					RequestedDestination:     "AUTO",
					DestinationLinkName:      "AutoRoute",
					Price:                    2,
					OrderLegCollections: []*OrderLegCollection{
						{
							OrderLegType: "EQUITY",
							LegID:        1,
							Instrument: &Instrument{
								AssetType: "EQUITY",
								Cusip:     "123456789",
								Symbol:    "VTI",
							},
							Instruction:    "BUY",
							PositionEffect: "OPENING",
							Quantity:       1,
						},
					},
					OrderStrategyType: "SINGLE",
					OrderID:           123456789,
					Cancelable:        false,
					Editable:          false,
					Status:            "REJECTED",
					EnteredTime:       "2018-11-03T08:26:11+0000",
					CloseTime:         "2018-11-03T08:26:11+0000",
					Tag:               "AA_ABCDEFGHIJK",
					AccountID:         123456789,
					StatusDescription: "Your limit price is too distant from the current quote. Try again.",
				},
				{
					Session:                  "NORMAL",
					Duration:                 "GOOD_TILL_CANCEL",
					OrderType:                "LIMIT",
					CancelTime:               "2019-03-01",
					ComplexOrderStrategyType: "NONE",
					Quantity:                 1,
					FilledQuantity:           0,
					RemainingQuantity:        0,
					RequestedDestination:     "AUTO",
					DestinationLinkName:      "AutoRoute",
					Price:                    20,
					OrderLegCollections: []*OrderLegCollection{
						{
							OrderLegType: "EQUITY",
							LegID:        1,
							Instrument: &Instrument{
								AssetType: "EQUITY",
								Cusip:     "123456789",
								Symbol:    "VTI",
							},
							Instruction:    "BUY",
							PositionEffect: "OPENING",
							Quantity:       1,
						},
					},
					OrderStrategyType: "SINGLE",
					OrderID:           123456789,
					Cancelable:        false,
					Editable:          false,
					Status:            "REJECTED",
					EnteredTime:       "2018-11-03T08:26:01+0000",
					CloseTime:         "2018-11-03T08:26:01+0000",
					Tag:               "AA_ABCDEFGHIJK",
					AccountID:         123456789,
					StatusDescription: "You do not have enough available cash/buying power for this order.",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersGetOrdersByQueryCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersGetOrdersByQueryCall.Do() = \n%+v\n, want \n%+v", got, tt.want)
			}
		})
	}
}

func TestOrdersGetOrdersByPathCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersGetOrdersByPathCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewOrdersService(tdTest).GetOrdersByPath("accountID"), "https://api.tdameritrade.com/v1/accounts/accountID/orders?", false},
		{"maxResults", NewOrdersService(tdTest).GetOrdersByPath("accountID").MaxResults(100), "https://api.tdameritrade.com/v1/accounts/accountID/orders?maxResults=100", false},
		{"fromEnteredTime", NewOrdersService(tdTest).GetOrdersByPath("accountID").FromEnteredTime(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "https://api.tdameritrade.com/v1/accounts/accountID/orders?fromEnteredTime=2020-01-01", false},
		{"toEnteredTime", NewOrdersService(tdTest).GetOrdersByPath("accountID").ToEnteredTime(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "https://api.tdameritrade.com/v1/accounts/accountID/orders?toEnteredTime=2020-01-01", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusAWAITINGPARENTORDER), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=AWAITING_PARENT_ORDER", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusAWAITINGCONDITION), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=AWAITING_CONDITION", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusAWAITINGMANUALREVIEW), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=AWAITING_MANUAL_REVIEW", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusACCEPTED), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=ACCEPTED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusAWAITINGUROUT), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=AWAITING_UR_OUT", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusPENDINGACTIVATION), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=PENDING_ACTIVATION", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusQUEUED), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=QUEUED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusWORKING), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=WORKING", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusREJECTED), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=REJECTED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusPENDINGCANCEL), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=PENDING_CANCEL", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusCANCELED), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=CANCELED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusPENDINGREPLACE), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=PENDING_REPLACE", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusREPLACED), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=REPLACED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusFILLED), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=FILLED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrdersStatusEXPIRED), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=EXPIRED", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderInstructionBuy), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=BUY", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderInstructionSELL), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=SELL", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderAssetTypeEQUITY), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=EQUITY", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderAssetTypeOPTION), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=OPTION", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderAssetTypeINDEX), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=INDEX", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderAssetTypeMUTUALFUND), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=MUTUAL_FUND", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderAssetTypeCASHEQUIVALENT), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=CASH_EQUIVALENT", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderAssetTypeFIXEDINCOME), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=FIXED_INCOME", false},
		{"status", NewOrdersService(tdTest).GetOrdersByPath("accountID").Status(OrderAssetTypeCURRENCYOrderAssetType), "https://api.tdameritrade.com/v1/accounts/accountID/orders?status=CURRENCYOrderAssetType", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersGetOrdersByPathCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersGetOrdersByPathCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersGetOrdersByPathCall_Do(t *testing.T) {
	client := clientTest(`[{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":2,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:26:11+0000","closeTime":"2018-11-03T08:26:11+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."},{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":20,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:26:01+0000","closeTime":"2018-11-03T08:26:01+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"You do not have enough available cash/buying power for this order."}]`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersGetOrdersByPathCall
		want    *OrderList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).GetOrdersByPath("accountID"), &OrderList{
			ServerResponse: ServerResponse{200, http.Header{}},
			Orders: []*Order{
				{
					Session:                  "NORMAL",
					Duration:                 "GOOD_TILL_CANCEL",
					OrderType:                "LIMIT",
					CancelTime:               "2019-03-01",
					ComplexOrderStrategyType: "NONE",
					Quantity:                 1,
					FilledQuantity:           0,
					RemainingQuantity:        0,
					RequestedDestination:     "AUTO",
					DestinationLinkName:      "AutoRoute",
					Price:                    2,
					OrderLegCollections: []*OrderLegCollection{
						{
							OrderLegType: "EQUITY",
							LegID:        1,
							Instrument: &Instrument{
								AssetType: "EQUITY",
								Cusip:     "123456789",
								Symbol:    "VTI",
							},
							Instruction:    "BUY",
							PositionEffect: "OPENING",
							Quantity:       1,
						},
					},
					OrderStrategyType: "SINGLE",
					OrderID:           123456789,
					Cancelable:        false,
					Editable:          false,
					Status:            "REJECTED",
					EnteredTime:       "2018-11-03T08:26:11+0000",
					CloseTime:         "2018-11-03T08:26:11+0000",
					Tag:               "AA_ABCDEFGHIJK",
					AccountID:         123456789,
					StatusDescription: "Your limit price is too distant from the current quote. Try again.",
				},
				{
					Session:                  "NORMAL",
					Duration:                 "GOOD_TILL_CANCEL",
					OrderType:                "LIMIT",
					CancelTime:               "2019-03-01",
					ComplexOrderStrategyType: "NONE",
					Quantity:                 1,
					FilledQuantity:           0,
					RemainingQuantity:        0,
					RequestedDestination:     "AUTO",
					DestinationLinkName:      "AutoRoute",
					Price:                    20,
					OrderLegCollections: []*OrderLegCollection{
						{
							OrderLegType: "EQUITY",
							LegID:        1,
							Instrument: &Instrument{
								AssetType: "EQUITY",
								Cusip:     "123456789",
								Symbol:    "VTI",
							},
							Instruction:    "BUY",
							PositionEffect: "OPENING",
							Quantity:       1,
						},
					},
					OrderStrategyType: "SINGLE",
					OrderID:           123456789,
					Cancelable:        false,
					Editable:          false,
					Status:            "REJECTED",
					EnteredTime:       "2018-11-03T08:26:01+0000",
					CloseTime:         "2018-11-03T08:26:01+0000",
					Tag:               "AA_ABCDEFGHIJK",
					AccountID:         123456789,
					StatusDescription: "You do not have enough available cash/buying power for this order.",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersGetOrdersByPathCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersGetOrdersByPathCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersGetOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersGetOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).GetOrder("accountID", 1234), "https://api.tdameritrade.com/v1/accounts/accountID/orders/1234?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersGetOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersGetOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersGetOrderCall_Do(t *testing.T) {
	client := clientTest(`{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":2,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:26:11+0000","closeTime":"2018-11-03T08:26:11+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersGetOrderCall
		want    *Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).GetOrder("accountID", 1234), &Order{
			ServerResponse:           ServerResponse{200, http.Header{}},
			Session:                  "NORMAL",
			Duration:                 "GOOD_TILL_CANCEL",
			OrderType:                "LIMIT",
			CancelTime:               "2019-03-01",
			ComplexOrderStrategyType: "NONE",
			Quantity:                 1,
			FilledQuantity:           0,
			RemainingQuantity:        0,
			RequestedDestination:     "AUTO",
			DestinationLinkName:      "AutoRoute",
			Price:                    2,
			OrderLegCollections: []*OrderLegCollection{
				{
					OrderLegType: "EQUITY",
					LegID:        1,
					Instrument: &Instrument{
						AssetType: "EQUITY",
						Cusip:     "123456789",
						Symbol:    "VTI",
					},
					Instruction:    "BUY",
					PositionEffect: "OPENING",
					Quantity:       1,
				},
			},
			OrderStrategyType: "SINGLE",
			OrderID:           123456789,
			Cancelable:        false,
			Editable:          false,
			Status:            "REJECTED",
			EnteredTime:       "2018-11-03T08:26:11+0000",
			CloseTime:         "2018-11-03T08:26:11+0000",
			Tag:               "AA_ABCDEFGHIJK",
			AccountID:         123456789,
			StatusDescription: "Your limit price is too distant from the current quote. Try again.",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersGetOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersGetOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersCancelOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersCancelOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).CancelOrder("accountID", 1234), "https://api.tdameritrade.com/v1/accounts/accountID/orders/1234?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersCancelOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersCancelOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersCancelOrderCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersCancelOrderCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).CancelOrder("accountID", 1234), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersCancelOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersCancelOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersPlaceOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersPlaceOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).PlaceOrder("accountID", &Order{}), "https://api.tdameritrade.com/v1/accounts/accountID/orders?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersPlaceOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersPlaceOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdersPlaceOrderCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *OrdersPlaceOrderCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(tdTest).PlaceOrder("accountID", &Order{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrdersPlaceOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdersPlaceOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
