package disassembler

import (
	"fmt"
)

// Disassemble takes hex and outputs 8080 assemlby code
func Disassemble(bytes []uint8) {
	for i := 0; i < len(bytes); i++ {
		if i == 0 {
			fmt.Printf("foo")
		}
	}
}
