package processor

import (
	"reflect"
)

// Dcr - decrement register or memory location by 1
func (cpu *Processor) Dcr(r interface{}) {

	if reflect.TypeOf(r).String() == "*[]unint*" {
		r, _ := r.(*[]uint8)
		(*r)[makeUint16(cpu.H, cpu.L)]--
		cpu.Status_flags((*r)[makeUint16(cpu.H, cpu.L)])
	} else {
		r, _ := r.(*uint8)
		*r--
		cpu.Status_flags(*r)
	}

	cpu.PC++
}

// Inr - Increment register or memory location by 1
func (cpu *Processor) Inr(r interface{}) {
	if reflect.TypeOf(r).String() == "*[]unint*" {
		r, _ := r.(*[]uint8)
		(*r)[makeUint16(cpu.H, cpu.L)]--
		cpu.Status_flags((*r)[makeUint16(cpu.H, cpu.L)])
	} else {
		r, _ := r.(*uint8)
		*r--
		cpu.Status_flags(*r)
	}

	cpu.PC++
}
