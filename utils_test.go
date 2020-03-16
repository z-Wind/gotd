package gotd

import "testing"

func TestResolveRelative(t *testing.T) {
	type args struct {
		basePath string
		elem     []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test", args{"https://api.tdameritrade.com/v1", []string{"accounts", "a", "b"}}, "https://api.tdameritrade.com/v1/accounts/a/b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveRelative(tt.args.basePath, tt.args.elem...); got != tt.want {
				t.Errorf("ResolveRelative() = %v, want %v", got, tt.want)
			}
		})
	}
}
