// 模拟电脑配件工厂
package simplefactory

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

// NewFactory 通过参数name创建工厂
func NewFactory(name string) CenterFactory {
	switch name {
	case "CPU":
		return &CPUFactory{}
	case "NetCard":
		return &NetCardFactory{}
	}
	return nil
}
