package stack

import (
	rte "petlangvm/petlang_vm/stack/petlangRTE"
)

type StackElement struct {
	Valtype rte.ValueType
	Value   rte.VmRTE
}

type Stack struct {
	Sp    int
	Bp    int
	stack [128]StackElement
}

//STACK OPERATING

func (s Stack) Push(element StackElement) {
	s.Sp--
	s.stack[s.Sp] = element
}
