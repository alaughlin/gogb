package main

// pops n bytes from ram
func pop(n uint16) []uint8 {
	res := make([]uint8, n)
	for i := uint16(0); i < n; i++ {
		res = append(res, ram[sp+i])
		sp++
	}
	return res
}

// pushes n bytes into ram
func push(vals ...uint8) {
	for val := range vals {
		ram[sp] = uint8(val)
		sp--
	}
}
