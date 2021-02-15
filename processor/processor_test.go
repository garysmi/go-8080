package processor_test

import (
	"testing"

	"github.com/garysmi/go-8080/processor"
)

func TestSetCary(t *testing.T) {
	proc := processor.NewProc()

	proc.SetCarry()
	result := proc.CheckCarry()

	if result != 0x01 {
		t.Errorf("cpu.setCarry() failed, expected %#x, got %#x", 0x01, result)
	} else {
		t.Logf("cpu.setCarry() success, expected %#x, got %#x", 0x01, result)
	}

	proc.SetCarry()
	result = proc.CheckCarry()

	if result != 0x00 {
		t.Errorf("cpu.setCarry() failed, expected %#x, got %#x", 0x00, result)
	} else {
		t.Logf("cpu.setCarry() success, expected %#x, got %#x", 0x00, result)
	}

}

func TestParity(t *testing.T) {
	proc := processor.NewProc()

	result := proc.Parity(0xAB)

	if !result {
		t.Logf("cpu.Parity() success, expected %v, got %v", false, result)
	} else {
		t.Errorf("cpu.setCarry() failed, expected %v, got %v", false, result)
	}

	result = proc.Parity(0xAA)

	if result {
		t.Logf("cpu.Parity() success, expected %v, got %v", true, result)
	} else {
		t.Errorf("cpu.setCarry() failed, expected %v, got %v", true, result)
	}
}
