package processor

// Ora - Or the specifed register with the accumlator
func (cpu *Processor) Ora(r string) {
	var res uint8
	switch r {
	case "B":
		res = cpu.A | cpu.B
	case "C":
		res = cpu.A | cpu.C
	case "H":
		res = cpu.A | cpu.H
	}

	cpu.A = res

	if res > 0xff {
		cpu.SetCarry()
	}

	if cpu.A&0b1000000 == 0 {
		cpu.SetSign()
	} else {
		cpu.ResetSign()
	}

	if cpu.A == 0x00 {
		cpu.SetZero()
	} else {
		cpu.ResetZero()
	}

	if parity(res) {
		cpu.SetParity()
	} else {
		cpu.ResetParity()
	}

	cpu.PC++
}
