package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// add test cases
		{
			name: "DB_CONNECTION",
			args: args{
				key: "DB_CONNECTION",
			},
			want: "mongodb://zeroPass:9674Ephzx&T7@127.0.0.1:27017/?retryWrites=true&w=majority",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetEnv(tt.args.key), "GetEnv(%v)", tt.args.key)
		})
	}
}
