package processor_test

import (
	"testing"

	"github.com/garysmi/go-8080/processor"
)

func TestJmp(t *testing.T) {
	proc := processor.NewProc()
	proc.PC = 0x00
	proc.MEMORY[proc.PC+1] = 0xbb
	proc.MEMORY[proc.PC+2] = 0xaa

	proc.Jmp()

	result := proc.PC
	if result != 0xaabb {
		t.Errorf("proc.Jmp() failed, expected %#x, got %#x", 0xaabb, result)
	} else {
		t.Logf("proc.Jmp() success, expected %#x, got %#x", 0xaabb, result)
	}
}
