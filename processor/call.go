package processor

// Call a subroutine at memory location
func (cpu *Processor) Call() {
	ret := cpu.PC + 0x03
	cpu.MEMORY[cpu.SP-1] = uint8((ret >> 8) & 0xff)
	cpu.MEMORY[cpu.SP-2] = uint8(ret & 0xff)
	cpu.SP = cpu.SP - 2
	var addr int
	addr = addr | int(cpu.MEMORY[cpu.PC+0x02])
	addr = (addr << 8) | int(cpu.MEMORY[cpu.PC+0x01])
	cpu.PC = uint16(addr)
}

// Cnz call the subroutine if zero flag is 1
func (cpu *Processor) Cnz() {
	if cpu.CheckZero() != 0 {
		cpu.PC = cpu.PC + 0x02
	} else {
		cpu.PC = cpu.PC + 0x02
	}
}
