package gotd

import (
	"reflect"
	"testing"
)

func TestTransactionHistoryGetTransactionListCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TransactionHistoryGetTransactionListCall
		want    *QuoteMap
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTransactionHistoryService(td).GetTransactionList(accountID), nil, false},
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

func TestTransactionHistoryGetTransactionCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TransactionHistoryGetTransactionCall
		want    *Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTransactionHistoryService(td).GetTransaction(accountID, "123"), nil, true},
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
