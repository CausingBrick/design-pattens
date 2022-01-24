package simplefactory

import (
	"reflect"
	"testing"
)

func TestNewFactory(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want CenterFactory
	}{
		{"CPU", args{"CPU"}, &CPUFactory{}},
		{"NetCard", args{"NetCard"}, &NetCardFactory{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFactory(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
