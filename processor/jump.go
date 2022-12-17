package processor

// Jmp to memory location addr
func (cpu *Processor) Jmp() {
	addr := makeUint16(cpu.MEMORY[cpu.PC+0x02], cpu.MEMORY[cpu.PC+0x01])
	cpu.PC = uint16(addr)
}
