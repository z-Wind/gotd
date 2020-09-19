package gotd

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func Test_Watchlist(t *testing.T) {
	if !onlineTest {
		t.Skipf("online Test setting is %v", onlineTest)
	}

	server := NewWatchlistService(tdReal)

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
func TestWatchlistGetWatchlistsForMultipleAccountsCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistsForMultipleAccountsCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).GetWatchlistsForMultipleAccounts(), "https://api.tdameritrade.com/v1/accounts/watchlists?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistGetWatchlistsForMultipleAccountsCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistGetWatchlistsForMultipleAccountsCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestWatchlistGetWatchlistsForMultipleAccountsCall_Do(t *testing.T) {
	client := clientTest(`[{"name":"default","watchlistId":"1234567893","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"SPY","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"AAPL","assetType":"EQUITY"}}]},{"name":"Quotes","watchlistId":"1234567898","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"IBM","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"MSFT","assetType":"EQUITY"}},{"sequenceId":3,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"EBAY","assetType":"EQUITY"}}]},{"name":"Indexes","watchlistId":"1234567896","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"$DJI","assetType":"INDEX"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"SPX","assetType":"EQUITY"}},{"sequenceId":3,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"COMP","assetType":"EQUITY"}},{"sequenceId":4,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VIX","assetType":"EQUITY"}}]},{"name":"test","watchlistId":"1234567890","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VTI","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VIX","assetType":"EQUITY"}}]},{"name":"ETF","watchlistId":"12345678","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VTI","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VBR","assetType":"EQUITY"}}]}]`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistsForMultipleAccountsCall
		want    *WatchlistList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewWatchlistService(tdTest).GetWatchlistsForMultipleAccounts(), &WatchlistList{
			ServerResponse: ServerResponse{200, http.Header{}},
			Watchlists: []*Watchlist{
				{
					Name:        "default",
					WatchlistID: "1234567893",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "SPY",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "AAPL",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "Quotes",
					WatchlistID: "1234567898",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "IBM",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "MSFT",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   3,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "EBAY",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "Indexes",
					WatchlistID: "1234567896",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "$DJI",
								AssetType: "INDEX",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "SPX",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   3,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "COMP",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   4,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VIX",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "test",
					WatchlistID: "1234567890",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VTI",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VIX",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "ETF",
					WatchlistID: "12345678",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VTI",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VBR",
								AssetType: "EQUITY",
							},
						},
					},
				},
			},
		}, false},
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

func TestWatchlistGetWatchlistsForSingleAccountCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistsForSingleAccountCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).GetWatchlistsForSingleAccount("accountID"), "https://api.tdameritrade.com/v1/accounts/accountID/watchlists?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistGetWatchlistsForSingleAccountCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistGetWatchlistsForSingleAccountCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestWatchlistGetWatchlistsForSingleAccountCall_Do(t *testing.T) {
	client := clientTest(`[{"name":"default","watchlistId":"1234567893","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"SPY","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"AAPL","assetType":"EQUITY"}}]},{"name":"Quotes","watchlistId":"1234567898","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"IBM","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"MSFT","assetType":"EQUITY"}},{"sequenceId":3,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"EBAY","assetType":"EQUITY"}}]},{"name":"Indexes","watchlistId":"1234567896","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"$DJI","assetType":"INDEX"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"SPX","assetType":"EQUITY"}},{"sequenceId":3,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"COMP","assetType":"EQUITY"}},{"sequenceId":4,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VIX","assetType":"EQUITY"}}]},{"name":"test","watchlistId":"1234567890","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VTI","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VIX","assetType":"EQUITY"}}]},{"name":"ETF","watchlistId":"12345678","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VTI","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"VBR","assetType":"EQUITY"}}]}]`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistsForSingleAccountCall
		want    *WatchlistList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewWatchlistService(tdTest).GetWatchlistsForSingleAccount(accountID), &WatchlistList{
			ServerResponse: ServerResponse{200, http.Header{}},
			Watchlists: []*Watchlist{
				{
					Name:        "default",
					WatchlistID: "1234567893",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "SPY",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "AAPL",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "Quotes",
					WatchlistID: "1234567898",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "IBM",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "MSFT",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   3,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "EBAY",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "Indexes",
					WatchlistID: "1234567896",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "$DJI",
								AssetType: "INDEX",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "SPX",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   3,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "COMP",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   4,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VIX",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "test",
					WatchlistID: "1234567890",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VTI",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VIX",
								AssetType: "EQUITY",
							},
						},
					},
				},
				{
					Name:        "ETF",
					WatchlistID: "12345678",
					AccountID:   "123456789",
					WatchlistItems: []*WatchlistItem{
						{
							SequenceID:   1,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VTI",
								AssetType: "EQUITY",
							},
						},
						{
							SequenceID:   2,
							Quantity:     0,
							AveragePrice: 0,
							Commission:   0,
							Instrument: &Instrument{
								Symbol:    "VBR",
								AssetType: "EQUITY",
							},
						},
					},
				},
			},
		}, false},
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

func TestWatchlistCreateWatchlistCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistCreateWatchlistCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).CreateWatchlist("accountID", &Watchlist{}), "https://api.tdameritrade.com/v1/accounts/accountID/watchlists?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistCreateWatchlistCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistCreateWatchlistCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistCreateWatchlistCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistCreateWatchlistCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).CreateWatchlist("accountID", &Watchlist{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistCreateWatchlistCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistCreateWatchlistCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistDeleteWatchlistCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistDeleteWatchlistCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).DeleteWatchlist("accountID", "watchlistID"), "https://api.tdameritrade.com/v1/accounts/accountID/watchlists/watchlistID?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistDeleteWatchlistCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistDeleteWatchlistCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistDeleteWatchlistCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistDeleteWatchlistCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).DeleteWatchlist("accountID", "watchlistID"), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistDeleteWatchlistCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistDeleteWatchlistCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistGetWatchlistCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).GetWatchlist("accountID", "watchlistID"), "https://api.tdameritrade.com/v1/accounts/accountID/watchlists/watchlistID?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistGetWatchlistCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistGetWatchlistCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistGetWatchlistCall_Do(t *testing.T) {
	client := clientTest(`{"name":"default","watchlistId":"1234567893","accountId":"123456789","watchlistItems":[{"sequenceId":1,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"SPY","assetType":"EQUITY"}},{"sequenceId":2,"quantity":0,"averagePrice":0,"commission":0,"instrument":{"symbol":"AAPL","assetType":"EQUITY"}}]}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistGetWatchlistCall
		want    *Watchlist
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).GetWatchlist("accountID", "watchlistID"), &Watchlist{
			ServerResponse: ServerResponse{200, http.Header{}},
			Name:           "default",
			WatchlistID:    "1234567893",
			AccountID:      "123456789",
			WatchlistItems: []*WatchlistItem{
				{
					SequenceID:   1,
					Quantity:     0,
					AveragePrice: 0,
					Commission:   0,
					Instrument: &Instrument{
						Symbol:    "SPY",
						AssetType: "EQUITY",
					},
				},
				{
					SequenceID:   2,
					Quantity:     0,
					AveragePrice: 0,
					Commission:   0,
					Instrument: &Instrument{
						Symbol:    "AAPL",
						AssetType: "EQUITY",
					},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistGetWatchlistCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistGetWatchlistCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistReplaceWatchlistCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistReplaceWatchlistCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).ReplaceWatchlist("accountID", "watchlistID", &Watchlist{}), "https://api.tdameritrade.com/v1/accounts/accountID/watchlists/watchlistID?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistReplaceWatchlistCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistReplaceWatchlistCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistReplaceWatchlistCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistReplaceWatchlistCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).ReplaceWatchlist("accountID", "watchlistID", &Watchlist{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistReplaceWatchlistCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistReplaceWatchlistCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistUpdateWatchlistCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistUpdateWatchlistCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).UpdateWatchlist("accountID", "watchlistID", &Watchlist{}), "https://api.tdameritrade.com/v1/accounts/accountID/watchlists/watchlistID?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistUpdateWatchlistCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistUpdateWatchlistCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchlistUpdateWatchlistCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *WatchlistUpdateWatchlistCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", NewWatchlistService(tdTest).UpdateWatchlist("accountID", "watchlistID", &Watchlist{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchlistUpdateWatchlistCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchlistUpdateWatchlistCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
