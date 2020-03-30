package gotd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_SavedOrder(t *testing.T) {
	if !onlineTest {
		t.Skipf("online Test setting is %v", onlineTest)
	}

	server := NewSavedOrdersService(tdReal)

	// 建立 save order
	fmt.Print("CreateSavedOrderCall\n=================================================================\n")
	savedOrder := &SavedOrder{
		Order: &Order{
			Session:    "NORMAL",
			Duration:   "GOOD_TILL_CANCEL",
			OrderType:  "LIMIT",
			CancelTime: time.Now().AddDate(0, 4, 0).UTC().Format("2006-01-02"),
			Price:      1.1,
			OrderLegCollections: []*OrderLegCollection{
				&OrderLegCollection{
					Instrument: &Instrument{
						Symbol:    "BWX",
						AssetType: OrderAssetTypeEQUITY,
					},
					Instruction: OrderInstructionBuy,
					Quantity:    2,
				},
			},
		},
	}

	CreateSavedOrderCall := server.CreateSavedOrder(accountID, savedOrder)
	resp, err := CreateSavedOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("CreateSavedOrderCall: %d\n", resp.HTTPStatusCode)
	fmt.Print("=================================================================\n")

	// 獲得建立的 save order
	fmt.Print("GetSavedOrdersByPathCall\n=================================================================\n")
	GetSavedOrdersByPathCall := server.GetSavedOrdersByPath(accountID)
	savedOrders, err := GetSavedOrdersByPathCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	content, _ := json.MarshalIndent(savedOrders.SavedOrders[0], "", "    ")
	t.Log(string(content))
	fmt.Print("=================================================================\n")

	// 取代建立的 save order
	fmt.Print("ReplaceSavedOrderCall\n=================================================================\n")
	savedOrder.OrderLegCollections[0].Instrument.Symbol = "VTI"
	ReplaceSavedOrderCall := server.ReplaceSavedOrder(accountID, savedOrders.SavedOrders[0].SavedOrderID, savedOrder)
	resp, err = ReplaceSavedOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ReplaceSavedOrderCall: %d\n", resp.HTTPStatusCode)
	fmt.Print("=================================================================\n")

	// 獲得取代的 save order
	fmt.Print("GetSavedOrderCall\n=================================================================\n")
	GetSavedOrderCall := server.GetSavedOrder(accountID, savedOrders.SavedOrders[0].SavedOrderID)
	savedOrder, err = GetSavedOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	content, _ = json.MarshalIndent(savedOrder, "", "    ")
	t.Log(string(content))
	fmt.Print("=================================================================\n")

	// 刪除取代的 save order
	fmt.Print("DeleteSavedOrderCall\n=================================================================\n")
	DeleteSavedOrderCall := server.DeleteSavedOrder(accountID, savedOrders.SavedOrders[0].SavedOrderID)
	_, err = DeleteSavedOrderCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print("=================================================================\n")

	// 獲得刪除的 save order，應為空
	fmt.Print("GetSavedOrderCall\n=================================================================\n")
	GetSavedOrderCall = server.GetSavedOrder(accountID, savedOrders.SavedOrders[0].SavedOrderID)
	_, err = GetSavedOrderCall.Do()
	if err == nil {
		content, _ = json.MarshalIndent(savedOrder, "", "    ")
		t.Fatal(string(content))
	}
	fmt.Print("=================================================================\n")
}

func TestSavedOrdersReplaceSavedOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersReplaceSavedOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).ReplaceSavedOrder("accountId", 1234, &SavedOrder{}), "https://api.tdameritrade.com/v1/accounts/accountId/savedorders/1234?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersReplaceSavedOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersReplaceSavedOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersReplaceSavedOrderCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersReplaceSavedOrderCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).ReplaceSavedOrder("accountId", 1234, &SavedOrder{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersReplaceSavedOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersReplaceSavedOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersGetSavedOrdersByPathCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersGetSavedOrdersByPathCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).GetSavedOrdersByPath("accountId"), "https://api.tdameritrade.com/v1/accounts/accountId/savedorders?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersGetSavedOrdersByPathCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersGetSavedOrdersByPathCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersGetSavedOrdersByPathCall_Do(t *testing.T) {
	client := clientTest(`[{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-03","complexOrderStrategyType":"NONE","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","symbol":"VBR"},"instruction":"BUY","quantity":1}],"orderStrategyType":"SINGLE","cancelable":true,"editable":true,"savedOrderId":12345678,"savedTime":"2018-11-03T09:25:19+0000"},{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-03","complexOrderStrategyType":"NONE","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","symbol":"VBR"},"instruction":"BUY","quantity":1}],"orderStrategyType":"SINGLE","cancelable":true,"editable":true,"savedOrderId":12345678,"savedTime":"2018-11-03T09:25:19+0000"}]`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersGetSavedOrdersByPathCall
		want    []*SavedOrder
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).GetSavedOrdersByPath("accountId"), []*SavedOrder{
			&SavedOrder{
				Order: &Order{
					Session:                  "NORMAL",
					Duration:                 "GOOD_TILL_CANCEL",
					OrderType:                "LIMIT",
					CancelTime:               "2019-03-03",
					ComplexOrderStrategyType: "NONE",
					Price:                    1,
					OrderLegCollections: []*OrderLegCollection{
						&OrderLegCollection{
							OrderLegType: "EQUITY",
							LegID:        1,
							Instrument: &Instrument{
								AssetType: "EQUITY",
								Symbol:    "VBR",
							},
							Instruction: "BUY",
							Quantity:    1,
						},
					},
					OrderStrategyType: "SINGLE",
					Cancelable:        true,
					Editable:          true,
				},
				SavedOrderID: 12345678,
				SavedTime:    "2018-11-03T09:25:19+0000",
			},
			&SavedOrder{
				Order: &Order{
					Session:                  "NORMAL",
					Duration:                 "GOOD_TILL_CANCEL",
					OrderType:                "LIMIT",
					CancelTime:               "2019-03-03",
					ComplexOrderStrategyType: "NONE",
					Price:                    1,
					OrderLegCollections: []*OrderLegCollection{
						&OrderLegCollection{
							OrderLegType: "EQUITY",
							LegID:        1,
							Instrument: &Instrument{
								AssetType: "EQUITY",
								Symbol:    "VBR",
							},
							Instruction: "BUY",
							Quantity:    1,
						},
					},
					OrderStrategyType: "SINGLE",
					Cancelable:        true,
					Editable:          true,
				},
				SavedOrderID: 12345678,
				SavedTime:    "2018-11-03T09:25:19+0000",
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersGetSavedOrdersByPathCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.SavedOrders
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersGetSavedOrdersByPathCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersGetSavedOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersGetSavedOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).GetSavedOrder("accountId", 1234), "https://api.tdameritrade.com/v1/accounts/accountId/savedorders/1234?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersGetSavedOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersGetSavedOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersGetSavedOrderCall_Do(t *testing.T) {
	client := clientTest(`{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-03","complexOrderStrategyType":"NONE","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","symbol":"VBR"},"instruction":"BUY","quantity":1}],"orderStrategyType":"SINGLE","cancelable":true,"editable":true,"savedOrderId":12345678,"savedTime":"2018-11-03T09:25:19+0000"}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersGetSavedOrderCall
		want    *SavedOrder
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).GetSavedOrder("accountId", 1234), &SavedOrder{
			ServerResponse: ServerResponse{200, http.Header{}},
			Order: &Order{
				Session:                  "NORMAL",
				Duration:                 "GOOD_TILL_CANCEL",
				OrderType:                "LIMIT",
				CancelTime:               "2019-03-03",
				ComplexOrderStrategyType: "NONE",
				Price:                    1,
				OrderLegCollections: []*OrderLegCollection{
					&OrderLegCollection{
						OrderLegType: "EQUITY",
						LegID:        1,
						Instrument: &Instrument{
							AssetType: "EQUITY",
							Symbol:    "VBR",
						},
						Instruction: "BUY",
						Quantity:    1,
					},
				},
				OrderStrategyType: "SINGLE",
				Cancelable:        true,
				Editable:          true,
			},
			SavedOrderID: 12345678,
			SavedTime:    "2018-11-03T09:25:19+0000",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersGetSavedOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersGetSavedOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersDeleteSavedOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersDeleteSavedOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).DeleteSavedOrder("accountId", 1234), "https://api.tdameritrade.com/v1/accounts/accountId/savedorders/1234?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersDeleteSavedOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersDeleteSavedOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersDeleteSavedOrderCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersDeleteSavedOrderCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).DeleteSavedOrder("accountId", 1234), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersDeleteSavedOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersDeleteSavedOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersCreateSavedOrderCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersCreateSavedOrderCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).CreateSavedOrder("accountId", &SavedOrder{}), "https://api.tdameritrade.com/v1/accounts/accountId/savedorders?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersCreateSavedOrderCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersCreateSavedOrderCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavedOrdersCreateSavedOrderCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *SavedOrdersCreateSavedOrderCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewSavedOrdersService(tdTest).CreateSavedOrder("accountId", &SavedOrder{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersCreateSavedOrderCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersCreateSavedOrderCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
