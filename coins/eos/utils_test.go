package eos

import (
	"reflect"
	"testing"

	"github.com/xiaohuasheng0x1/blockchains/coins/eos/types"
)

func TestHexToChecksum256(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want types.Checksum256
	}{
		{name: "test1", args: args{data: "a"}, want: []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToChecksum256(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexToChecksum256() = %v, want %v", got, tt.want)
			}
		})
	}
}
