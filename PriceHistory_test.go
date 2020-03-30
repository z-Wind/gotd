package gotd

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestPriceHistoryGetPriceHistoryCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *PriceHistoryGetPriceHistoryCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewPriceHistoryService(tdTest).GetPriceHistory("symbol"), "https://api.tdameritrade.com/v1/marketdata/symbol/pricehistory?", false},
		{"All", NewPriceHistoryService(tdTest).GetPriceHistory("symbol").
			PeriodType(PriceHistoryPeriodTypeDay).
			Period(1).
			FrequencyType(PriceHistoryFrequencyTypeMinute).
			Frequency(2).
			EndDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)).
			StartDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)).
			NeedExtendedHoursData(false), "https://api.tdameritrade.com/v1/marketdata/symbol/pricehistory?endDate=1577808000000&frequency=2&frequencyType=minute&needExtendedHoursData=false&period=1&periodType=day&startDate=1577808000000", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("PriceHistoryGetPriceHistoryCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceHistoryGetPriceHistoryCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPriceHistoryGetPriceHistoryCall_Do(t *testing.T) {
	client := clientTest(`{"candles":[{"open":139,"high":139,"low":139,"close":139,"volume":912,"datetime":1234567890000},{"open":139.04,"high":139.31,"low":139.01,"close":139.21,"volume":107396,"datetime":1234567890000}],"symbol":"VTI","empty":false}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *PriceHistoryGetPriceHistoryCall
		want    *PriceHistory
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewPriceHistoryService(tdTest).GetPriceHistory("VTI"), &PriceHistory{
			ServerResponse: ServerResponse{200, http.Header{}},
			Candles: []Candle{
				{
					Open:     139,
					High:     139,
					Low:      139,
					Close:    139,
					Volume:   912,
					Datetime: Time(time.Unix(0, 1234567890000*int64(time.Millisecond))),
				},
				{
					Open:     139.04,
					High:     139.31,
					Low:      139.01,
					Close:    139.21,
					Volume:   107396,
					Datetime: Time(time.Unix(0, 1234567890000*int64(time.Millisecond))),
				},
			},
			Symbol: "VTI",
			Empty:  false,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("PriceHistoryGetPriceHistoryCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceHistoryGetPriceHistoryCall.Do() = %+v, want %v", got, tt.want)
			}
		})
	}
}
