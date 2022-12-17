package processor

import "fmt"

// Ret - Pops the last instruction of the stack and executes it
func (cpu *Processor) Ret() {
	addr := makeUint16(cpu.MEMORY[cpu.SP+1], cpu.MEMORY[cpu.SP])
	fmt.Printf("RET ADDRESS: %#x\n", addr)
	cpu.SP += 2
	cpu.PC = addr
}
