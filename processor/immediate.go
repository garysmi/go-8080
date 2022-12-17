package processor

//Lxi Loads immediate data into register pair
func (cpu *Processor) Lxi(rp string) {
	switch rp {
	case "SP":
		data := int(0x00)
		data = data | int(cpu.MEMORY[cpu.PC+0x02])
		data = data << 8
		data = data | int(cpu.MEMORY[cpu.PC+0x01])
		cpu.SP = uint16(data)
	case "B":
		cpu.B = cpu.MEMORY[cpu.PC+0x02]
		cpu.C = cpu.MEMORY[cpu.PC+0x01]
	case "D":
		cpu.D = cpu.MEMORY[cpu.PC+0x02]
		cpu.E = cpu.MEMORY[cpu.PC+0x01]
	case "H":
		cpu.H = cpu.MEMORY[cpu.PC+0x02]
		cpu.L = cpu.MEMORY[cpu.PC+0x01]
	}

	cpu.PC = cpu.PC + 0x03
}

//Mvi Loads immediate data into register
func (cpu *Processor) Mvi(r string) {
	switch r {
	case "A":
		cpu.A = cpu.MEMORY[cpu.PC+0x01]
	case "B":
		cpu.B = cpu.MEMORY[cpu.PC+0x01]
	case "C":
		cpu.C = cpu.MEMORY[cpu.PC+0x01]
	case "D":
		cpu.D = cpu.MEMORY[cpu.PC+0x01]
	case "E":
		cpu.E = cpu.MEMORY[cpu.PC+0x01]
	case "H":
		cpu.H = cpu.MEMORY[cpu.PC+0x01]
	case "L":
		cpu.L = cpu.MEMORY[cpu.PC+0x01]
	case "M":
		addr := makeUint16(cpu.H, cpu.L)
		cpu.MEMORY[addr] = cpu.MEMORY[cpu.PC+0x01]
	}
	cpu.PC = cpu.PC + 0x02
}

// Sta Store accumlator to memory location addr
func (cpu *Processor) Sta() {
	addr := makeUint16(cpu.MEMORY[cpu.PC+0x02], cpu.MEMORY[cpu.PC+0x01])
	cpu.MEMORY[addr] = cpu.A
	cpu.PC = cpu.PC + 0x03
}

// Xra xor registrer pair with accumulator
func (cpu *Processor) Xra() {
	cpu.ResetCarry()
	cpu.PC++
}
