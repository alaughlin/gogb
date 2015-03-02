package main

// pops n bytes from zram
func pop(n uint16) []uint8 {
	res := make([]uint8, n)
	for i := uint16(0); i < n; i++ {
		res = append(res, zram[(sp+i)-0xFF80])
		sp++
	}
	return res
}

// pushes n bytes into zram
func push(vals ...uint8) {
	for val := range vals {
		zram[sp-0xFF80] = uint8(val)
		sp--
	}
}

func getMemorySlice(address uint16) ([]uint8, uint16) {
	if address < 0x8000 {
		return rom, 0x0000
	} else if address < 0xA000 {
		return vram, 0x8000
	} else if address < 0xC000 {
		return eram, 0xA000
	} else if address < 0xE000 {
		return wram, 0xC000
	} else if address >= 0xFF80 && address <= 0xFFFE {
		return zram, 0xFF80
	}

	return zram, 0xFF80
}
