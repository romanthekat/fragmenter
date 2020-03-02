package main

import "testing"

func Test_getOffsetAndFlag(t *testing.T) {
	type args struct {
		offset     uint16
		moreFollow bool
	}

	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"full zeroes", args{offset: 0, moreFollow: false}, 0},
		{"zeroes(last)", args{offset: 0, moreFollow: true}, 1},
		{"some value(last)", args{offset: 1, moreFollow: true}, 0b0000_0000_0000_1001},
		{"big value", args{offset: 0b0000_1111_0100_1, moreFollow: false}, 0b0000_1111_0100_1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOffsetAndFlag(tt.args.offset, tt.args.moreFollow); got != tt.want {
				t.Errorf("getOffsetAndFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
