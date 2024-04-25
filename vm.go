package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	// opcodes
	PUSH   int = 0xF0
	POP    int = 0xF1
	RETURN int = 0xF2
	CALL   int = 0xF3
	SFRAME int = 0xF4 //create stack frame
	SUM    int = 0xA0
	DIV    int = 0xA1
	MUL    int = 0xA2

	MODSEC  int = 0xAD //modules section
	ADDMOD  int = 0xD0 // add mod to loader
	MODNEND int = 0xD1
	MDSEND  int = 0xDA // end module section

	//error codes

	SUCCESSEVAL    int = 0xDEADD011
	NOTBYTECODE    int = 0xBADC0DE
	MODSECFINDFAIL int = 0xFA11EDF1
)

var ip int = 4
var bytecode []byte

var stack []int
var sp int = 0
var bp int = 0

func main() {
	fmt.Println("Petlang v0.1")
	bytecode = loadByteCode(os.Args[1])
	errcode := evalByteCode()
	decodeErrors(errcode)
	fmt.Println("Petlang exit")
}

func decodeErrors(errcode int) {
	if errcode != SUCCESSEVAL {
		fmt.Printf("Petlang vm error, code: 0x%02x\n", errcode)
		switch errcode {
		case NOTBYTECODE:
			fmt.Printf("file isn't petlang bytecode\n")
		case MODSECFINDFAIL:
			fmt.Printf("Can't find modules import section")
		}
	}
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

func evalByteCode() int {
	//check bytecode valid
	if readWordX32(0) != 0xFEE1DEAD {
		return NOTBYTECODE
	}
	//find where is import section
	if ip := findImportSection(); ip == 0 {
		return MODSECFINDFAIL
	}

	return SUCCESSEVAL
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
		word += int(bytecode[start+i]) << offset
		offset -= 8
	}
	fmt.Printf("word: 0x%02x\n", word)
	return word
}

//STACK OPERATING

func push(value int) {
	stack = append(stack, value)
	sp++
}

func pop() int {
	value := stack[sp]
	stack = stack[:len(stack)-1]
	return value
}

func sum() {
	leftop := stack[sp-1]
	rightop := stack[sp]
	push(leftop + rightop)
}

func div() {
	leftop := stack[sp-1]
	rightop := stack[sp]
	push(leftop - rightop)
}

func mul() {
	leftop := stack[sp-1]
	rightop := stack[sp]
	push(leftop * rightop)
}
