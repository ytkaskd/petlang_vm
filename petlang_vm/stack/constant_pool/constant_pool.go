package constantpool

type ConstantPool struct {
	pool map[string]int
}

func (cp *ConstantPool) Add(name string, sp int) {
	cp.pool[name] = sp
}

func (cp *ConstantPool) Get(name string) int {
	return cp.pool[name]
}
