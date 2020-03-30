package gotd

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestAccountsGetAccountCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *AccountsGetAccountCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewAccountsService(tdTest).GetAccount("accountId"), "https://api.tdameritrade.com/v1/accounts/accountId?", false},
		{"Positions", NewAccountsService(tdTest).GetAccount("accountId").Fields(AccountFieldPositions), "https://api.tdameritrade.com/v1/accounts/accountId?fields=positions", false},
		{"Orders", NewAccountsService(tdTest).GetAccount("accountId").Fields(AccountFieldOrders), "https://api.tdameritrade.com/v1/accounts/accountId?fields=orders", false},
		{"Positions,Orders", NewAccountsService(tdTest).GetAccount("accountId").Fields(AccountFieldPositions + "," + AccountFieldOrders), "https://api.tdameritrade.com/v1/accounts/accountId?fields=positions%2Corders", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountsGetAccountCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountsGetAccountCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountsGetAccountCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *AccountsGetAccountCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Real", NewAccountsService(tdReal).GetAccount(accountID), accountID, false},
		{"Test", func() *AccountsGetAccountCall {
			client := clientTest(`{"securitiesAccount":{"type":"CASH","accountId":"123456789","roundTrips":0,"isDayTrader":false,"isClosingOnlyRestricted":false,"positions":[{"shortQuantity":0,"averagePrice":88.888,"currentDayProfitLoss":88888.88,"currentDayProfitLossPercentage":8.88,"longQuantity":88.888,"settledLongQuantity":88.888,"settledShortQuantity":0,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"marketValue":88.888},{"shortQuantity":0,"averagePrice":88.888,"currentDayProfitLoss":-88.888,"currentDayProfitLossPercentage":-8.88,"longQuantity":88.888,"settledLongQuantity":88.888,"settledShortQuantity":0,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"BND"},"marketValue":888.888}],"orderStrategies":[{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T07:30:23+0000","closeTime":"2018-11-03T07:30:23+0000","tag":"WEB_GRID_SNAP","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."},{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:25:41+0000","closeTime":"2018-11-03T08:25:41+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."}],"initialBalances":{"accruedInterest":0,"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88,"cashBalance":8.88,"bondValue":0,"cashReceipts":0,"liquidationValue":88888.88,"longOptionMarketValue":0,"longStockValue":88888.88,"moneyMarketFund":0,"mutualFundValue":0,"shortOptionMarketValue":0,"shortStockValue":0,"isInCall":false,"unsettledCash":0,"cashDebitCallValue":0,"pendingDeposits":0,"accountValue":88888.88},"currentBalances":{"accruedInterest":0,"cashBalance":8.88,"cashReceipts":0,"longOptionMarketValue":0,"liquidationValue":88888.88,"longMarketValue":88888.88,"moneyMarketFund":0,"savings":0,"shortMarketValue":0,"pendingDeposits":0,"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88,"cashCall":0,"longNonMarginableMarketValue":8.88,"totalCash":8.88,"shortOptionMarketValue":0,"bondValue":0,"cashDebitCallValue":0,"unsettledCash":0},"projectedBalances":{"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88}}}`, http.StatusOK)
			tdTest, _ := New(client)
			return NewAccountsService(tdTest).GetAccount(accountID)
		}(), "123456789", false},
	}
	for _, tt := range tests {
		if strings.Contains(tt.name, "Real") && !onlineTest {
			continue
		}
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

func TestAccountsGetAccountListCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *AccountsGetAccountListCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewAccountsService(tdTest).GetAccountList(), "https://api.tdameritrade.com/v1/accounts?", false},
		{"Positions", NewAccountsService(tdTest).GetAccountList().Fields(AccountFieldPositions), "https://api.tdameritrade.com/v1/accounts?fields=positions", false},
		{"Orders", NewAccountsService(tdTest).GetAccountList().Fields(AccountFieldOrders), "https://api.tdameritrade.com/v1/accounts?fields=orders", false},
		{"Positions,Orders", NewAccountsService(tdTest).GetAccountList().Fields(AccountFieldPositions + "," + AccountFieldOrders), "https://api.tdameritrade.com/v1/accounts?fields=positions%2Corders", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountsGetAccountListCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountsGetAccountListCall.doRequest() = %v, want %v", got, tt.want)
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
		{"Real", NewAccountsService(tdReal).GetAccountList(), accountID, false},
		{"Test", func() *AccountsGetAccountListCall {
			client := clientTest(`[{"securitiesAccount":{"type":"CASH","accountId":"123456789","roundTrips":0,"isDayTrader":false,"isClosingOnlyRestricted":false,"positions":[{"shortQuantity":0,"averagePrice":88.888,"currentDayProfitLoss":88.888,"currentDayProfitLossPercentage":88.888,"longQuantity":88.888,"settledLongQuantity":88.888,"settledShortQuantity":0,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"marketValue":88.888},{"shortQuantity":0,"averagePrice":88.888,"currentDayProfitLoss":-88.888,"currentDayProfitLossPercentage":-8.88,"longQuantity":88.888,"settledLongQuantity":88.888,"settledShortQuantity":0,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"BND"},"marketValue":88.888}],"orderStrategies":[{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T07:30:23+0000","closeTime":"2018-11-03T07:30:23+0000","tag":"WEB_GRID_SNAP","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."},{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:25:41+0000","closeTime":"2018-11-03T08:25:41+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."}],"initialBalances":{"accruedInterest":0,"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88,"cashBalance":8.88,"bondValue":0,"cashReceipts":0,"liquidationValue":88888.88,"longOptionMarketValue":0,"longStockValue":88888.88,"moneyMarketFund":0,"mutualFundValue":0,"shortOptionMarketValue":0,"shortStockValue":0,"isInCall":false,"unsettledCash":0,"cashDebitCallValue":0,"pendingDeposits":0,"accountValue":88888.88},"currentBalances":{"accruedInterest":0,"cashBalance":8.88,"cashReceipts":0,"longOptionMarketValue":0,"liquidationValue":88888.88,"longMarketValue":88888.88,"moneyMarketFund":0,"savings":0,"shortMarketValue":0,"pendingDeposits":0,"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88,"cashCall":0,"longNonMarginableMarketValue":8.88,"totalCash":8.88,"shortOptionMarketValue":0,"bondValue":0,"cashDebitCallValue":0,"unsettledCash":0},"projectedBalances":{"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88}}},{"securitiesAccount":{"type":"CASH","accountId":"123456789","roundTrips":0,"isDayTrader":false,"isClosingOnlyRestricted":false,"positions":[{"shortQuantity":0,"averagePrice":88.888,"currentDayProfitLoss":88.888,"currentDayProfitLossPercentage":88.888,"longQuantity":88.888,"settledLongQuantity":88.888,"settledShortQuantity":0,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"marketValue":88.888},{"shortQuantity":0,"averagePrice":88.888,"currentDayProfitLoss":-88.888,"currentDayProfitLossPercentage":-8.88,"longQuantity":88.888,"settledLongQuantity":88.888,"settledShortQuantity":0,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"BND"},"marketValue":88.888}],"orderStrategies":[{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T07:30:23+0000","closeTime":"2018-11-03T07:30:23+0000","tag":"WEB_GRID_SNAP","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."},{"session":"NORMAL","duration":"GOOD_TILL_CANCEL","orderType":"LIMIT","cancelTime":"2019-03-01","complexOrderStrategyType":"NONE","quantity":1,"filledQuantity":0,"remainingQuantity":0,"requestedDestination":"AUTO","destinationLinkName":"AutoRoute","price":1,"orderLegCollection":[{"orderLegType":"EQUITY","legId":1,"instrument":{"assetType":"EQUITY","cusip":"123456789","symbol":"VTI"},"instruction":"BUY","positionEffect":"OPENING","quantity":1}],"orderStrategyType":"SINGLE","orderId":123456789,"cancelable":false,"editable":false,"status":"REJECTED","enteredTime":"2018-11-03T08:25:41+0000","closeTime":"2018-11-03T08:25:41+0000","tag":"AA_ABCDEFGHIJK","accountId":123456789,"statusDescription":"Your limit price is too distant from the current quote. Try again."}],"initialBalances":{"accruedInterest":0,"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88,"cashBalance":8.88,"bondValue":0,"cashReceipts":0,"liquidationValue":88888.88,"longOptionMarketValue":0,"longStockValue":88888.88,"moneyMarketFund":0,"mutualFundValue":0,"shortOptionMarketValue":0,"shortStockValue":0,"isInCall":false,"unsettledCash":0,"cashDebitCallValue":0,"pendingDeposits":0,"accountValue":88888.88},"currentBalances":{"accruedInterest":0,"cashBalance":8.88,"cashReceipts":0,"longOptionMarketValue":0,"liquidationValue":88888.88,"longMarketValue":88888.88,"moneyMarketFund":0,"savings":0,"shortMarketValue":0,"pendingDeposits":0,"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88,"cashCall":0,"longNonMarginableMarketValue":8.88,"totalCash":8.88,"shortOptionMarketValue":0,"bondValue":0,"cashDebitCallValue":0,"unsettledCash":0},"projectedBalances":{"cashAvailableForTrading":8.88,"cashAvailableForWithdrawal":8.88}}}]`, http.StatusOK)
			tdTest, _ := New(client)
			return NewAccountsService(tdTest).GetAccountList()
		}(), "123456789", false},
	}
	for _, tt := range tests {
		if strings.Contains(tt.name, "Real") && !onlineTest {
			continue
		}
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
