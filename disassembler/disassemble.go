package main

import (
	"fmt"
)

// Disassemble takes hex and outputs 8080 assemlby code
func Disassemble(bytes []byte) {
	fmt.Printf("% x\n", bytes)
	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		case 0x00:
			fmt.Printf("%04x %#02x NOP\n", i, bytes[i])
		case 0x01:
			fmt.Printf("%04x %#02x LXI B,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x02:
			fmt.Printf("%04x %#02x STAX B\n", i, bytes[i])
		case 0x03:
			fmt.Printf("%04x %#02x INX B\n", i, bytes[i])
		case 0x04:
			fmt.Printf("%04x %#02x INR B\n", i, bytes[i])
		case 0x05:
			fmt.Printf("%04x %#02x DCR B\n", i, bytes[i])
		case 0x06:
			fmt.Printf("%04x %#02x MVI B,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x07:
			fmt.Printf("%04x %#02x RLC\n", i, bytes[i])
		case 0x09:
			fmt.Printf("%04x %#02x DAD B\n", i, bytes[i])
		case 0x0a:
			fmt.Printf("%04x %#02x LDAX B\n", i, bytes[i])
		case 0x0b:
			fmt.Printf("%04x %#02x DCX B\n", i, bytes[i])
		case 0x0c:
			fmt.Printf("%04x %#02x INR C\n", i, bytes[i])
		case 0x0d:
			fmt.Printf("%04x %#02x DCR C\n", i, bytes[i])
		case 0x0e:
			fmt.Printf("%04x %#02x MVI C,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x0f:
			fmt.Printf("%04x %#02x RRC\n", i, bytes[i])
		case 0x11:
			fmt.Printf("%04x %#02x LXI D,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x12:
			fmt.Printf("%04x %#02x STAX D\n", i, bytes[i])
		case 0x13:
			fmt.Printf("%04x %#02x INX D\n", i, bytes[i])
		case 0x14:
			fmt.Printf("%04x %#02x INR D\n", i, bytes[i])
		case 0x15:
			fmt.Printf("%04x %#02x DCR D\n", i, bytes[i])
		case 0x16:
			fmt.Printf("%04x %#02x MVI D,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x17:
			fmt.Printf("%04x %#02x RAL\n", i, bytes[i])
		case 0x19:
			fmt.Printf("%04x %#02x DAD D\n", i, bytes[i])
		case 0x1a:
			fmt.Printf("%04x %#02x LDAX D\n", i, bytes[i])
		case 0x1b:
			fmt.Printf("%04x %#02x DCX D\n", i, bytes[i])
		case 0x1c:
			fmt.Printf("%04x %#02x INR E\n", i, bytes[i])
		case 0x1d:
			fmt.Printf("%04x %#02x DCR E\n", i, bytes[i])
		case 0x1e:
			fmt.Printf("%04x %#02x MVI E,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x1f:
			fmt.Printf("%04x %#02x RAR\n", i, bytes[i])
		case 0x21:
			fmt.Printf("%04x %#02x LXI H,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x22:
			fmt.Printf("%04x %#02x SHLD,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x23:
			fmt.Printf("%04x %#02x INX H\n", i, bytes[i])
		case 0x24:
			fmt.Printf("%04x %#02x INR H\n", i, bytes[i])
		case 0x25:
			fmt.Printf("%04x %#02x DCR H\n", i, bytes[i])
		case 0x26:
			fmt.Printf("%04x %#02x MVI H,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x27:
			fmt.Printf("%04x %#02x DAA\n", i, bytes[i])
		case 0x29:
			fmt.Printf("%04x %#02x DAD H\n", i, bytes[i])
		case 0x2a:
			fmt.Printf("%04x %#02x LHLD,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x2c:
			fmt.Printf("%04x %#02x INR L\n", i, bytes[i])
		case 0x2b:
			fmt.Printf("%04x %#02x DCX H\n", i, bytes[i])
		case 0x2d:
			fmt.Printf("%04x %#02x DCR L\n", i, bytes[i])
		case 0x2e:
			fmt.Printf("%04x %#02x MVI L,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x2f:
			fmt.Printf("%04x %#02x CMA\n", i, bytes[i])
		case 0x31:
			fmt.Printf("%04x %#02x LXI SP,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x32:
			fmt.Printf("%04x %#02x STA,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x33:
			fmt.Printf("%04x %#02x INX SP\n", i, bytes[i])
		case 0x34:
			fmt.Printf("%04x %#02x INR M\n", i, bytes[i])
		case 0x35:
			fmt.Printf("%04x %#02x DCR M\n", i, bytes[i])
		case 0x36:
			fmt.Printf("%04x %#02x MVI M,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x37:
			fmt.Printf("%04x %#02x STC\n", i, bytes[i])
		case 0x39:
			fmt.Printf("%04x %#02x DAD SP\n", i, bytes[i])
		case 0x3a:
			fmt.Printf("%04x %#02x LDA,%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0x3b:
			fmt.Printf("%04x %#02x DCX SP\n", i, bytes[i])
		case 0x3c:
			fmt.Printf("%04x %#02x INR A\n", i, bytes[i])
		case 0x3d:
			fmt.Printf("%04x %#02x DCR A\n", i, bytes[i])
		case 0x3e:
			fmt.Printf("%04x %#02x MVI A,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0x3F:
			fmt.Printf("%04x %#02x CMC\n", i, bytes[i])
		case 0x40:
			fmt.Printf("%04x %#02x MOV B,B\n", i, bytes[i])
		case 0x41:
			fmt.Printf("%04x %#02x MOV B,C\n", i, bytes[i])
		case 0x42:
			fmt.Printf("%04x %#02x MOV B,D\n", i, bytes[i])
		case 0x43:
			fmt.Printf("%04x %#02x MOV B,E\n", i, bytes[i])
		case 0x44:
			fmt.Printf("%04x %#02x MOV B,H\n", i, bytes[i])
		case 0x45:
			fmt.Printf("%04x %#02x MOV B,L\n", i, bytes[i])
		case 0x46:
			fmt.Printf("%04x %#02x MOV B,M\n", i, bytes[i])
		case 0x47:
			fmt.Printf("%04x %#02x MOV B,A\n", i, bytes[i])
		case 0x48:
			fmt.Printf("%04x %#02x MOV C,B\n", i, bytes[i])
		case 0x49:
			fmt.Printf("%04x %#02x MOV C,C\n", i, bytes[i])
		case 0x4a:
			fmt.Printf("%04x %#02x MOV C,D\n", i, bytes[i])
		case 0x4b:
			fmt.Printf("%04x %#02x MOV C,E\n", i, bytes[i])
		case 0x4c:
			fmt.Printf("%04x %#02x MOV C,H\n", i, bytes[i])
		case 0x4d:
			fmt.Printf("%04x %#02x MOV C,L\n", i, bytes[i])
		case 0x4e:
			fmt.Printf("%04x %#02x MOV C,M\n", i, bytes[i])
		case 0x4f:
			fmt.Printf("%04x %#02x MOV C,A\n", i, bytes[i])
		case 0x50:
			fmt.Printf("%04x %#02x MOV D,B\n", i, bytes[i])
		case 0x51:
			fmt.Printf("%04x %#02x MOV D,C\n", i, bytes[i])
		case 0x52:
			fmt.Printf("%04x %#02x MOV D,D\n", i, bytes[i])
		case 0x53:
			fmt.Printf("%04x %#02x MOV D,E\n", i, bytes[i])
		case 0x54:
			fmt.Printf("%04x %#02x MOV D,H\n", i, bytes[i])
		case 0x55:
			fmt.Printf("%04x %#02x MOV D,L\n", i, bytes[i])
		case 0x56:
			fmt.Printf("%04x %#02x MOV D,M\n", i, bytes[i])
		case 0x57:
			fmt.Printf("%04x %#02x MOV D,A\n", i, bytes[i])
		case 0x58:
			fmt.Printf("%04x %#02x MOV E,B\n", i, bytes[i])
		case 0x59:
			fmt.Printf("%04x %#02x MOV E,C\n", i, bytes[i])
		case 0x5a:
			fmt.Printf("%04x %#02x MOV E,D\n", i, bytes[i])
		case 0x5b:
			fmt.Printf("%04x %#02x MOV E,E\n", i, bytes[i])
		case 0x5c:
			fmt.Printf("%04x %#02x MOV E,H\n", i, bytes[i])
		case 0x5d:
			fmt.Printf("%04x %#02x MOV E,L\n", i, bytes[i])
		case 0x5e:
			fmt.Printf("%04x %#02x MOV E,M\n", i, bytes[i])
		case 0x5f:
			fmt.Printf("%04x %#02x MOV E,A\n", i, bytes[i])
		case 0x60:
			fmt.Printf("%04x %#02x MOV H,B\n", i, bytes[i])
		case 0x61:
			fmt.Printf("%04x %#02x MOV H,C\n", i, bytes[i])
		case 0x62:
			fmt.Printf("%04x %#02x MOV H,D\n", i, bytes[i])
		case 0x63:
			fmt.Printf("%04x %#02x MOV H,E\n", i, bytes[i])
		case 0x64:
			fmt.Printf("%04x %#02x MOV H,H\n", i, bytes[i])
		case 0x65:
			fmt.Printf("%04x %#02x MOV H,L\n", i, bytes[i])
		case 0x66:
			fmt.Printf("%04x %#02x MOV H,M\n", i, bytes[i])
		case 0x67:
			fmt.Printf("%04x %#02x MOV H,A\n", i, bytes[i])
		case 0x68:
			fmt.Printf("%04x %#02x MOV L,B\n", i, bytes[i])
		case 0x69:
			fmt.Printf("%04x %#02x MOV L,C\n", i, bytes[i])
		case 0x6a:
			fmt.Printf("%04x %#02x MOV L,D\n", i, bytes[i])
		case 0x6b:
			fmt.Printf("%04x %#02x MOV L,E\n", i, bytes[i])
		case 0x6c:
			fmt.Printf("%04x %#02x MOV L,H\n", i, bytes[i])
		case 0x6d:
			fmt.Printf("%04x %#02x MOV L,L\n", i, bytes[i])
		case 0x6e:
			fmt.Printf("%04x %#02x MOV L,M\n", i, bytes[i])
		case 0x6f:
			fmt.Printf("%04x %#02x MOV L,A\n", i, bytes[i])
		case 0x70:
			fmt.Printf("%04x %#02x MOV M,B\n", i, bytes[i])
		case 0x71:
			fmt.Printf("%04x %#02x MOV M,C\n", i, bytes[i])
		case 0x72:
			fmt.Printf("%04x %#02x MOV M,D\n", i, bytes[i])
		case 0x73:
			fmt.Printf("%04x %#02x MOV M,E\n", i, bytes[i])
		case 0x74:
			fmt.Printf("%04x %#02x MOV M,H\n", i, bytes[i])
		case 0x75:
			fmt.Printf("%04x %#02x MOV M,L\n", i, bytes[i])
		case 0x76:
			fmt.Printf("%04x %#02x HLT\n", i, bytes[i])
		case 0x77:
			fmt.Printf("%04x %#02x MOV M,A\n", i, bytes[i])
		case 0x78:
			fmt.Printf("%04x %#02x MOV A,B\n", i, bytes[i])
		case 0x79:
			fmt.Printf("%04x %#02x MOV A,C\n", i, bytes[i])
		case 0x7a:
			fmt.Printf("%04x %#02x MOV A,D\n", i, bytes[i])
		case 0x7b:
			fmt.Printf("%04x %#02x MOV A,E\n", i, bytes[i])
		case 0x7c:
			fmt.Printf("%04x %#02x MOV A,H\n", i, bytes[i])
		case 0x7d:
			fmt.Printf("%04x %#02x MOV A,L\n", i, bytes[i])
		case 0x7e:
			fmt.Printf("%04x %#02x MOV A,M\n", i, bytes[i])
		case 0x7f:
			fmt.Printf("%04x %#02x MOV A,A\n", i, bytes[i])
		case 0x80:
			fmt.Printf("%04x %#02x ADD B\n", i, bytes[i])
		case 0x81:
			fmt.Printf("%04x %#02x ADD C\n", i, bytes[i])
		case 0x82:
			fmt.Printf("%04x %#02x ADD D\n", i, bytes[i])
		case 0x83:
			fmt.Printf("%04x %#02x ADD E\n", i, bytes[i])
		case 0x84:
			fmt.Printf("%04x %#02x ADD H\n", i, bytes[i])
		case 0x85:
			fmt.Printf("%04x %#02x ADD L\n", i, bytes[i])
		case 0x86:
			fmt.Printf("%04x %#02x ADD M\n", i, bytes[i])
		case 0x87:
			fmt.Printf("%04x %#02x ADD A\n", i, bytes[i])
		case 0x88:
			fmt.Printf("%04x %#02x ADC B\n", i, bytes[i])
		case 0x89:
			fmt.Printf("%04x %#02x ADC C\n", i, bytes[i])
		case 0x8a:
			fmt.Printf("%04x %#02x ADC D\n", i, bytes[i])
		case 0x8b:
			fmt.Printf("%04x %#02x ADC E\n", i, bytes[i])
		case 0x8c:
			fmt.Printf("%04x %#02x ADC H\n", i, bytes[i])
		case 0x8d:
			fmt.Printf("%04x %#02x ADC L\n", i, bytes[i])
		case 0x8e:
			fmt.Printf("%04x %#02x ADC M\n", i, bytes[i])
		case 0x8f:
			fmt.Printf("%04x %#02x ADC A\n", i, bytes[i])
		case 0x90:
			fmt.Printf("%04x %#02x SUB B\n", i, bytes[i])
		case 0x91:
			fmt.Printf("%04x %#02x SUB C\n", i, bytes[i])
		case 0x92:
			fmt.Printf("%04x %#02x SUB D\n", i, bytes[i])
		case 0x93:
			fmt.Printf("%04x %#02x SUB E\n", i, bytes[i])
		case 0x94:
			fmt.Printf("%04x %#02x SUB H\n", i, bytes[i])
		case 0x95:
			fmt.Printf("%04x %#02x SUB L\n", i, bytes[i])
		case 0x96:
			fmt.Printf("%04x %#02x SUB M\n", i, bytes[i])
		case 0x97:
			fmt.Printf("%04x %#02x SUB A\n", i, bytes[i])
		case 0x98:
			fmt.Printf("%04x %#02x SBB B\n", i, bytes[i])
		case 0x99:
			fmt.Printf("%04x %#02x SBB C\n", i, bytes[i])
		case 0x9a:
			fmt.Printf("%04x %#02x SBB D\n", i, bytes[i])
		case 0x9b:
			fmt.Printf("%04x %#02x SBB E\n", i, bytes[i])
		case 0x9c:
			fmt.Printf("%04x %#02x SBB H\n", i, bytes[i])
		case 0x9d:
			fmt.Printf("%04x %#02x SBB L\n", i, bytes[i])
		case 0x9e:
			fmt.Printf("%04x %#02x SBB M\n", i, bytes[i])
		case 0x9f:
			fmt.Printf("%04x %#02x SBB A\n", i, bytes[i])
		case 0xa0:
			fmt.Printf("%04x %#02x ANA B\n", i, bytes[i])
		case 0xa1:
			fmt.Printf("%04x %#02x ANA C\n", i, bytes[i])
		case 0xa2:
			fmt.Printf("%04x %#02x ANA D\n", i, bytes[i])
		case 0xa3:
			fmt.Printf("%04x %#02x ANA E\n", i, bytes[i])
		case 0xa4:
			fmt.Printf("%04x %#02x ANA H\n", i, bytes[i])
		case 0xa5:
			fmt.Printf("%04x %#02x ANA L\n", i, bytes[i])
		case 0xa6:
			fmt.Printf("%04x %#02x ANA M\n", i, bytes[i])
		case 0xa7:
			fmt.Printf("%04x %#02x ANA A\n", i, bytes[i])
		case 0xa8:
			fmt.Printf("%04x %#02x XRA B\n", i, bytes[i])
		case 0xa9:
			fmt.Printf("%04x %#02x XRA C\n", i, bytes[i])
		case 0xaa:
			fmt.Printf("%04x %#02x XRA D\n", i, bytes[i])
		case 0xab:
			fmt.Printf("%04x %#02x XRA E\n", i, bytes[i])
		case 0xac:
			fmt.Printf("%04x %#02x XRA H\n", i, bytes[i])
		case 0xad:
			fmt.Printf("%04x %#02x XRA L\n", i, bytes[i])
		case 0xae:
			fmt.Printf("%04x %#02x XRA M\n", i, bytes[i])
		case 0xaf:
			fmt.Printf("%04x %#02x XRA A\n", i, bytes[i])
		case 0xb0:
			fmt.Printf("%04x %#02x ORA B\n", i, bytes[i])
		case 0xb1:
			fmt.Printf("%04x %#02x ORA C\n", i, bytes[i])
		case 0xb2:
			fmt.Printf("%04x %#02x ORA D\n", i, bytes[i])
		case 0xb3:
			fmt.Printf("%04x %#02x ORA E\n", i, bytes[i])
		case 0xb4:
			fmt.Printf("%04x %#02x ORA H\n", i, bytes[i])
		case 0xb5:
			fmt.Printf("%04x %#02x ORA L\n", i, bytes[i])
		case 0xb6:
			fmt.Printf("%04x %#02x ORA M\n", i, bytes[i])
		case 0xb7:
			fmt.Printf("%04x %#02x ORA A\n", i, bytes[i])
		case 0xb8:
			fmt.Printf("%04x %#02x CMP B\n", i, bytes[i])
		case 0xb9:
			fmt.Printf("%04x %#02x CMP C\n", i, bytes[i])
		case 0xba:
			fmt.Printf("%04x %#02x CMP D\n", i, bytes[i])
		case 0xbb:
			fmt.Printf("%04x %#02x CMP E\n", i, bytes[i])
		case 0xbc:
			fmt.Printf("%04x %#02x CMP H\n", i, bytes[i])
		case 0xbd:
			fmt.Printf("%04x %#02x CMP L\n", i, bytes[i])
		case 0xbe:
			fmt.Printf("%04x %#02x CMP M\n", i, bytes[i])
		case 0xbf:
			fmt.Printf("%04x %#02x CMP A\n", i, bytes[i])
		case 0xc0:
			fmt.Printf("%04x %#02x RNZ\n", i, bytes[i])
		case 0xc1:
			fmt.Printf("%04x %#02x POP B\n", i, bytes[i])
		case 0xc2:
			fmt.Printf("%04x %#02x JNZ $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xc3:
			fmt.Printf("%04x %#02x JMP $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xc4:
			fmt.Printf("%04x %#02x CNZ $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xc5:
			fmt.Printf("%04x %#02x PUSH B\n", i, bytes[i])
		case 0xc6:
			fmt.Printf("%04x %#02x ADI $%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xc7:
			fmt.Printf("%04x %#02x RST 1\n", i, bytes[i])
		case 0xc8:
			fmt.Printf("%04x %#02x RZ\n", i, bytes[i])
		case 0xc9:
			fmt.Printf("%04x %#02x RET\n", i, bytes[i])
		case 0xca:
			fmt.Printf("%04x %#02x JZ $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xcc:
			fmt.Printf("%04x %#02x CZ $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xcd:
			fmt.Printf("%04x %#02x CALL $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xce:
			fmt.Printf("%04x %#02x ACI $%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xcf:
			fmt.Printf("%04x %#02x RST 5\n", i, bytes[i])
		case 0xd0:
			fmt.Printf("%04x %#02x RNC\n", i, bytes[i])
		case 0xd1:
			fmt.Printf("%04x %#02x POP D\n", i, bytes[i])
		case 0xd2:
			fmt.Printf("%04x %#02x JNC $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xd3:
			fmt.Printf("%04x %#02x OUT $%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xd4:
			fmt.Printf("%04x %#02x CNC $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xd5:
			fmt.Printf("%04x %#02x PUSH D\n", i, bytes[i])
		case 0xd6:
			fmt.Printf("%04x %#02x SUI $%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xd7:
			fmt.Printf("%04x %#02x RST 2\n", i, bytes[i])
		case 0xd8:
			fmt.Printf("%04x %#02x RC\n", i, bytes[i])
		case 0xda:
			fmt.Printf("%04x %#02x JC $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xdb:
			fmt.Printf("%04x %#02x IN $%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xdc:
			fmt.Printf("%04x %#02x CC $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xde:
			fmt.Printf("%04x %#02x SBI $%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xdf:
			fmt.Printf("%04x %#02x RST 6\n", i, bytes[i])
		case 0xe0:
			fmt.Printf("%04x %#02x RPE\n", i, bytes[i])
		case 0xe1:
			fmt.Printf("%04x %#02x POP H\n", i, bytes[i])
		case 0xe2:
			fmt.Printf("%04x %#02x JPO $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xe3:
			fmt.Printf("%04x %#02x XTHL\n", i, bytes[i])
		case 0xe4:
			fmt.Printf("%04x %#02x CPO $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xe5:
			fmt.Printf("%04x %#02x PUSH H\n", i, bytes[i])
		case 0xe6:
			fmt.Printf("%04x %#02x ANI,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xe7:
			fmt.Printf("%04x %#02x RST 3\n", i, bytes[i])
		case 0xe8:
			fmt.Printf("%04x %#02x RPO\n", i, bytes[i])
		case 0xe9:
			fmt.Printf("%04x %#02x PCHL H\n", i, bytes[i])
		case 0xea:
			fmt.Printf("%04x %#02x JPE $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xec:
			fmt.Printf("%04x %#02x CPE $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xeb:
			fmt.Printf("%04x %#02x XCHG\n", i, bytes[i])
		case 0xee:
			fmt.Printf("%04x %#02x XRI,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xef:
			fmt.Printf("%04x %#02x RST 7\n", i, bytes[i])
		case 0xf0:
			fmt.Printf("%04x %#02x RP\n", i, bytes[i])
		case 0xf1:
			fmt.Printf("%04x %#02x POP PSW\n", i, bytes[i])
		case 0xf2:
			fmt.Printf("%04x %#02x JP $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xf3:
			fmt.Printf("%04x %#02x DI\n", i, bytes[i])
		case 0xf4:
			fmt.Printf("%04x %#02x CP $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xf5:
			fmt.Printf("%04x %#02x PUSH PSW\n", i, bytes[i])
		case 0xf6:
			fmt.Printf("%04x %#02x ORI,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xf7:
			fmt.Printf("%04x %#02x RST 4\n", i, bytes[i])
		case 0xf8:
			fmt.Printf("%04x %#02x RM\n", i, bytes[i])
		case 0xf9:
			fmt.Printf("%04x %#02x SPHL\n", i, bytes[i])
		case 0xfa:
			fmt.Printf("%04x %#02x JM $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xfb:
			fmt.Printf("%04x %#02x EI\n", i, bytes[i])
		case 0xfc:
			fmt.Printf("%04x %#02x CM $%02x%02x\n", i, bytes[i], bytes[i+2], bytes[i+1])
			i = i + 2
		case 0xfe:
			fmt.Printf("%04x %#02x CPI,%02x\n", i, bytes[i], bytes[i+1])
			i = i + 1
		case 0xff:
			fmt.Printf("%04x %#02x RST 8\n", i, bytes[i])
		default:
			//fmt.Printf("%04x unknown opcode %#02x\n", i, bytes[i])
		}
	}
}
