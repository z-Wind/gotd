package gotd

import (
	"testing"
)

func TestAccountsGetAccountCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *AccountsGetAccountCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewAccountsService(td).GetAccount(accountID), accountID, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountsGetAccountCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.SecuritiesAccount.AccountID != tt.want {
				t.Errorf("AccountsGetAccountCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountsGetAccountListCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *AccountsGetAccountListCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewAccountsService(td).GetAccountList(), accountID, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountsGetAccountListCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Accounts[0].SecuritiesAccount.AccountID != tt.want {
				t.Errorf("AccountsGetAccountListCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
