package processor

func makeUint16(lo uint8, ho uint8) uint16 {
	addr := uint(lo)
	addr = (addr << 8) | uint(ho)
	return uint16(addr)
}
