package gotd

import (
	"reflect"
	"testing"
)

func TestQuotesGetQuoteCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *QuotesGetQuoteCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewQuotesService(td).GetQuote("VTI"), "VTI", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetQuoteCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Symbol != tt.want {
				t.Errorf("QuotesGetQuoteCall.Do() = %+v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuotesGetQuoteListCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *QuotesGetQuoteListCall
		want    map[string]struct{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewQuotesService(td).GetQuoteList("VTI", "VBR"), map[string]struct{}{"VTI": struct{}{}, "VBR": struct{}{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetQuoteListCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ans := map[string]struct{}{}
			for _, q := range got.Quotes {
				ans[q.Symbol] = struct{}{}
			}
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("QuotesGetQuoteListCall.Do() = %v, want %v", ans, tt.want)
			}
		})
	}
}
