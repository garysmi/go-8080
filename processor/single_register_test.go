package processor_test

import (
	"testing"

	"github.com/garysmi/go-8080/processor"
)

func TestDcr(t *testing.T) {

	proc := processor.NewProc()

	proc.Dcr(&proc.A)

	if proc.A != 0xff {
		t.Errorf("cpu.Inx() failed, expected %#x, got %#x", 0xff, proc.A)
	} else {
		t.Logf("cpu.Inx() success, expected %#x, got %#x", 0xff, proc.A)
	}
}
