package processor

// Jmp to memory location addr
func (cpu *Processor) Jmp() {
	addr := int(0x00)
	addr = addr | int(cpu.MEMORY[cpu.PC+0x02])
	addr = addr << 8
	addr = addr | int(cpu.MEMORY[cpu.PC+0x01])
	cpu.PC = uint16(addr)
}
