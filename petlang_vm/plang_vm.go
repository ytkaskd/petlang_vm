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
	vmstack.Sp = 255
	vmstack.Bp = 255
}

func EvalByteCode() {

	//check bytecode valid
	if readWordX32(0) != 0xFEE1DEAD {
		vm_errors.ThrowError(vm_errors.NOTBYTECODE, ip)
	}
	//find where is import section
	if ip := findImportSection(); ip == 0 {
		vm_errors.ThrowError(vm_errors.MODSECFINDFAIL, ip)
	}
	for ; ip != len(bytecode); ip++ {
		fmt.Printf("\n\n INSTRUCTION: 0x%02x\n\n", bytecode[ip])
		switch bytecode[ip] {
		//TODO: Implement other opcodes from opcodes.go
		case opcode.PUSHBYTE:
			fmt.Println("\nPUSH BYTE command")
			ip++
			value := rte.PetlangByte{Value: bytecode[ip]}
			se := stack.StackElement{Valtype: rte.Byte, Value: value}
			vmstack.Push(se)

		case opcode.PUSHINT:
			fmt.Println("\nPUSH INT command")
			ip++
			value := rte.PetlangInt{Value: readWordX32(ip)}
			se := stack.StackElement{Valtype: rte.Integer, Value: value}
			vmstack.Push(se)

		case opcode.PUSHFLOAT:
			fmt.Println("\nPUSH FLOAT command")
			ip++
			value := rte.PetlangFLoat32{Value: float32(readWordX32(ip))}
			se := stack.StackElement{Valtype: rte.Float, Value: value}
			vmstack.Push(se)

		case opcode.PUSHREF:
			fmt.Println("\nPUSH REF command")
			ip++
			value := rte.PetlangRef{Value: readWordX32(ip)}
			se := stack.StackElement{Valtype: rte.Reference, Value: value}
			vmstack.Push(se)

		case opcode.POP:
			fmt.Println("\nPOP command")
			ip++

		default:
			vm_errors.ThrowError(vm_errors.UNKNOWNOPCODE, ip)
		}

	}

}

func findImportSection() int {
	for ind, value := range bytecode {
		if value == opcode.MODSEC {
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

func PrintStack() {
	for index, value := range vmstack.Stack {
		fmt.Printf(": %d :||: %02x :", index, value.Value)
	}
}
