package gotd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSavedOrdersGetSavedOrdersByPathCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *SavedOrdersGetSavedOrdersByPathCall
		want    *SavedOrder
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewSavedOrdersService(td).GetSavedOrdersByPath(accountID), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("SavedOrdersGetSavedOrdersByPathCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavedOrdersGetSavedOrdersByPathCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SavedOrder(t *testing.T) {
	server := NewSavedOrdersService(td)

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
