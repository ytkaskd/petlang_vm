package petlangvm

import (
	"fmt"
	"petlangvm/petlang_vm/stack"
	rte "petlangvm/petlang_vm/stack/petlangRTE"
	vm_errors "petlangvm/petlang_vm/vm_error_service"
	opcode "petlangvm/petlang_vm/vm_opcodes"
)

var ip int = 0
var bytecode []byte
var vmstack stack.Stack

func Preload(bc []byte) {
	bytecode = bc
	vmstack.Sp = 127
	vmstack.Bp = 127
}

func EvalByteCode(bc []byte) {

	//check bytecode valid
	if readWordX32(0) != 0xFEE1DEAD {
		vm_errors.ThrowError(vm_errors.NOTBYTECODE)
	}
	//find where is import section
	// if ip := findImportSection(); ip == 0 {
	// 	return MODSECFINDFAIL
	// }
	for ; ip != len(bytecode); ip++ {
		fmt.Printf("\n\n INSTRUCTION: 0x%02x\n\n", bytecode[ip])
		switch bytecode[ip] {
		//TODO: Implement PUSHINT, PUSHFLOAT, PUSHREF and other opcodes from opcodes.go
		case opcode.PUSHBYTE:
			fmt.Println("PUSH BYTE command")
			ip++
			value := rte.PetlangByte{Value: bytecode[ip]}
			se := stack.StackElement{Valtype: rte.Byte, Value: value}
			vmstack.Push(se)
		}
	}

}

func findImportSection() int {
	for ind, value := range bytecode {
		if value == 0xAD {
			return ind
		}
	}
	return 0
}

func readWordX32(start int) int {
	var word int
	var offset int = 24
	for i := 0; i != 4; i++ {
		ip++
		word += int(bytecode[start+i]) << offset
		offset -= 8
	}
	fmt.Printf("word: 0x%02x\n", word)
	return word
}
