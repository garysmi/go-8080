package processor

func makeUint16(lo uint8, ho uint8) uint16 {
	addr := uint(lo)
	addr = (addr << 8) | uint(ho)
	return uint16(addr)
}

//Parity returns true if the number of set bits is odd
func parity(x uint8) bool {
	x ^= (x >> 4)
	x ^= (x >> 2)
	x ^= (x >> 1)

	if x&0x01 == 0 {
		return true
	}

	return false
}
