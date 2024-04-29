package vm_errors

import (
	"fmt"
	"os"
)

// error codes
const (
	NOTBYTECODE    int = 0xBADC0DE
	MODSECFINDFAIL int = 0xFA11EDF1
)

func ThrowError(errcode int) {
	fmt.Printf("\nPetlang vm error, code: 0x%02x\n", errcode)
	switch errcode {
	case NOTBYTECODE:
		fmt.Printf("file isn't petlang bytecode\n")
	case MODSECFINDFAIL:
		fmt.Printf("Can't find modules import section\n")
	}

	os.Exit(1)
}
