package opcode

const (
	// opcodes
	PUSHBYTE  byte = 0xC0
	PUSHINT   byte = 0xC1
	PUSHFLOAT byte = 0xC2
	PUSHREF   byte = 0xC3
	IFSTAT    byte = 0xC4
	ELSE      byte = 0xC5

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
)
