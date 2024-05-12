package petlangvm

import (
	"fmt"
	"petlangvm/petlang_vm/stack"
	constantpool "petlangvm/petlang_vm/stack/constant_pool"
	rte "petlangvm/petlang_vm/stack/petlangRTE"
	vm_errors "petlangvm/petlang_vm/vm_error_service"
	opcode "petlangvm/petlang_vm/vm_opcodes"
)

type VM struct {
	ip       int
	bytecode []byte
	stack    *stack.Stack
	Pool     *constantpool.ConstantPool
}

func (vm *VM) Preload(bc []byte, stackSize int) {
	vm.bytecode = bc
	vm.stack = new(stack.Stack)
	vm.stack.StackSize = stackSize
	vm.stack.Bp = stackSize - 1
	vm.stack.Sp = stackSize - 1
	vm.stack.Stack = make([]stack.StackElement, stackSize)
	vm.Pool = new(constantpool.ConstantPool)
}

func (vm *VM) EvalByteCode() {

	//check bytecode valid
	if vm.readWordX32(0) != 0xFEE1DEAD {
		vm_errors.ThrowError(vm_errors.NOTBYTECODE, vm.ip)
	}
	//find where is import section
	if ip := vm.findImportSection(); ip == 0 {
		vm_errors.ThrowError(vm_errors.MODSECFINDFAIL, ip)
	}
	for ; vm.ip != len(vm.bytecode); vm.ip++ {
		fmt.Printf("\n\n INSTRUCTION: 0x%02x\n\n", vm.bytecode[vm.ip])
		switch vm.bytecode[vm.ip] {
		//TODO: Implement other opcodes from opcodes.go
		case opcode.PUSHBYTE:
			fmt.Println("\nPUSH BYTE command")
			vm.ip++
			value := vm.bytecode[vm.ip]
			se := stack.StackElement{Valtype: rte.Byte, Value: value}
			vm.ip++
			vm.stack.Push(se)

		case opcode.PUSHINT:
			fmt.Println("\nPUSH INT command")
			vm.ip++
			se := stack.StackElement{Valtype: rte.Integer, Value: vm.readWordX32(vm.ip)}
			vm.ip++
			vm.stack.Push(se)

		case opcode.PUSHFLOAT:
			fmt.Println("\nPUSH FLOAT command")
			vm.ip++
			se := stack.StackElement{Valtype: rte.Float, Value: vm.readWordX32(vm.ip)}
			vm.ip++
			vm.stack.Push(se)

		case opcode.PUSHREF:
			fmt.Println("\nPUSH REF command")
			vm.ip++
			se := stack.StackElement{Valtype: rte.Reference, Value: vm.readWordX32(vm.ip)}
			vm.ip++
			vm.stack.Push(se)

		case opcode.POP:
			fmt.Println("\nPOP command")
			vm.ip++

		case opcode.ALIGNVAR:
			fmt.Println("Allign")
			vm.ip++
			name := vm.readName(vm.ip)
			addr := vm.stack.Sp
			vm.Pool.Add(name, addr)

		default:
			vm_errors.ThrowError(vm_errors.UNKNOWNOPCODE, vm.ip)
		}

	}

}

func (vm *VM) findImportSection() int {
	for ind, value := range vm.bytecode {
		if value == opcode.MODSEC {
			return ind
		}
	}
	return 0
}

func (vm *VM) readWordX32(start int) int {
	var word int
	var offset int = 24
	for i := 0; i != 4; i++ {
		vm.ip++
		word += int(vm.bytecode[start+i]) << offset
		offset -= 8
	}
	fmt.Printf("word: 0x%02x\n", word)
	return word
}

func (vm *VM) readName(start int) string {
	name := ""
	for i := start; vm.bytecode[i] != '$'; i++ {
		name += string(vm.bytecode[i])
		vm.ip++
		fmt.Printf("String: %s\n", name)
	}
	if name != "" {
		return name
	}
	vm_errors.ThrowError(vm_errors.NOTBYTECODE, vm.ip)
	return name
}

func (vm *VM) PrintStack() {
	for index, value := range vm.stack.Stack {
		fmt.Printf(": %d :||: %02x :", index, value.Value)
	}
}
