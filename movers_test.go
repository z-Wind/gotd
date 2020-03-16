package gotd

import (
	"testing"
)

// 假日無法測試
func TestMoversGetMoversCall_Do(t *testing.T) {
	t.Skip("假日無法測試")

	tests := []struct {
		name    string
		c       *MoversGetMoversCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewMoversService(td).GetMovers(MoversIndexCOMPX), MoversIndexCOMPX, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("MoversGetMoversCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Symbol != tt.want {
				t.Errorf("MoversGetMoversCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
