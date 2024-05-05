package stack

import (
	rte "petlangvm/petlang_vm/stack/petlangRTE"
	vm_errors "petlangvm/petlang_vm/vm_error_service"
)

type StackElement struct {
	Valtype rte.ValueType
	Value   any
}

type Stack struct {
	StackSize int
	Sp        int
	Bp        int
	Stack     []StackElement
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

func (s *Stack) Sum() {
	var se StackElement
	lsv := s.Stack[s.Sp+1]
	rsv := s.Stack[s.Sp]
	s.cast(&lsv, &se)
	s.cast(&lsv, &rsv)
}

func (s *Stack) cast(castTo *StackElement, se *StackElement) {
	//Cast by left operand(for safety)
	if castTo.Valtype != rte.Reference && se.Valtype != rte.Reference {
		se.Valtype = castTo.Valtype
	} else {
		vm_errors.ThrowError(vm_errors.UNSUPADAR, s.Sp)
	}
}
