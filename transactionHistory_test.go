package gotd

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestTransactionHistoryGetTransactionListCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TransactionHistoryGetTransactionListCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewTransactionHistoryService(tdTest).GetTransactionList("accountId"), "https://api.tdameritrade.com/v1/accounts/accountId/transactions?", false},
		{"All", NewTransactionHistoryService(tdTest).GetTransactionList("accountId").
			Type(TransactionsKindALL).
			Symbol("symbol").
			StartDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)).
			EndDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "https://api.tdameritrade.com/v1/accounts/accountId/transactions?endDate=2020-01-01&startDate=2020-01-01&symbol=symbol&type=ALL", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryGetTransactionListCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionHistoryGetTransactionListCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTransactionHistoryGetTransactionListCall_Do(t *testing.T) {
	client := clientTest(`[{"type":"TRADE","subAccount":"1","settlementDate":"2018-10-09","netAmount":-88.888,"transactionDate":"2018-10-04T19:17:23+0000","transactionSubType":"DR","transactionId":12345678919,"cashBalanceEffectFlag":true,"description":"DIVIDEND REINVEST","fees":{"rFee":0,"additionalFee":0,"cdscFee":0,"regFee":0,"otherCharges":0,"commission":0,"optRegFee":0,"secFee":0},"transactionItem":{"accountId":123456789,"amount":8.88,"price":88.888,"cost":-88.888,"instruction":"BUY","instrument":{"symbol":"VTI","cusip":"123456789","assetType":"EQUITY"}}},{"type":"DIVIDEND_OR_INTEREST","subAccount":"1","settlementDate":"2018-10-04","netAmount":88.888,"transactionDate":"2018-10-04T19:14:41+0000","transactionSubType":"OD","transactionId":12345678918,"cashBalanceEffectFlag":true,"description":"ORDINARY DIVIDEND","fees":{"rFee":0,"additionalFee":0,"cdscFee":0,"regFee":0,"otherCharges":0,"commission":0,"optRegFee":0,"secFee":0},"transactionItem":{"accountId":123456789,"cost":0,"instrument":{"symbol":"VTI","cusip":"123456789","assetType":"EQUITY"}}},{"type":"JOURNAL","subAccount":"1","settlementDate":"2018-10-04","netAmount":-88.888,"transactionDate":"2018-10-04T19:14:40+0000","transactionSubType":"WF","transactionId":12345678962,"cashBalanceEffectFlag":true,"description":"W-8 WITHHOLDING","fees":{"rFee":0,"additionalFee":0,"cdscFee":0,"regFee":0,"otherCharges":0,"commission":0,"optRegFee":0,"secFee":0},"transactionItem":{"accountId":123456789,"cost":0,"instrument":{"symbol":"VTI","cusip":"123456789","assetType":"EQUITY"}}}]`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TransactionHistoryGetTransactionListCall
		want    *TransactionList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTransactionHistoryService(tdTest).GetTransactionList("accountID"), &TransactionList{
			ServerResponse: ServerResponse{200, http.Header{}},
			Transactions: []*Transaction{
				{
					Type:                  "TRADE",
					SubAccount:            "1",
					SettlementDate:        "2018-10-09",
					NetAmount:             -88.888,
					TransactionDate:       "2018-10-04T19:17:23+0000",
					TransactionSubType:    "DR",
					TransactionID:         12345678919,
					CashBalanceEffectFlag: true,
					Description:           "DIVIDEND REINVEST",
					Fees: &Fees{
						RFee:          0,
						AdditionalFee: 0,
						CDSCFee:       0,
						RegFee:        0,
						OtherCharges:  0,
						Commission:    0,
						OptRegFee:     0,
						SecFee:        0,
					},
					TransactionItem: &TransactionItem{
						AccountID:   123456789,
						Amount:      8.88,
						Price:       88.888,
						Cost:        -88.888,
						Instruction: "BUY",
						Instrument: &Instrument{
							Symbol:    "VTI",
							Cusip:     "123456789",
							AssetType: "EQUITY",
						},
					},
				},
				{
					Type:                  "DIVIDEND_OR_INTEREST",
					SubAccount:            "1",
					SettlementDate:        "2018-10-04",
					NetAmount:             88.888,
					TransactionDate:       "2018-10-04T19:14:41+0000",
					TransactionSubType:    "OD",
					TransactionID:         12345678918,
					CashBalanceEffectFlag: true,
					Description:           "ORDINARY DIVIDEND",
					Fees: &Fees{
						RFee:          0,
						AdditionalFee: 0,
						CDSCFee:       0,
						RegFee:        0,
						OtherCharges:  0,
						Commission:    0,
						OptRegFee:     0,
						SecFee:        0,
					},
					TransactionItem: &TransactionItem{
						AccountID: 123456789,
						Cost:      0,
						Instrument: &Instrument{
							Symbol:    "VTI",
							Cusip:     "123456789",
							AssetType: "EQUITY",
						},
					},
				},
				{
					Type:                  "JOURNAL",
					SubAccount:            "1",
					SettlementDate:        "2018-10-04",
					NetAmount:             -88.888,
					TransactionDate:       "2018-10-04T19:14:40+0000",
					TransactionSubType:    "WF",
					TransactionID:         12345678962,
					CashBalanceEffectFlag: true,
					Description:           "W-8 WITHHOLDING",
					Fees: &Fees{
						RFee:          0,
						AdditionalFee: 0,
						CDSCFee:       0,
						RegFee:        0,
						OtherCharges:  0,
						Commission:    0,
						OptRegFee:     0,
						SecFee:        0,
					},
					TransactionItem: &TransactionItem{
						AccountID: 123456789,
						Cost:      0,
						Instrument: &Instrument{
							Symbol:    "VTI",
							Cusip:     "123456789",
							AssetType: "EQUITY",
						},
					},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryGetTransactionListCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionHistoryGetTransactionListCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionHistoryGetTransactionCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TransactionHistoryGetTransactionCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewTransactionHistoryService(tdTest).GetTransaction("accountId", "transactionId"), "https://api.tdameritrade.com/v1/accounts/accountId/transactions/transactionId?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryGetTransactionCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionHistoryGetTransactionCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTransactionHistoryGetTransactionCall_Do(t *testing.T) {
	client := clientTest(`{"type":"TRADE","subAccount":"1","settlementDate":"2018-10-09","netAmount":-88.888,"transactionDate":"2018-10-04T19:17:23+0000","transactionSubType":"DR","transactionId":12345678919,"cashBalanceEffectFlag":true,"description":"DIVIDEND REINVEST","fees":{"rFee":0,"additionalFee":0,"cdscFee":0,"regFee":0,"otherCharges":0,"commission":0,"optRegFee":0,"secFee":0},"transactionItem":{"accountId":123456789,"amount":8.88,"price":88.888,"cost":-88.888,"instruction":"BUY","instrument":{"symbol":"VTI","cusip":"123456789","assetType":"EQUITY"}}}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TransactionHistoryGetTransactionCall
		want    *Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTransactionHistoryService(tdTest).GetTransaction("accountID", "12345678919"), &Transaction{
			ServerResponse:        ServerResponse{200, http.Header{}},
			Type:                  "TRADE",
			SubAccount:            "1",
			SettlementDate:        "2018-10-09",
			NetAmount:             -88.888,
			TransactionDate:       "2018-10-04T19:17:23+0000",
			TransactionSubType:    "DR",
			TransactionID:         12345678919,
			CashBalanceEffectFlag: true,
			Description:           "DIVIDEND REINVEST",
			Fees: &Fees{
				RFee:          0,
				AdditionalFee: 0,
				CDSCFee:       0,
				RegFee:        0,
				OtherCharges:  0,
				Commission:    0,
				OptRegFee:     0,
				SecFee:        0,
			},
			TransactionItem: &TransactionItem{
				AccountID:   123456789,
				Amount:      8.88,
				Price:       88.888,
				Cost:        -88.888,
				Instruction: "BUY",
				Instrument: &Instrument{
					Symbol:    "VTI",
					Cusip:     "123456789",
					AssetType: "EQUITY",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryGetTransactionCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionHistoryGetTransactionCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
