package singleton

import (
	"reflect"
	"testing"
)

func TestNewSingelton(t *testing.T) {
	tests := []struct {
		name string
		want *Singleton
	}{
		{"case1", NewSingelton()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingelton(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingelton() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSingelton1(t *testing.T) {
	tests := []struct {
		name string
		want *Singleton
	}{
		{"case1", NewSingelton1()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingelton1(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingelton1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSingletonLazy(t *testing.T) {
	tests := []struct {
		name string
		want *Singleton
	}{
		{"case1",NewSingletonLazy()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingletonLazy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingletonLazy() = %v, want %v", got, tt.want)
			}
		})
	}
}
