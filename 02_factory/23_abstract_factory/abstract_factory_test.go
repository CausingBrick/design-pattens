// 模拟电脑配件生产厂
package abstractfactory

import (
	"reflect"
	"testing"
)

func TestFactoryInShanghai_NewCPUFactory(t *testing.T) {
	tests := []struct {
		name string
		want *CPU
	}{
		{
			"Produce CPU",
			&CPU{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FactoryInShanghai{}
			if got := f.NewCPUFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FactoryInShanghai.NewCPUFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactoryInShanghai_NewNetCardFactory(t *testing.T) {
	tests := []struct {
		name string
		want *NetCard
	}{
		{
			"Produce netcard",
			&NetCard{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FactoryInShanghai{}
			if got := f.NewNetCardFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FactoryInShanghai.NewNetCardFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
