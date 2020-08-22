package main

import (
	"reflect"
	"testing"
)

func Test_isOdd(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
