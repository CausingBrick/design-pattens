package singleton

import "sync"

type Singleton struct{}

var hunger *Singleton

// init 饿汉式 init函数加载
func init() {
	hunger = &Singleton{}
}

func NewSingelton() *Singleton {
	return hunger
}

// 饿汉式 初始化加载
var hunger1 = &Singleton{}

func NewSingelton1() *Singleton {
	return hunger1
}

// 懒汉式 运行时加载
var (
	lazy     *Singleton
	lazyOnce = &sync.Once{}
)

func NewSingletonLazy() *Singleton {
	lazyOnce.Do(func() {
		lazy = &Singleton{}
	})
	return lazy
}
