package keycloakmiddleware

import (
	"testing"
)

func Test_isScopesValid(t *testing.T) {
	type args struct {
		claims claims
		scopes []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "scope valid",
			args: args{
				claims: claims{
					Authorization: authorization{
						Permissions: []permission{
							{Scopes: []string{"foo", "bar", "baz"}},
							{Scopes: []string{"qux", "fred"}},
						},
					},
				},
				scopes: []string{"bar", "fred"},
			},
			want: true,
		},
		{
			name: "scope valid 2",
			args: args{
				claims: claims{
					Authorization: authorization{
						Permissions: []permission{
							{Scopes: []string{"foo", "bar", "baz"}},
							{Scopes: []string{"qux", "fred"}},
						},
					},
				},
				scopes: []string{"qux", "thud"},
			},
			want: true,
		},
		{
			name: "scope valid 2",
			args: args{
				claims: claims{
					Authorization: authorization{
						Permissions: []permission{
							{Scopes: []string{"foo", "bar", "baz"}},
							{Scopes: []string{"qux", "fred"}},
						},
					},
				},
				scopes: []string{"thud", "chips"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScopesValid(tt.args.claims, tt.args.scopes); got != tt.want {
				t.Errorf("isScopesValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
