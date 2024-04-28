package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	// opcodes
	PUSH   byte = 0xF0
	POP    byte = 0xF1
	RETURN byte = 0xF2
	CALL   byte = 0xF3
	SFRAME byte = 0xF4 //create stack frame
	SUM    byte = 0xA0
	DIV    byte = 0xA1
	MUL    byte = 0xA2

	MODSEC  byte = 0xAD //modules section
	ADDMOD  byte = 0xD0 // add mod to loader
	MODNEND byte = 0xD1
	MDSEND  byte = 0xDA // end module section

	//error codes

	SUCCESSEVAL    int = 0xDEADD011
	NOTBYTECODE    int = 0xBADC0DE
	MODSECFINDFAIL int = 0xFA11EDF1
)

var ip int = 0
var bytecode []byte

var stack [128]int
var sp int = 127
var bp int = 127

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
	// if ip := findImportSection(); ip == 0 {
	// 	return MODSECFINDFAIL
	// }
	for ; ip != len(bytecode); ip++ {
		fmt.Printf("\n\n INSTRUCTION: 0x%02x\n\n", bytecode[ip])
		switch bytecode[ip] {
		case PUSH:
			fmt.Println("PUSH command")
			ip++
			push(readWordX32(ip))
		case SUM:
			sum()
		}
	}
	debugMode()

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
		ip++
		word += int(bytecode[start+i]) << offset
		offset -= 8
	}
	fmt.Printf("word: 0x%02x\n", word)
	return word
}

//STACK OPERATING

func push(value int) {
	stack[sp] = value
	fmt.Printf("\npush: 0x%02x\n", stack[sp])
	sp--
}

func pop() int {
	value := stack[sp]
	sp++
	return value
}

func sum() {
	leftop := stack[sp]
	fmt.Println(leftop)
	rightop := stack[sp+1]
	fmt.Println(rightop)
	push(leftop + rightop)
	fmt.Print(pop())
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

func debugMode() {
	for stackpose, stackvalue := range stack {
		fmt.Printf("\npose: %d, value: 0x%02x", stackpose, stackvalue)
	}
}
