package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	// opcodes
	ENDWRD int = 0x0F
	PUSH   int = 0xF6
	POP    int = 0xF5
	RETURN int = 0xFD
	CALL   int = 0xFA
	SUM    int = 0xA0

	MODSEC int = 0xAD //modules section
	ADDMOD int = 0xDD // add mod to loader
	MDSEND int = 0xDA // end module section

	//error codes

	SUCCESSEVAL    int = 0xDEADD011
	NOTBYTECODE    int = 0xBADC0DE
	MODSECFINDFAIL int = 0xFA11EDF1
)

var ip int = 4
var bytecode []byte

func main() {
	fmt.Println("Petlang v0.1")
	bytecode = loadByteCode(os.Args[1])
	errcode := evalByteCode()
	if errcode != SUCCESSEVAL {
		fmt.Printf("Petlang vm error, code: 0x%02x\n", errcode)
		switch errcode {
		case NOTBYTECODE:
			fmt.Printf("file isn't petlang bytecode\n")
		case MODSECFINDFAIL:

		}
	}

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
