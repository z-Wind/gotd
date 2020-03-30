package gotd

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMoversGetMoversCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *MoversGetMoversCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewMoversService(tdTest).GetMovers("index"), "https://api.tdameritrade.com/v1/marketdata/index/movers?", false},
		{"direction", NewMoversService(tdTest).GetMovers("index").Direction(MoversDirectionUp), "https://api.tdameritrade.com/v1/marketdata/index/movers?direction=up", false},
		{"direction", NewMoversService(tdTest).GetMovers("index").Direction(MoversDirectionDown), "https://api.tdameritrade.com/v1/marketdata/index/movers?direction=down", false},
		{"change", NewMoversService(tdTest).GetMovers("index").Change(MoversChangePercent), "https://api.tdameritrade.com/v1/marketdata/index/movers?change=percent", false},
		{"change", NewMoversService(tdTest).GetMovers("index").Change(MoversChangeValue), "https://api.tdameritrade.com/v1/marketdata/index/movers?change=value", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("MoversGetMoversCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoversGetMoversCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 假日無法測試
func TestMoversGetMoversCall_Do(t *testing.T) {
	client := clientTest(`{"index":{"change":0,"description":"string","direction":"'up' or 'down'","last":0,"symbol":"string","totalVolume":0}}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *MoversGetMoversCall
		want    *Mover
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewMoversService(tdTest).GetMovers("index"), &Mover{
			ServerResponse: ServerResponse{200, http.Header{}},
			Change:         0,
			Description:    "string",
			Direction:      "'up' or 'down'",
			Last:           0,
			Symbol:         "string",
			TotalVolume:    0,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("MoversGetMoversCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoversGetMoversCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
