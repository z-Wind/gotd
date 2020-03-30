package gotd

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestQuotesGetQuoteCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *QuotesGetQuoteCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewQuotesService(tdTest).GetQuote("symbol"), "https://api.tdameritrade.com/v1/marketdata/symbol/quotes?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetQuoteCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetQuoteCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestQuotesGetQuoteCall_Do(t *testing.T) {
	client := clientTest(`{"VTI":{"assetType":"ETF","symbol":"VTI","description":"Vanguard Total Stock Market ETF","bidPrice":139.2,"bidSize":100,"bidId":"P","askPrice":139.66,"askSize":700,"askId":"P","lastPrice":139.56,"lastSize":11000,"lastId":"P","openPrice":140.93,"highPrice":141.2,"lowPrice":138.39,"bidTick":" ","closePrice":139.56,"netChange":0,"totalVolume":3171400,"quoteTimeInLong":1234567896878,"tradeTimeInLong":1234567890004,"mark":139.56,"exchange":"p","exchangeName":"Pacific","marginable":true,"shortable":true,"volatility":0.123456789,"digits":4,"52WkHigh":151.839,"52WkLow":129.841,"nAV":0,"peRatio":0,"divAmount":2.5567,"divYield":1.83,"divDate":"2018-09-28 00:00:00.0","securityStatus":"Closed","regularMarketLastPrice":139.56,"regularMarketLastSize":110,"regularMarketNetChange":0,"regularMarketTradeTimeInLong":1234567890004,"delayed":false}}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *QuotesGetQuoteCall
		want    *Quote
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewQuotesService(tdTest).GetQuote("VTI"), &Quote{
			ServerResponse: ServerResponse{200, http.Header{}},

			AssetType:                    "ETF",
			Symbol:                       "VTI",
			Description:                  "Vanguard Total Stock Market ETF",
			BidPrice:                     139.2,
			BidSize:                      100,
			BidID:                        "P",
			AskPrice:                     139.66,
			AskSize:                      700,
			AskID:                        "P",
			LastPrice:                    139.56,
			LastSize:                     11000,
			LastID:                       "P",
			OpenPrice:                    140.93,
			HighPrice:                    141.2,
			LowPrice:                     138.39,
			BidTick:                      " ",
			ClosePrice:                   139.56,
			NetChange:                    0,
			TotalVolume:                  3171400,
			QuoteTimeInLong:              1234567896878,
			TradeTimeInLong:              1234567890004,
			Mark:                         139.56,
			Exchange:                     "p",
			ExchangeName:                 "Pacific",
			Marginable:                   true,
			Shortable:                    true,
			Volatility:                   0.123456789,
			Digits:                       4,
			Wk52High:                     151.839,
			Wk52Low:                      129.841,
			NAV:                          0,
			PeRatio:                      0,
			DivAmount:                    2.5567,
			DivYield:                     1.83,
			DivDate:                      "2018-09-28 00:00:00.0",
			SecurityStatus:               "Closed",
			RegularMarketLastPrice:       139.56,
			RegularMarketLastSize:        110,
			RegularMarketNetChange:       0,
			RegularMarketTradeTimeInLong: 1234567890004,
			Delayed:                      false,
		}, false},
		{"Real fail", NewQuotesService(tdReal).GetQuote("0050"), nil, true},
	}
	for _, tt := range tests {
		if strings.Contains(tt.name, "Real") && !onlineTest {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetQuoteCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetQuoteCall.Do() = %+v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuotesGetQuoteListCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *QuotesGetQuoteListCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewQuotesService(tdTest).GetQuoteList("VTI", "VBR"), "https://api.tdameritrade.com/v1/marketdata/quotes?symbol=VTI%2CVBR", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetQuoteListCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetQuoteListCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuotesGetQuoteListCall_Do(t *testing.T) {
	client := clientTest(`{"BND":{"assetType":"ETF","symbol":"BND","description":"Vanguard Total Bond Market ETF","bidPrice":77.01,"bidSize":2000,"bidId":"K","askPrice":77.89,"askSize":2000,"askId":"Q","lastPrice":78.1,"lastSize":0,"lastId":"Q","openPrice":77.69,"highPrice":77.71,"lowPrice":77.4601,"bidTick":" ","closePrice":77.5,"netChange":0.6,"totalVolume":2024625,"quoteTimeInLong":1234567890395,"tradeTimeInLong":1234567897119,"mark":77.5,"exchange":"q","exchangeName":"NASDAQ","marginable":true,"shortable":true,"volatility":0.12345678,"digits":4,"52WkHigh":81.95,"52WkLow":77.4601,"nAV":0,"peRatio":0,"divAmount":2.1404,"divYield":2.76,"divDate":"2018-11-01 00:00:00.0","securityStatus":"Normal","regularMarketLastPrice":77.5,"regularMarketLastSize":102,"regularMarketNetChange":0,"regularMarketTradeTimeInLong":1234567890246,"delayed":false},"VTI":{"assetType":"ETF","symbol":"VTI","description":"Vanguard Total Stock Market ETF","bidPrice":139.2,"bidSize":100,"bidId":"P","askPrice":139.66,"askSize":700,"askId":"P","lastPrice":139.56,"lastSize":11000,"lastId":"P","openPrice":140.93,"highPrice":141.2,"lowPrice":138.39,"bidTick":" ","closePrice":139.56,"netChange":0,"totalVolume":3171400,"quoteTimeInLong":1234567896878,"tradeTimeInLong":1234567890004,"mark":139.56,"exchange":"p","exchangeName":"Pacific","marginable":true,"shortable":true,"volatility":0.123456789,"digits":4,"52WkHigh":151.839,"52WkLow":129.841,"nAV":0,"peRatio":0,"divAmount":2.5567,"divYield":1.83,"divDate":"2018-09-28 00:00:00.0","securityStatus":"Closed","regularMarketLastPrice":139.56,"regularMarketLastSize":110,"regularMarketNetChange":0,"regularMarketTradeTimeInLong":1234567890004,"delayed":false}}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *QuotesGetQuoteListCall
		want    *QuoteMap
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewQuotesService(tdTest).GetQuoteList("BND", "VTI"), &QuoteMap{
			ServerResponse: ServerResponse{200, http.Header{}},
			Quotes: map[string]*Quote{
				"BND": &Quote{
					AssetType:                    "ETF",
					Symbol:                       "BND",
					Description:                  "Vanguard Total Bond Market ETF",
					BidPrice:                     77.01,
					BidSize:                      2000,
					BidID:                        "K",
					AskPrice:                     77.89,
					AskSize:                      2000,
					AskID:                        "Q",
					LastPrice:                    78.1,
					LastSize:                     0,
					LastID:                       "Q",
					OpenPrice:                    77.69,
					HighPrice:                    77.71,
					LowPrice:                     77.4601,
					BidTick:                      " ",
					ClosePrice:                   77.5,
					NetChange:                    0.6,
					TotalVolume:                  2024625,
					QuoteTimeInLong:              1234567890395,
					TradeTimeInLong:              1234567897119,
					Mark:                         77.5,
					Exchange:                     "q",
					ExchangeName:                 "NASDAQ",
					Marginable:                   true,
					Shortable:                    true,
					Volatility:                   0.12345678,
					Digits:                       4,
					Wk52High:                     81.95,
					Wk52Low:                      77.4601,
					NAV:                          0,
					PeRatio:                      0,
					DivAmount:                    2.1404,
					DivYield:                     2.76,
					DivDate:                      "2018-11-01 00:00:00.0",
					SecurityStatus:               "Normal",
					RegularMarketLastPrice:       77.5,
					RegularMarketLastSize:        102,
					RegularMarketNetChange:       0,
					RegularMarketTradeTimeInLong: 1234567890246,
					Delayed:                      false,
				},
				"VTI": &Quote{
					AssetType:                    "ETF",
					Symbol:                       "VTI",
					Description:                  "Vanguard Total Stock Market ETF",
					BidPrice:                     139.2,
					BidSize:                      100,
					BidID:                        "P",
					AskPrice:                     139.66,
					AskSize:                      700,
					AskID:                        "P",
					LastPrice:                    139.56,
					LastSize:                     11000,
					LastID:                       "P",
					OpenPrice:                    140.93,
					HighPrice:                    141.2,
					LowPrice:                     138.39,
					BidTick:                      " ",
					ClosePrice:                   139.56,
					NetChange:                    0,
					TotalVolume:                  3171400,
					QuoteTimeInLong:              1234567896878,
					TradeTimeInLong:              1234567890004,
					Mark:                         139.56,
					Exchange:                     "p",
					ExchangeName:                 "Pacific",
					Marginable:                   true,
					Shortable:                    true,
					Volatility:                   0.123456789,
					Digits:                       4,
					Wk52High:                     151.839,
					Wk52Low:                      129.841,
					NAV:                          0,
					PeRatio:                      0,
					DivAmount:                    2.5567,
					DivYield:                     1.83,
					DivDate:                      "2018-09-28 00:00:00.0",
					SecurityStatus:               "Closed",
					RegularMarketLastPrice:       139.56,
					RegularMarketLastSize:        110,
					RegularMarketNetChange:       0,
					RegularMarketTradeTimeInLong: 1234567890004,
					Delayed:                      false,
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetQuoteListCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetQuoteListCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
