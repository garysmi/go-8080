package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/garysmi/go-8080/processor"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("testdata/invaders")
	checkErr(err)
	defer file.Close()

	checkErr(err)

	cpu := processor.NewProc()

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(cpu.MEMORY)

	cpu.PC = uint16(cpu.MEMORY[0x00])

	for {
		opcode := cpu.MEMORY[cpu.PC]
		fmt.Printf("Intruction %#02x @ %#06x\n", opcode, cpu.PC)

		switch opcode {
		case 0x00:
			cpu.Nop()
		case 0x01:
			cpu.Lxi("B")
		case 0x03:
			cpu.Inx("B")
		case 0x05:
			cpu.Dcr(&cpu.B)
		case 0x06:
			cpu.Mvi("B")
		case 0x0a:
			cpu.Ldax("B")
		case 0x0b:
			cpu.Dcx("B")
		case 0x011:
			cpu.Lxi("D")
		case 0x13:
			cpu.Inx("D")
		case 0x19:
			cpu.Dad("D")
		case 0x1a:
			cpu.Ldax("D")
		case 0x1b:
			cpu.Dcx("D")
		case 0x21:
			cpu.Lxi("H")
		case 0x23:
			cpu.Inx("H")
		case 0x2b:
			cpu.Dcx("B")
		case 0x33:
			cpu.Inx("SP")
		case 0x36:
			cpu.Mvi("M")
		case 0x3b:
			cpu.Dcx("SP")
		case 0x40:
			cpu.Mov(&cpu.B, cpu.B)
		case 0x41:
			cpu.Mov(&cpu.B, cpu.C)
		case 0x42:
			cpu.Mov(&cpu.B, cpu.D)
		case 0x43:
			cpu.Mov(&cpu.B, cpu.E)
		case 0x44:
			cpu.Mov(&cpu.B, cpu.H)
		case 0x45:
			cpu.Mov(&cpu.B, cpu.L)
		case 0x46:
			cpu.Mov(&cpu.B, cpu.MEMORY)
		case 0x47:
			cpu.Mov(&cpu.B, cpu.A)
		case 0x48:
			cpu.Mov(&cpu.C, cpu.B)
		case 0x49:
			cpu.Mov(&cpu.C, cpu.C)
		case 0x4a:
			cpu.Mov(&cpu.C, cpu.D)
		case 0x4b:
			cpu.Mov(&cpu.C, cpu.E)
		case 0x4c:
			cpu.Mov(&cpu.C, cpu.H)
		case 0x4d:
			cpu.Mov(&cpu.C, cpu.L)
		case 0x4e:
			cpu.Mov(&cpu.C, cpu.MEMORY)
		case 0x4f:
			cpu.Mov(&cpu.C, cpu.A)
		case 0x50:
			cpu.Mov(&cpu.D, cpu.B)
		case 0x51:
			cpu.Mov(&cpu.D, cpu.C)
		case 0x52:
			cpu.Mov(&cpu.D, cpu.D)
		case 0x53:
			cpu.Mov(&cpu.D, cpu.E)
		case 0x54:
			cpu.Mov(&cpu.D, cpu.H)
		case 0x55:
			cpu.Mov(&cpu.D, cpu.L)
		case 0x56:
			cpu.Mov(&cpu.D, cpu.MEMORY)
		case 0x57:
			cpu.Mov(&cpu.D, cpu.A)
		case 0x58:
			cpu.Mov(&cpu.E, cpu.B)
		case 0x59:
			cpu.Mov(&cpu.E, cpu.C)
		case 0x5a:
			cpu.Mov(&cpu.E, cpu.D)
		case 0x5b:
			cpu.Mov(&cpu.E, cpu.E)
		case 0x5c:
			cpu.Mov(&cpu.E, cpu.H)
		case 0x5d:
			cpu.Mov(&cpu.E, cpu.L)
		case 0x5e:
			cpu.Mov(&cpu.E, cpu.MEMORY)
		case 0x5f:
			cpu.Mov(&cpu.E, cpu.A)
		case 0x60:
			cpu.Mov(&cpu.H, cpu.B)
		case 0x61:
			cpu.Mov(&cpu.H, cpu.C)
		case 0x62:
			cpu.Mov(&cpu.H, cpu.D)
		case 0x63:
			cpu.Mov(&cpu.H, cpu.E)
		case 0x64:
			cpu.Mov(&cpu.H, cpu.H)
		case 0x65:
			cpu.Mov(&cpu.H, cpu.L)
		case 0x66:
			cpu.Mov(&cpu.H, cpu.MEMORY)
		case 0x67:
			cpu.Mov(&cpu.H, cpu.A)
		case 0x68:
			cpu.Mov(&cpu.L, cpu.B)
		case 0x69:
			cpu.Mov(&cpu.L, cpu.C)
		case 0x6a:
			cpu.Mov(&cpu.L, cpu.D)
		case 0x6b:
			cpu.Mov(&cpu.L, cpu.E)
		case 0x6c:
			cpu.Mov(&cpu.L, cpu.H)
		case 0x6d:
			cpu.Mov(&cpu.L, cpu.L)
		case 0x6e:
			cpu.Mov(&cpu.L, cpu.MEMORY)
		case 0x6f:
			cpu.Mov(&cpu.L, cpu.A)
		case 0x70:
			cpu.Mov(&cpu.MEMORY, cpu.B)
		case 0x71:
			cpu.Mov(&cpu.MEMORY, cpu.C)
		case 0x72:
			cpu.Mov(&cpu.MEMORY, cpu.D)
		case 0x73:
			cpu.Mov(&cpu.MEMORY, cpu.E)
		case 0x74:
			cpu.Mov(&cpu.MEMORY, cpu.H)
		case 0x75:
			cpu.Mov(&cpu.MEMORY, cpu.L)
		case 0x77:
			cpu.Mov(&cpu.MEMORY, cpu.A)
		case 0x78:
			cpu.Mov(&cpu.A, cpu.B)
		case 0x79:
			cpu.Mov(&cpu.A, cpu.C)
		case 0x7a:
			cpu.Mov(&cpu.A, cpu.D)
		case 0x7b:
			cpu.Mov(&cpu.A, cpu.E)
		case 0x7c:
			cpu.Mov(&cpu.A, cpu.H)
		case 0x7d:
			cpu.Mov(&cpu.A, cpu.L)
		case 0x7e:
			cpu.Mov(&cpu.A, cpu.MEMORY)
		case 0x7f:
			cpu.Mov(&cpu.A, cpu.A)
		case 0xaf:
			cpu.Xra()
		case 0xb0:
			cpu.Ora("B")
		case 0xb1:
			cpu.Ora("C")
		case 0xb4:
			cpu.Ora("H")
		case 0x31:
			cpu.Lxi("SP")
		case 0x32:
			cpu.Sta()
		case 0x3e:
			cpu.Mvi("A")
		case 0xc3:
			cpu.Jmp()
		case 0xc9:
			cpu.Ret()
		case 0xcd:
			cpu.Call()
		case 0xc2:
			cpu.Cnz()
		case 0xd3:
			cpu.Out()
		default:
			fmt.Printf("Uinplemented Intruction %#02x\n", opcode)
			cpu.PrintState()
			os.Exit(1)
		}

	}
}
