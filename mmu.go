package main

// 0000-3FFF   16KB ROM Bank 00            (ROM)  (in cartridge, fixed at bank 00)
// 4000-7FFF   16KB ROM Bank 01..NN        (ROM)  (in cartridge, switchable bank number)
// 8000-9FFF   8KB Video RAM               (VRAM) (switchable bank 0-1 in CGB Mode)
// A000-BFFF   8KB External RAM            (in cartridge, switchable bank, if any)
// C000-CFFF   4KB Work RAM Bank 0         (WRAM)
// D000-DFFF   4KB Work RAM Bank 1         (WRAM) (switchable bank 1-7 in CGB Mode)
// E000-FDFF   Same as C000-DDFF           (ECHO) (typically not used)
// FE00-FE9F   Sprite Attribute Table      (OAM)
// FEA0-FEFF   Not Usable
// FF00-FF7F   I/O Ports
// FF80-FFFE   High RAM                    (ZRAM)
// FFFF        Interrupt Enable Register

type MMU struct{}

func (mmu *MMU) write(address uint16, data uint8) {
	memorySlice, offset := getMemorySlice(address)
	memorySlice[address-offset] = data
}

func (mmu *MMU) read(address uint16) uint8 {
	memorySlice, offset := getMemorySlice(address)
	return memorySlice[address-offset]
}

func (mmu *MMU) incr(address uint16) {
	memorySlice, offset := getMemorySlice(address)
	memorySlice[address-offset]++
}

func (mmu *MMU) decr(address uint16) {
	memorySlice, offset := getMemorySlice(address)
	memorySlice[address-offset]--
}
