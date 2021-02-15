package processor_test

import (
	"testing"

	"github.com/garysmi/go-8080/processor"
)

func TestOra(t *testing.T) {
	proc := processor.NewProc()
	proc.A = 0x33
	proc.C = 0x0F

	proc.Ora("C")

	if proc.A != 0x3F {
		t.Errorf("proc.Ora() failed, expected %#x, got %#x", 0x3F, proc.A)
	} else {
		t.Logf("proc.Ldax() success, expected %#x, got %#x", 0xbb, proc.A)
	}

}
