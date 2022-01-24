package factorymethod

// CenterFactory 工厂接口
type CenterFactory interface {
	Produce(num int)
}

// NetCardFactory 网卡工厂
type NetCardFactory struct {
}

func (n *NetCardFactory) Produce(num int) {

}

// CPUFactory CPU工厂
type CPUFactory struct {
}

func (c *CPUFactory) Produce(num int) {

}

// CenterFactoryBuilder 工厂方法接口
type CenterFactoryBuilder interface {
	Build() CenterFactory
}

// NetCardFactoryBuilder 网卡工厂接口
type NetCardFactoryBuilder struct {
}

func (n *NetCardFactoryBuilder) Build() CenterFactory {
	return &NetCardFactory{}
}

// CPUFactoryBuilder CPU工厂接口
type CPUFactoryBuilder struct {
}

func (c *CPUFactoryBuilder) Build() CenterFactory {
	return &CPUFactory{}
}

// NewFactoryBuilder 通过参数name创建工厂接口
func NewFactoryBuilder(name string) CenterFactoryBuilder {
	switch name {
	case "CPU":
		return &CPUFactoryBuilder{}
	case "NetCard":
		return &NetCardFactoryBuilder{}
	}
	return nil
}
