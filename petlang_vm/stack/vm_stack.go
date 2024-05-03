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

// func (s *Stack) Sum() {

// 	var se StackElement
// 	lsv := s.Stack[s.Sp+1]
// 	rsv := s.Stack[s.Sp]

// 	//Cast by left operand(for safety)
// 	switch lsv.Valtype {
// 	case rte.Byte:
// 		se = StackElement{Valtype: rte.Byte}
// 	case rte.Integer:
// 		se = StackElement{Valtype: rte.Integer}
// 	case rte.Float:
// 		se = StackElement{Valtype: rte.Float}
// 	case rte.Reference:
// 		se = StackElement{Valtype: rte.Reference}
// 	}

// 	if(rsv.Valtype != se.Valtype){

// 	}

// }
