package gotd

import (
	"reflect"
	"testing"
)

func TestMarketHoursGetHoursForMultipleMarketsCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *MarketHoursGetHoursForMultipleMarketsCall
		want    map[string]struct{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewMarketHoursService(td).GetHoursForMultipleMarkets(), map[string]struct{}{HoursMarketsFUTURE: struct{}{}, HoursMarketsEQUITY: struct{}{}, HoursMarketsBOND: struct{}{}, HoursMarketsFOREX: struct{}{}, HoursMarketsOPTION: struct{}{}}, false},
		{"Test", NewMarketHoursService(td).GetHoursForMultipleMarkets(HoursMarketsFUTURE, HoursMarketsEQUITY, HoursMarketsBOND), map[string]struct{}{HoursMarketsFUTURE: struct{}{}, HoursMarketsEQUITY: struct{}{}, HoursMarketsBOND: struct{}{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketHoursGetHoursForMultipleMarketsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ans := map[string]struct{}{}
			for _, prods := range got.MarketHourProductMaps {
				for _, prod := range prods {
					ans[prod.MarketType] = struct{}{}
					break
				}
			}
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("MarketHoursGetHoursForMultipleMarketsCall.Do() = %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMarketHoursGetHoursForASingleMarketCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *MarketHoursGetHoursForASingleMarketCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewMarketHoursService(td).GetHoursForASingleMarket(HoursMarketsBOND), HoursMarketsBOND, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketHoursGetHoursForASingleMarketCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var ans string
			for _, prod := range got.MarketHours {
				ans = prod.MarketType
				break
			}
			if ans != tt.want {
				t.Errorf("MarketHoursGetHoursForASingleMarketCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
