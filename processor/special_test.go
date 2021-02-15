package processor

import (
	"testing"
)

func TestNop(t *testing.T) {
	proc := new(Processor)
	proc.PC = 0x00
	proc.Nop()

	result := proc.PC
	if result != 0x01 {
		t.Errorf("proc.nop() failed, expected %v, got %v", 0x01, result)
	} else {
		t.Logf("proc.nop() success, expected %v, got %v", 0x01, result)
	}
}
