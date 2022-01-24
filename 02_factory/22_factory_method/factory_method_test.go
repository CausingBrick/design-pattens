package factorymethod

import (
	"reflect"
	"testing"
)

func TestNewFactoryBuilder(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want CenterFactoryBuilder
	}{
		{
			name:"CPU",
			args: args{
				"CPU",
			},
			want: &CPUFactoryBuilder{},
		},
		{
			name:"NetCard",
			args: args{
				"NetCard",
			},
			want: &NetCardFactoryBuilder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFactoryBuilder(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFactoryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
