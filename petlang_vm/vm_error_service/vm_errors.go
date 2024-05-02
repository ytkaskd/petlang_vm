package vm_errors

import (
	"fmt"
	"os"
)

// error codes
const (
	NOTBYTECODE    byte = 0x00
	MODSECFINDFAIL byte = 0x01
	UNKNOWNOPCODE  byte = 0x02
	STACKOVERFLOW  byte = 0x03
)

func ThrowError(errcode byte, point int) {
	fmt.Printf("\nPetlang vm error, code: 0x%02x\nat point: %d\n", errcode, point)
	switch errcode {
	case NOTBYTECODE:
		fmt.Println("file isn't petlang bytecode")
	case MODSECFINDFAIL:
		fmt.Println("Can't find modules import section")
	case UNKNOWNOPCODE:
		fmt.Println("Can't recognize bytecode")
	case STACKOVERFLOW:
		fmt.Printf("\nFatal error: stack overflow! sp = %d\n", point)
	}

	os.Exit(1)
}
