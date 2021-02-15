package processor

import "fmt"

// Processor models an 8080 Processor
type Processor struct {
	A      uint8
	B      uint8
	C      uint8
	D      uint8
	E      uint8
	H      uint8
	L      uint8
	PC     uint16
	SP     uint16
	FLAGS  uint8
	MEMORY []uint8
}

// NewProc creates a new instance of processor
func NewProc() *Processor {
	return &Processor{MEMORY: make([]uint8, 0xffff), FLAGS: 0x02}
}

// SetCarry sets the carry bit
func (cpu *Processor) SetCarry() {
	cpu.FLAGS = cpu.FLAGS ^ 0b00000001
}

// SetParity sets the parity bit
func (cpu *Processor) SetParity() {
	cpu.FLAGS = cpu.FLAGS ^ 0b00000100
}

// SetAuxCarry sets the Auxcillary Carry Bit
func (cpu *Processor) SetAuxCarry() {
	cpu.FLAGS = cpu.FLAGS ^ 0b00010000
}

// SetZero sets the Zero bit
func (cpu *Processor) SetZero() {
	cpu.FLAGS = cpu.FLAGS ^ 0b01000000
}

// SetSign sets the Sign bit
func (cpu *Processor) SetSign() {
	cpu.FLAGS = cpu.FLAGS ^ 0b1000000
}

// CheckCarry returns the current state of the carry bit
func (cpu *Processor) CheckCarry() uint8 {
	return cpu.FLAGS & 0b00000001
}

// CheckParity returns the current state of the Parity bit
func (cpu *Processor) CheckParity() uint8 {
	return cpu.FLAGS & 0b00000100
}

// CheckAuxCarry returns the current state of the Auxcillary Carry bit
func (cpu *Processor) CheckAuxCarry() uint8 {
	return cpu.FLAGS & 0b00010000
}

// CheckZero returns the current state of the Zero bit
func (cpu *Processor) CheckZero() uint8 {
	return cpu.FLAGS & 0b01000000
}

// CheckSign returns the current state of the Sign bit
func (cpu *Processor) CheckSign() uint8 {
	return cpu.FLAGS & 0b1000000
}

// ResetCarry resets the carry bit if set
func (cpu *Processor) ResetCarry() {
	if cpu.CheckCarry() == 1 {
		cpu.SetCarry()
	}
}

// ResetParity resets the Parity bit if set
func (cpu *Processor) ResetParity() {
	if cpu.CheckParity() == 1 {
		cpu.SetParity()
	}
}

// ResetAuxCarry resets the Auxcillary Carry bit if set
func (cpu *Processor) ResetAuxCarry() {
	if cpu.CheckAuxCarry() == 1 {
		cpu.SetAuxCarry()
	}
}

// ResetZero resets the Zero bit if set
func (cpu *Processor) resetZero() {
	if cpu.CheckZero() == 1 {
		cpu.SetZero()
	}
}

// ResetSign resets the Sign bit if set
func (cpu *Processor) resetSign() {
	cpu.FLAGS = cpu.FLAGS ^ 0b1000000
}

// PrintState prints the current state of the CPU for debugging
func (cpu *Processor) PrintState() {
	fmt.Printf("Register B: %#x\n", cpu.B)
	fmt.Printf("Register C: %#x\n", cpu.C)
	fmt.Printf("Register D: %#x\n", cpu.D)
	fmt.Printf("Register E: %#x\n", cpu.E)
	fmt.Printf("Register H: %#x\n", cpu.H)
	fmt.Printf("Register L: %#x\n", cpu.L)
	fmt.Printf("Register A: %#x\n", cpu.A)
	fmt.Printf("Register PC: %#x\n", cpu.PC)
	fmt.Printf("Register SP: %#x\n", cpu.SP)
	fmt.Printf("Register FLAGS: %#x\n", cpu.FLAGS)
	// fmt.Printf("MEMORY: %#x\n", cpu.MEMORY)
}

//Parity returns true if the number of set bits is odd
func (cpu *Processor) Parity(x uint8) bool {
	x ^= (x >> 4)
	x ^= (x >> 2)
	x ^= (x >> 1)

	if x&0x01 == 0 {
		return true
	}

	return false
}
