package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	petlangvm "petlangvm/petlang_vm"
	vm_errors "petlangvm/petlang_vm/vm_error_service"
)

func main() {
	fmt.Println("Petlang v0.1")

	pvm := new(petlangvm.VM)

	stackSize := flag.Int("ss", 256, "\033[94mPetlangVM stack size\033[0m")
	dbgMode := flag.Bool("dbg", false, "\033[94mEnable debug mode\033[0m")
	flag.Parse()

	if bc := loadByteCode(os.Args[len(os.Args)-1]); bc != nil {
		pvm.Preload(bc, *stackSize)
		pvm.EvalByteCode()
		if *dbgMode {
			pvm.PrintStack()
		}
		fmt.Println("Petlang exit")
	} else {
		vm_errors.ThrowError(0x00, 0x00)
	}

}

func loadByteCode(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open file %s", filename)
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	bc := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bc)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return nil
	}
	//fmt.Println("Openned")
	return bc
}
