package processor

// Dad - Double add the contents of the specified regiser pair to HL pair
func (cpu *Processor) Dad(r string) {
	switch r {
	case "D":
		d := (cpu.D << 8) | cpu.E
		h := (cpu.H << 8) | cpu.L
		res := h + d
		cpu.H = uint8((res >> 8) & 0xff)
		cpu.L = uint8(res & 0xff)
		if res > 0xff {
			cpu.SetCarry()
		}
	}
	cpu.PC++
}

// Dcx - Decrement specified register pair by 1
func (cpu *Processor) Dcx(r string) {
	switch r {
	case "B":
		cpu.B--
		cpu.C--
	case "D":
		cpu.D--
		cpu.E--
	case "H":
		cpu.H--
		cpu.L--
	case "SP":
		cpu.SP--
	}
	cpu.PC++
}
