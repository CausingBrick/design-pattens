// 模拟电脑配件生产厂
package abstractfactory

// Accessories 配件产线抽象
type Accessories interface {
	Produce(num int)
}

// CPU 产线抽象
type CPU struct {
}

// Produce CPU生产
func (c *CPU) Produce(num int) {

}

// NetCard 网卡产线抽象
type NetCard struct {
}

// Produce 网卡生产
func (n *NetCard) Produce(num int) {

}

//AccessoriesFactory 电脑配件工厂抽象
type AccessoriesFactory interface {
	NewCPUFactory() *CPU
	NewNetCardFactory() *NetCard
}

// FactoryInShanghai 位于上海的工厂，满足成为生产配件的工厂的条件
type FactoryInShanghai struct {
}

func (f *FactoryInShanghai) NewCPUFactory() *CPU {
	return &CPU{}
}

func (f *FactoryInShanghai) NewNetCardFactory() *NetCard {
	return &NetCard{}
}
