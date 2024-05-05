package vm_errors

import (
	"fmt"
	"os"
)

// error codes
const (
	NOTBYTECODE byte = iota + 1
	MODSECFINDFAIL
	UNKNOWNOPCODE
	STACKOVERFLOW
	CASTERROR
	UNSUPADAR
)

func ThrowError(errcode byte, point int) {
	fmt.Printf("\033[91m\nPetlang vm error, code: 0x%02x\nat point: %d\n", errcode, point)
	switch errcode {
	case NOTBYTECODE:
		fmt.Println("file isn't petlang bytecode")
	case MODSECFINDFAIL:
		fmt.Println("Can't find modules import section")
	case UNKNOWNOPCODE:
		fmt.Println("Can't recognize bytecode")
	case STACKOVERFLOW:
		fmt.Printf("\nFatal error: stack overflow! sp = %d\n", point)
	case UNSUPADAR:
		fmt.Printf("\nAddress arithmetic unsupported")
	default:
		fmt.Printf("Unknown error")
	}
	fmt.Print("\033[0m")

	os.Exit(1)
}
