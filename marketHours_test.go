package gotd

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestMarketHoursGetHoursForMultipleMarketsCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *MarketHoursGetHoursForMultipleMarketsCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewMarketHoursService(tdTest).GetHoursForMultipleMarkets(HoursMarketsBOND, HoursMarketsEQUITY, HoursMarketsFOREX, HoursMarketsFUTURE, HoursMarketsOPTION), "https://api.tdameritrade.com/v1/marketdata/hours?markets=BOND%2CEQUITY%2CFOREX%2CFUTURE%2COPTION", false},
		{"date", NewMarketHoursService(tdTest).GetHoursForMultipleMarkets(HoursMarketsBOND, HoursMarketsEQUITY, HoursMarketsFOREX, HoursMarketsFUTURE, HoursMarketsOPTION).Date(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "https://api.tdameritrade.com/v1/marketdata/hours?date=2020-01-01&markets=BOND%2CEQUITY%2CFOREX%2CFUTURE%2COPTION", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketHoursGetHoursForMultipleMarketsCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketHoursGetHoursForMultipleMarketsCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMarketHoursGetHoursForMultipleMarketsCall_Do(t *testing.T) {
	client := clientTest(`{"bond":{"bond":{"date":"2018-11-03","marketType":"BOND","product":"bond","isOpen":false}},"equity":{"equity":{"date":"2018-11-03","marketType":"EQUITY","product":"equity","isOpen":false}},"future":{"CTW":{"date":"2018-11-03","marketType":"FUTURE","exchange":"ICE","category":"Agriculture","product":"CTW","productName":"weekly cotton no. 2 options","isOpen":true,"sessionHours":{"preMarket":[{"start":"2018-11-03T19:30:00-04:00","end":"2018-11-03T21:00:00-04:00"}],"regularMarket":[{"start":"2018-11-03T21:00:00-04:00","end":"2018-11-04T14:20:00-05:00"}]}},"PS":{"date":"2018-11-03","marketType":"FUTURE","exchange":"ICE","category":"FX","product":"PS","productName":"cross currency pairs british pound sterling swedish krona futures","isOpen":true,"sessionHours":{"preMarket":[{"start":"2018-11-03T19:30:00-04:00","end":"2018-11-03T20:00:00-04:00"}],"regularMarket":[{"start":"2018-11-03T20:00:00-04:00","end":"2018-11-04T17:00:00-05:00"}]}}},"option":{"option":{"date":"2018-11-03","marketType":"OPTION","product":"option","isOpen":false}}}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *MarketHoursGetHoursForMultipleMarketsCall
		want    *MarketHourMap
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewMarketHoursService(tdTest).GetHoursForMultipleMarkets(), &MarketHourMap{
			ServerResponse: ServerResponse{200, http.Header{}},
			MarketHourProductMaps: map[string]map[string]MarketHour{
				"bond": map[string]MarketHour{
					"bond": MarketHour{
						Date:       "2018-11-03",
						MarketType: "BOND",
						Product:    "bond",
						IsOpen:     false,
					},
				},
				"equity": map[string]MarketHour{
					"equity": MarketHour{
						Date:       "2018-11-03",
						MarketType: "EQUITY",
						Product:    "equity",
						IsOpen:     false,
					},
				},
				"future": map[string]MarketHour{
					"CTW": MarketHour{
						Date:        "2018-11-03",
						MarketType:  "FUTURE",
						Exchange:    "ICE",
						Category:    "Agriculture",
						Product:     "CTW",
						ProductName: "weekly cotton no. 2 options",
						IsOpen:      true,
						SessionHours: SessionHours{
							PreMarket: []Period{
								Period{
									Start: "2018-11-03T19:30:00-04:00",
									End:   "2018-11-03T21:00:00-04:00",
								},
							},
							RegularMarket: []Period{
								Period{
									Start: "2018-11-03T21:00:00-04:00",
									End:   "2018-11-04T14:20:00-05:00",
								},
							},
						},
					},
					"PS": MarketHour{
						Date:        "2018-11-03",
						MarketType:  "FUTURE",
						Exchange:    "ICE",
						Category:    "FX",
						Product:     "PS",
						ProductName: "cross currency pairs british pound sterling swedish krona futures",
						IsOpen:      true,
						SessionHours: SessionHours{
							PreMarket: []Period{
								Period{
									Start: "2018-11-03T19:30:00-04:00",
									End:   "2018-11-03T20:00:00-04:00",
								},
							},
							RegularMarket: []Period{
								Period{
									Start: "2018-11-03T20:00:00-04:00",
									End:   "2018-11-04T17:00:00-05:00",
								},
							},
						},
					},
				},
				"option": map[string]MarketHour{
					"option": MarketHour{
						Date:       "2018-11-03",
						MarketType: "OPTION",
						Product:    "option",
						IsOpen:     false,
					},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketHoursGetHoursForMultipleMarketsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketHoursGetHoursForMultipleMarketsCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketHoursGetHoursForASingleMarketCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *MarketHoursGetHoursForASingleMarketCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewMarketHoursService(tdTest).GetHoursForASingleMarket(HoursMarketsBOND), "https://api.tdameritrade.com/v1/marketdata/BOND/hours?", false},
		{"date", NewMarketHoursService(tdTest).GetHoursForASingleMarket(HoursMarketsBOND).Date(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "https://api.tdameritrade.com/v1/marketdata/BOND/hours?date=2020-01-01", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketHoursGetHoursForASingleMarketCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketHoursGetHoursForASingleMarketCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketHoursGetHoursForASingleMarketCall_Do(t *testing.T) {
	client := clientTest(`{"equity":{"equity":{"date":"2018-11-03","marketType":"EQUITY","product":"equity","isOpen":false}}}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *MarketHoursGetHoursForASingleMarketCall
		want    *MarketHourProductMap
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewMarketHoursService(tdTest).GetHoursForASingleMarket(HoursMarketsEQUITY), &MarketHourProductMap{
			ServerResponse: ServerResponse{200, http.Header{}},
			MarketHours: map[string]MarketHour{
				"equity": MarketHour{
					Date:       "2018-11-03",
					MarketType: "EQUITY",
					Product:    "equity",
					IsOpen:     false,
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketHoursGetHoursForASingleMarketCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketHoursGetHoursForASingleMarketCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
