package gotd

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWatchlistGetWatchlistsForMultipleAccountsCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistsForMultipleAccountsCall
		want    *WatchlistList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewWatchlistService(td).GetWatchlistsForMultipleAccounts(), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistGetWatchlistsForMultipleAccountsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistGetWatchlistsForMultipleAccountsCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistGetWatchlistsForSingleAccountCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistsForSingleAccountCall
		want    *WatchlistList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewWatchlistService(td).GetWatchlistsForSingleAccount(accountID), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistGetWatchlistsForSingleAccountCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistGetWatchlistsForSingleAccountCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Watchlist(t *testing.T) {
	server := NewWatchlistService(td)

	// 建立 watchlist
	symbols := []string{"VTI", "BND"}
	watchlistItems := make([]*WatchlistItem, len(symbols))
	for i, symbol := range symbols {
		watchlistItems[i] = &WatchlistItem{
			Instrument: &Instrument{
				Symbol:    symbol,
				AssetType: "EQUITY",
			},
		}
	}

	wb := &Watchlist{
		Name:           "Test",
		WatchlistItems: watchlistItems,
	}

	CreateWatchlistCall := server.CreateWatchlist(accountID, wb)
	_, err := CreateWatchlistCall.Do()
	if err != nil {
		panic(err)
	}

	// 獲得建立的 watchlist by single
	GetWatchlistsForSingleAccountCall := server.GetWatchlistsForSingleAccount(accountID)
	ws, err := GetWatchlistsForSingleAccountCall.Do()
	if err != nil {
		panic(err)
	}
	var w *Watchlist
	for _, v := range ws.Watchlists {
		if v.Name == "Test" {
			w = v
			break
		}
	}
	content, _ := json.MarshalIndent(w, "", "    ")
	t.Log(string(content))

	// 取代建立的 watchlist
	wb.WatchlistItems = wb.WatchlistItems[1:]
	wb.Name = "TestReplace"
	ReplaceWatchlistCall := server.ReplaceWatchlist(accountID, w.WatchlistID, wb)
	_, err = ReplaceWatchlistCall.Do()
	if err != nil {
		panic(err)
	}

	// 獲得取代的 watchlist by multiple
	GetWatchlistsforMultipleAccountsCall := server.GetWatchlistsForMultipleAccounts()
	ws, err = GetWatchlistsforMultipleAccountsCall.Do()
	if err != nil {
		panic(err)
	}

	for _, v := range ws.Watchlists {
		if v.Name == "TestReplace" {
			w = v
			break
		}
	}
	content, _ = json.MarshalIndent(w, "", "    ")
	t.Log(string(content))

	// 更正 watchlist
	wb.Name = "TestUpdate"
	UpdateWatchlistCall := server.UpdateWatchlist(accountID, w.WatchlistID, wb)
	_, err = UpdateWatchlistCall.Do()
	if err != nil {
		panic(err)
	}

	// 獲得更正的 watchlist
	GetWatchlistCall := server.GetWatchlist(accountID, w.WatchlistID)
	w, err = GetWatchlistCall.Do()
	if err != nil {
		panic(err)
	}
	content, _ = json.MarshalIndent(w, "", "    ")
	t.Log(string(content))

	// 刪除取代的 watchlist
	DeleteWatchlistCall := server.DeleteWatchlist(accountID, w.WatchlistID)
	_, err = DeleteWatchlistCall.Do()
	if err != nil {
		panic(err)
	}

	// 獲得刪除的 watchlist，應為空
	GetWatchlistCall = server.GetWatchlist(accountID, w.WatchlistID)
	w, err = GetWatchlistCall.Do()
	if err == nil {
		panic(err)
	}
}
