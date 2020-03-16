package gotd

import (
	"reflect"
	"testing"
)

func TestInstrumentsSearchInstrumentsCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *InstrumentsSearchInstrumentsCall
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SymbolSearch", NewInstrumentsService(td).SearchInstruments("VTI", InstrumentsProjectionSymbolSearch), []string{"VTI"}, false},
		{"SymbolSearch", NewInstrumentsService(td).SearchInstruments("VTI,BND", InstrumentsProjectionSymbolSearch), []string{"VTI", "BND"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("InstrumentsSearchInstrumentsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ans := []string{}
			for key := range got.Instruments {
				ans = append(ans, key)
			}
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("InstrumentsSearchInstrumentsCall.Do() = %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestInstrumentsGetInstrumentCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *InstrumentsGetInstrumentCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewInstrumentsService(td).GetInstrument("921937835"), "BND", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("InstrumentsGetInstrumentCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Symbol != tt.want {
				t.Errorf("InstrumentsGetInstrumentCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
