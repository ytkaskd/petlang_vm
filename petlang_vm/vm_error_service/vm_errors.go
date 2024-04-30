package vm_errors

import (
	"fmt"
	"os"
)

// error codes
const (
	NOTBYTECODE    int = 0xBADC0DE
	MODSECFINDFAIL int = 0xFA11EDF1
	UNKNOWNOPCODE  int = 0xCAFEBABE
)

func ThrowError(errcode int, point int) {
	fmt.Printf("\nPetlang vm error, code: 0x%02x\nat point: %d\n", errcode, point)
	switch errcode {
	case NOTBYTECODE:
		fmt.Println("file isn't petlang bytecode")
	case MODSECFINDFAIL:
		fmt.Println("Can't find modules import section")
	case UNKNOWNOPCODE:
		fmt.Println("Can't recognize bytecode")
	}

	os.Exit(1)
}
