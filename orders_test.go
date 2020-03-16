package gotd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestOrdersGetOrdersByPathCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *OrdersGetOrdersByPathCall
		want    *OrderList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewOrdersService(td).GetOrdersByPath(accountID), nil, false},
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

// 暫不測試，因需戶口有錢
func Test_Order(t *testing.T) {
	t.Skip("暫不測試，因需戶口有錢")

	server := NewOrdersService(td)

	// 建立 order
	fmt.Print("PlaceOrderCall\n=================================================================\n")
	order := &Order{
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
