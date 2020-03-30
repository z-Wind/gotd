package gotd

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestInstrumentsSearchInstrumentsCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *InstrumentsSearchInstrumentsCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewInstrumentsService(tdTest).SearchInstruments("symbol", "projection"), "https://api.tdameritrade.com/v1/instruments?projection=projection&symbol=symbol", false},
		{"Projection", NewInstrumentsService(tdTest).SearchInstruments("symbol", InstrumentsProjectionDescRegex), "https://api.tdameritrade.com/v1/instruments?projection=desc-regex&symbol=symbol", false},
		{"Projection", NewInstrumentsService(tdTest).SearchInstruments("symbol", InstrumentsProjectionDescSearch), "https://api.tdameritrade.com/v1/instruments?projection=desc-search&symbol=symbol", false},
		{"Projection", NewInstrumentsService(tdTest).SearchInstruments("symbol", InstrumentsProjectionFundamental), "https://api.tdameritrade.com/v1/instruments?projection=fundamental&symbol=symbol", false},
		{"Projection", NewInstrumentsService(tdTest).SearchInstruments("symbol", InstrumentsProjectionSymbolRegex), "https://api.tdameritrade.com/v1/instruments?projection=symbol-regex&symbol=symbol", false},
		{"Projection", NewInstrumentsService(tdTest).SearchInstruments("symbol", InstrumentsProjectionSymbolSearch), "https://api.tdameritrade.com/v1/instruments?projection=symbol-search&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("InstrumentsSearchInstrumentsCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstrumentsSearchInstrumentsCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInstrumentsSearchInstrumentsCall_Do(t *testing.T) {
	client := clientTest(`{"VTI":{"cusip":"123456789","symbol":"VTI","description":"Vanguard Total Stock Market ETF","exchange":"Pacific","assetType":"ETF"}}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *InstrumentsSearchInstrumentsCall
		want    *InstrumentMap
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewInstrumentsService(tdTest).SearchInstruments("symbol", InstrumentsProjectionSymbolSearch), &InstrumentMap{
			ServerResponse: ServerResponse{200, http.Header{}},
			Instruments: map[string]*Instrument{
				"VTI": &Instrument{
					Cusip:       "123456789",
					Symbol:      "VTI",
					Description: "Vanguard Total Stock Market ETF",
					Exchange:    "Pacific",
					AssetType:   "ETF",
				},
			},
		}, false},
		{"Real", NewInstrumentsService(tdReal).SearchInstruments("VTI", InstrumentsProjectionSymbolSearch), &InstrumentMap{
			ServerResponse: ServerResponse{200, http.Header{}},
			Instruments: map[string]*Instrument{
				"VTI": &Instrument{
					Cusip:       "922908769",
					Symbol:      "VTI",
					Description: "Vanguard Total Stock Market ETF",
					Exchange:    "Pacific",
					AssetType:   "ETF",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		if strings.Contains(tt.name, "Real") && !onlineTest {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("InstrumentsSearchInstrumentsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.ServerResponse = got.ServerResponse
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstrumentsSearchInstrumentsCall.Do() = %v, want %v", got.Instruments, tt.want.Instruments)
			}
		})
	}
}

func TestInstrumentsGetInstrumentCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *InstrumentsGetInstrumentCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewInstrumentsService(tdTest).GetInstrument("cursip"), "https://api.tdameritrade.com/v1/instruments/cursip?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("InstrumentsGetInstrumentCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstrumentsGetInstrumentCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstrumentsGetInstrumentCall_Do(t *testing.T) {
	client := clientTest(`[{"cusip":"123456789","symbol":"VTI","description":"Vanguard Total Stock Market ETF","exchange":"Pacific","assetType":"ETF"}]`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *InstrumentsGetInstrumentCall
		want    *Instrument
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewInstrumentsService(tdTest).GetInstrument("123456789"), &Instrument{
			ServerResponse: ServerResponse{200, http.Header{}},
			Cusip:          "123456789",
			Symbol:         "VTI",
			Description:    "Vanguard Total Stock Market ETF",
			Exchange:       "Pacific",
			AssetType:      "ETF",
		}, false},
		{"Real", NewInstrumentsService(tdReal).GetInstrument("921937835"), &Instrument{
			ServerResponse: ServerResponse{200, http.Header{}},
			Cusip:          "921937835",
			Symbol:         "BND",
			Description:    "Vanguard Total Bond Market ETF",
			Exchange:       "NASDAQ",
			AssetType:      "ETF",
		}, false},
	}
	for _, tt := range tests {
		if strings.Contains(tt.name, "Real") && !onlineTest {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("InstrumentsGetInstrumentCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.ServerResponse = got.ServerResponse
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstrumentsGetInstrumentCall.Do() = %+v\n, want \n%+v", got, tt.want)
			}
		})
	}
}
