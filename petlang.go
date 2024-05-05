package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	petlangvm "petlangvm/petlang_vm"
)

func main() {
	fmt.Println("Petlang v0.1")

	stackSize := flag.Int("ss", 256, "\033[94mPetlangVM stack size\033[0m")

	flag.Parse()
	fmt.Printf("Stack size -ss: %d\n", *stackSize)
	petlangvm.Preload(loadByteCode(os.Args[1]), *stackSize)
	// petlangvm.EvalByteCode()
	// petlangvm.PrintStack()
	fmt.Println("Petlang exit")
}

func loadByteCode(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open file %s, \nerror:\n%v", filename, err)
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
	fmt.Println("Openned")
	return bc
}
