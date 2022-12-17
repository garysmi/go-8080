package processor_test

import (
	"testing"

	"github.com/garysmi/go-8080/processor"
)

func TestInx(t *testing.T) {

	proc := processor.NewProc()

	proc.Inx("B")

	if proc.B != 0x01 && proc.C != 0x01 {
		t.Errorf("cpu.Inx() failed, expected %#x and %#x, got %#x and %#x", 0x01, 0x01, proc.B, proc.C)
	} else {
		t.Logf("cpu.Inx() success, expected  %#x and %#x, got %#x and %#x", 0x01, 0x01, proc.B, proc.C)
	}

	proc.Inx("D")

	if proc.B != 0x01 && proc.C != 0x01 {
		t.Errorf("cpu.Inx() failed, expected  %#x and %#x, got %#x and %#x", 0x01, 0x01, proc.D, proc.E)
	} else {
		t.Logf("cpu.Inx() success, expected  %#x and %#x, got %#x and %#x", 0x01, 0x01, proc.D, proc.E)
	}

	proc.Inx("H")

	if proc.H != 0x01 && proc.L != 0x01 {
		t.Errorf("cpu.Inx() failed, expected  %#x and %#x, got %#x and %#x", 0x01, 0x01, proc.H, proc.L)
	} else {
		t.Logf("cpu.Inx() success, expected  %#x and %#x, got %#x and %#x", 0x01, 0x01, proc.H, proc.L)
	}

	proc.Inx("SP")

	if proc.SP != 0x01 {
		t.Errorf("cpu.Inx() failed, expected %#x, got %#x", 0x01, proc.SP)
	} else {
		t.Logf("cpu.Inx() success, expected %#x, got %#x", 0x01, proc.SP)
	}
}

func TestDcx(t *testing.T) {

	proc := processor.NewProc()

	proc.Dcx("B")

	if proc.B != 0xff && proc.C != 0xff {
		t.Errorf("cpu.Inx() failed, expected %#x and %#x, got %#x and %#x", 0xff, 0xff, proc.B, proc.C)
	} else {
		t.Logf("cpu.Inx() success, expected  %#x and %#x, got %#x and %#x", 0xff, 0xff, proc.B, proc.C)
	}

	proc.Dcx("D")

	if proc.B != 0xff && proc.C != 0xff {
		t.Errorf("cpu.Inx() failed, expected  %#x and %#x, got %#x and %#x", 0xff, 0xff, proc.D, proc.E)
	} else {
		t.Logf("cpu.Inx() success, expected  %#x and %#x, got %#x and %#x", 0xff, 0xff, proc.D, proc.E)
	}

	proc.Dcx("H")

	if proc.H != 0xff && proc.L != 0xff {
		t.Errorf("cpu.Inx() failed, expected  %#x and %#x, got %#x and %#x", 0xff, 0xff, proc.H, proc.L)
	} else {
		t.Logf("cpu.Inx() success, expected  %#x and %#x, got %#x and %#x", 0xff, 0xff, proc.H, proc.L)
	}

	proc.Dcx("SP")

	if proc.SP != 0xffff {
		t.Errorf("cpu.Inx() failed, expected %#x, got %#x", 0xffff, proc.SP)
	} else {
		t.Logf("cpu.Inx() success, expected %#x, got %#x", 0xffff, proc.SP)
	}
}
