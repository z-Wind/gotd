package gotd

import (
	"testing"
)

func TestPriceHistoryGetPriceHistoryCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *PriceHistoryGetPriceHistoryCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewPriceHistoryService(td).GetPriceHistory("VTI"), "VTI", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("PriceHistoryGetPriceHistoryCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Symbol != tt.want {
				t.Errorf("PriceHistoryGetPriceHistoryCall.Do() = %+v, want %v", got, tt.want)
			}
		})
	}
}
