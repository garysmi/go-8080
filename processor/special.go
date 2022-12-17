package processor

// Nop increments Program count by 1
func (cpu *Processor) Nop() {
	cpu.PC++
}

// Out Sends contents of register to output device
func (cpu *Processor) Out() {
	cpu.PC = cpu.PC + 0x02
}
