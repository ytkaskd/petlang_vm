package constantpool

type ConstantPool struct {
	pool map[string]int
}

func (cp *ConstantPool) Add(name string, addr int) {
	cp.pool[name] = addr
}

func (cp *ConstantPool) Get(name string) int {
	return cp.pool[name]
}
