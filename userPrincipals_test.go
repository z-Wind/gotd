package gotd

import (
	"reflect"
	"testing"
)

func Test_UserPrincipalGetPreferencesCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *UserPrincipalGetPreferencesCall
		want    *Preferences
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalService(td).GetPreferences(accountID), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalGetPreferencesCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalGetPreferencesCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UserPrincipalGetStreamerSubscriptionKeysCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *UserPrincipalGetStreamerSubscriptionKeysCall
		want    *SubscriptionKey
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalService(td).GetStreamerSubscriptionKeys(accountID), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalGetStreamerSubscriptionKeysCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPrincipalGetStreamerSubscriptionKeysCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UserPrincipalGetUserPrincipalsCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *UserPrincipalGetUserPrincipalsCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewUserPrincipalService(td).GetUserPrincipals(), accountID, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPrincipalGetUserPrincipalsCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.PrimaryAccountID != tt.want {
				t.Errorf("UserPrincipalGetUserPrincipalsCall.Do() = %+v, want %v", got, tt.want)
			}
		})
	}
}
