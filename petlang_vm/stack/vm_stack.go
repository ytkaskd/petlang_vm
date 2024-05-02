package stack

import (
	rte "petlangvm/petlang_vm/stack/petlangRTE"
	vm_errors "petlangvm/petlang_vm/vm_error_service"
)

type StackElement struct {
	Valtype rte.ValueType
	Value   rte.VmRTE
}

type Stack struct {
	Sp    int
	Bp    int
	Stack [256]StackElement
}

//STACK OPERATING

func (s *Stack) Push(element StackElement) {
	s.Sp--
	if s.Sp < 0 {
		vm_errors.ThrowError(vm_errors.STACKOVERFLOW, s.Sp)
	} else {
		s.Stack[s.Sp] = element
	}
}
