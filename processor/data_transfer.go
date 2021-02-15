package processor

import (
	"reflect"
)

// Ldax - loads data from memory address by register pair into accumulator
func (cpu *Processor) Ldax(r string) {
	switch r {
	case "B":
		cpu.A = cpu.MEMORY[makeUint16(cpu.B, cpu.C)]
	case "D":
		cpu.A = cpu.MEMORY[makeUint16(cpu.D, cpu.E)]
	}

	cpu.PC++
}

// Mov - move the contents of the src into dst
func (cpu *Processor) Mov(dst interface{}, src interface{}) {
	srcT := reflect.TypeOf(src).String()
	dstT := reflect.TypeOf(dst).String()

	if srcT == "[]uint8" {
		d, _ := dst.(*uint8)
		s, _ := src.([]uint8)
		*d = s[makeUint16(cpu.H, cpu.L)]
	} else if dstT == "*[]uint8" {
		d, _ := dst.(*[]uint8)
		s, _ := src.(uint8)
		(*d)[makeUint16(cpu.H, cpu.L)] = s
	} else {
		d, _ := dst.(*uint8)
		s, _ := src.(uint8)
		*d = s
	}

	cpu.PC++
}
