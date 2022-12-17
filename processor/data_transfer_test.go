package processor_test

import (
	"testing"

	"github.com/garysmi/go-8080/processor"
)

func TestLdax(t *testing.T) {
	proc := processor.NewProc()
	proc.A = 0xaa
	proc.B = 0x93
	proc.C = 0x8B
	proc.D = 0xDE
	proc.E = 0xAD
	proc.MEMORY[0x938B] = 0xFF
	proc.MEMORY[0xDEAD] = 0xbb

	proc.Ldax("B")

	result := proc.A
	if result != 0xFF {
		t.Errorf("proc.Ldax() failed, expected %#x, got %#x", 0xFF, result)
	} else {
		t.Logf("proc.Ldax() success, expected %#x, got %#x", 0xFF, result)
	}

	proc.Ldax("D")

	result = proc.A
	if result != 0xbb {
		t.Errorf("proc.Ldax() failed, expected %#x, got %#x", 0xbb, result)
	} else {
		t.Logf("proc.Ldax() success, expected %#x, got %#x", 0xbb, result)
	}
}
