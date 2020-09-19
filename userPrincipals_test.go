package gotd

import (
	"net/http"
	"reflect"
	"testing"
)

func TestUserPrincipalsGetPreferencesCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsGetPreferencesCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).GetPreferences("accountID"), "https://api.tdameritrade.com/v1/accounts/accountID/preferences?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsGetPreferencesCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsGetPreferencesCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrincipalsGetPreferencesCall_Do(t *testing.T) {
	client := clientTest(`{"expressTrading":false,"directOptionsRouting":false,"directEquityRouting":false,"defaultEquityOrderLegInstruction":"NONE","defaultEquityOrderType":"LIMIT","defaultEquityOrderPriceLinkType":"NONE","defaultEquityOrderDuration":"GOOD_TILL_CANCEL","defaultEquityOrderMarketSession":"NORMAL","defaultEquityQuantity":0,"mutualFundTaxLotMethod":"FIFO","optionTaxLotMethod":"FIFO","equityTaxLotMethod":"FIFO","defaultAdvancedToolLaunch":"NONE","authTokenTimeout":"FIFTY_FIVE_MINUTES"}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsGetPreferencesCall
		want    *Preferences
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).GetPreferences("accountID"), &Preferences{
			ServerResponse:                   ServerResponse{200, http.Header{}},
			ExpressTrading:                   false,
			DirectOptionsRouting:             false,
			DirectEquityRouting:              false,
			DefaultEquityOrderLegInstruction: "NONE",
			DefaultEquityOrderType:           "LIMIT",
			DefaultEquityOrderPriceLinkType:  "NONE",
			DefaultEquityOrderDuration:       "GOOD_TILL_CANCEL",
			DefaultEquityOrderMarketSession:  "NORMAL",
			DefaultEquityQuantity:            0,
			MutualFundTaxLotMethod:           "FIFO",
			OptionTaxLotMethod:               "FIFO",
			EquityTaxLotMethod:               "FIFO",
			DefaultAdvancedToolLaunch:        "NONE",
			AuthTokenTimeout:                 "FIFTY_FIVE_MINUTES",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsGetPreferencesCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsGetPreferencesCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrincipalsGetStreamerSubscriptionKeysCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsGetStreamerSubscriptionKeysCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).GetStreamerSubscriptionKeys("accountID1", "accountID2"), "https://api.tdameritrade.com/v1/userprincipals/streamersubscriptionkeys?accountIds=accountID1%2CaccountID2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsGetStreamerSubscriptionKeysCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsGetStreamerSubscriptionKeysCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrincipalsGetStreamerSubscriptionKeysCall_Do(t *testing.T) {
	client := clientTest(`{"keys":[{"key":"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567891011"}]}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsGetStreamerSubscriptionKeysCall
		want    *SubscriptionKey
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).GetStreamerSubscriptionKeys("accountID1", "accountID2"), &SubscriptionKey{
			ServerResponse: ServerResponse{200, http.Header{}},
			Keys: []*Key{
				{
					"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567891011",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsGetStreamerSubscriptionKeysCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsGetStreamerSubscriptionKeysCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrincipalsGetUserPrincipalsCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsGetUserPrincipalsCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).GetUserPrincipals(), "https://api.tdameritrade.com/v1/userprincipals?", false},
		{"All", NewUserPrincipalsService(tdTest).GetUserPrincipals().
			Fields(
				UserPrincipalsFieldsStreamerSubscriptionKeys,
				UserPrincipalsFieldsStreamerConnectionInfo,
				UserPrincipalsFieldsPreferences,
				UserPrincipalsFieldsSurrogateIds,
			), "https://api.tdameritrade.com/v1/userprincipals?fields=streamerSubscriptionKeys%2CstreamerConnectionInfo%2Cpreferences%2CsurrogateIds", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsGetUserPrincipalsCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsGetUserPrincipalsCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrincipalsGetUserPrincipalsCall_Do(t *testing.T) {
	client := clientTest(`{"userId":"abcdefghijkl","userCdDomainId":"A123456789703389","primaryAccountId":"123456789","lastLoginTime":"2018-11-03T08:52:14+0000","tokenExpirationTime":"2018-11-03T09:51:28+0000","loginTime":"2018-11-03T09:21:28+0000","accessLevel":"CUS","stalePassword":false,"streamerInfo":{"streamerBinaryUrl":"streamer-bin.tdameritrade.com","streamerSocketUrl":"streamer-ws.tdameritrade.com","token":"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn","tokenTimestamp":"2018-11-03T09:35:37+0000","userGroup":"ACCT","accessLevel":"ACCT","acl":"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz12345678910111213141516171819202","appId":"ABCDEFGHIJK"},"professionalStatus":"NON_PROFESSIONAL","quotes":{"isNyseDelayed":false,"isNasdaqDelayed":false,"isOpraDelayed":false,"isAmexDelayed":false,"isCmeDelayed":true,"isIceDelayed":true,"isForexDelayed":true},"streamerSubscriptionKeys":{"keys":[{"key":"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567891011"}]},"accounts":[{"accountId":"123456789","displayName":"ABCDEFGHIJ","accountCdDomainId":"A123456789703390","company":"AMER","segment":"AMER","surrogateIds":{"SCARR":"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn","DART":"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn"},"preferences":{"expressTrading":false,"directOptionsRouting":false,"directEquityRouting":false,"defaultEquityOrderLegInstruction":"NONE","defaultEquityOrderType":"LIMIT","defaultEquityOrderPriceLinkType":"NONE","defaultEquityOrderDuration":"GOOD_TILL_CANCEL","defaultEquityOrderMarketSession":"NORMAL","defaultEquityQuantity":0,"mutualFundTaxLotMethod":"FIFO","optionTaxLotMethod":"FIFO","equityTaxLotMethod":"FIFO","defaultAdvancedToolLaunch":"NONE","authTokenTimeout":"FIFTY_FIVE_MINUTES"},"acl":"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456","authorizations":{"apex":false,"levelTwoQuotes":false,"stockTrading":true,"marginTrading":false,"streamingNews":false,"optionTradingLevel":"NONE","streamerAccess":true,"advancedMargin":false,"scottradeAccount":false}}]}`, http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsGetUserPrincipalsCall
		want    *UserPrincipal
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).GetUserPrincipals(), &UserPrincipal{
			ServerResponse:      ServerResponse{200, http.Header{}},
			UserID:              "abcdefghijkl",
			UserCdDomainID:      "A123456789703389",
			PrimaryAccountID:    "123456789",
			LastLoginTime:       "2018-11-03T08:52:14+0000",
			TokenExpirationTime: "2018-11-03T09:51:28+0000",
			LoginTime:           "2018-11-03T09:21:28+0000",
			AccessLevel:         "CUS",
			StalePassword:       false,
			StreamerInfo: &StreamerInfo{
				StreamerBinaryURL: "streamer-bin.tdameritrade.com",
				StreamerSocketURL: "streamer-ws.tdameritrade.com",
				Token:             "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn",
				TokenTimestamp:    "2018-11-03T09:35:37+0000",
				UserGroup:         "ACCT",
				AccessLevel:       "ACCT",
				ACL:               "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz12345678910111213141516171819202",
				AppID:             "ABCDEFGHIJK",
			},
			ProfessionalStatus: "NON_PROFESSIONAL",
			Quotes: &Quotes{
				IsNyseDelayed:   false,
				IsNasdaqDelayed: false,
				IsOpraDelayed:   false,
				IsAmexDelayed:   false,
				IsCmeDelayed:    true,
				IsIceDelayed:    true,
				IsForexDelayed:  true,
			},
			StreamerSubscriptionKeys: &SubscriptionKey{
				Keys: []*Key{
					{
						"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567891011",
					},
				},
			},
			Accounts: []*AccountU{
				{
					AccountID:         "123456789",
					DisplayName:       "ABCDEFGHIJ",
					AccountCdDomainID: "A123456789703390",
					Company:           "AMER",
					Segment:           "AMER",
					SurrogateIds: map[string]string{
						"SCARR": "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn",
						"DART":  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn",
					},
					Preferences: Preferences{
						ExpressTrading:                   false,
						DirectOptionsRouting:             false,
						DirectEquityRouting:              false,
						DefaultEquityOrderLegInstruction: "NONE",
						DefaultEquityOrderType:           "LIMIT",
						DefaultEquityOrderPriceLinkType:  "NONE",
						DefaultEquityOrderDuration:       "GOOD_TILL_CANCEL",
						DefaultEquityOrderMarketSession:  "NORMAL",
						DefaultEquityQuantity:            0,
						MutualFundTaxLotMethod:           "FIFO",
						OptionTaxLotMethod:               "FIFO",
						EquityTaxLotMethod:               "FIFO",
						DefaultAdvancedToolLaunch:        "NONE",
						AuthTokenTimeout:                 "FIFTY_FIVE_MINUTES",
					},
					ACL: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456",
					Authorizations: Authorizations{
						Apex:               false,
						LevelTwoQuotes:     false,
						StockTrading:       true,
						MarginTrading:      false,
						StreamingNews:      false,
						OptionTradingLevel: "NONE",
						StreamerAccess:     true,
						AdvancedMargin:     false,
						ScottradeAccount:   false,
					},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsGetUserPrincipalsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsGetUserPrincipalsCall.Do() = %+v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrincipalsUpdatePreferencesCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsUpdatePreferencesCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).UpdatePreferences("accountID", &Preferences{}), "https://api.tdameritrade.com/v1/accounts/accountID/preferences?", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsUpdatePreferencesCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsUpdatePreferencesCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrincipalsUpdatePreferencesCall_Do(t *testing.T) {
	client := clientTest("", http.StatusOK)
	tdTest, _ := New(client)

	tests := []struct {
		name    string
		c       *UserPrincipalsUpdatePreferencesCall
		want    *ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(tdTest).UpdatePreferences("accountID", &Preferences{}), &ServerResponse{200, http.Header{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsUpdatePreferencesCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalsUpdatePreferencesCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
