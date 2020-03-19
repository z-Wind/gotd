package gotd

import (
	"reflect"
	"testing"
)

func Test_UserPrincipalsGetPreferencesCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *UserPrincipalsGetPreferencesCall
		want    *Preferences
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(td).GetPreferences(accountID), nil, false},
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

func Test_UserPrincipalsGetStreamerSubscriptionKeysCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *UserPrincipalsGetStreamerSubscriptionKeysCall
		want    *SubscriptionKey
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(td).GetStreamerSubscriptionKeys(accountID), nil, false},
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

func Test_UserPrincipalsGetUserPrincipalsCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *UserPrincipalsGetUserPrincipalsCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalsService(td).GetUserPrincipals(), accountID, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalsGetUserPrincipalsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.PrimaryAccountID != tt.want {
				t.Errorf("UserPrincipalsGetUserPrincipalsCall.Do() = %+v, want %v", got, tt.want)
			}
		})
	}
}
