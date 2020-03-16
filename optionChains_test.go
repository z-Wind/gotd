package gotd

import (
	"testing"
)

func TestOptionChainsGetOptionChainCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *OptionChainsGetOptionChainCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewOptionChainsService(td).GetOptionChain("VTI"), "VTI", false},
		{"SINGLE ", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategySINGLE), "VTI", false},
		{"ANALYTICAL ", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyANALYTICAL), "VTI", false},
		{"COVERED", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyCOVERED), "VTI", false},
		{"VERTICAL", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyVERTICAL), "VTI", false},
		{"CALENDAR", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyCALENDAR), "VTI", false},
		{"STRANGLE", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategySTRANGLE), "VTI", false},
		{"STRADDLE", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategySTRADDLE), "VTI", false},
		// {"BUTTERFLY", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyBUTTERFLY), "VTI", false},
		// {"CONDOR", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyCONDOR), "VTI", false},
		// {"DIAGONAL", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyDIAGONAL), "VTI", false},
		// {"COLLAR", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyCOLLAR), "VTI", false},
		// {"ROLL", NewOptionChainsService(td).GetOptionChain("VTI").Strategy(OptionChainStrategyROLL), "VTI", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("OptionChainsGetOptionChainCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Symbol != tt.want {
				t.Errorf("OptionChainsGetOptionChainCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
