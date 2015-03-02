package main

import "os"

var (
	_a, _b, _c, _d, _e, _f, _h, _l uint8   // cpu registers
	pc                             uint16  // program counter
	sp                             uint16  // stack pointer
	rom                            []uint8 // 32k rom (0x0000-0x7FFF)
	vram                           []uint8 // 8k vram (0x8000-0x9FFF)
	eram                           []uint8 // 8k eram (0xA000-0xBFFF)
	wram                           []uint8 // 8k wram (0xC000-0xDFFF)
	zram                           []uint8 // 128b zram (0xFF80-0xFFFE)
)

type GameBoy struct{}

func (gb *GameBoy) reset() {
	mmu = &MMU{}
	rom = make([]uint8, 32768)
	vram = make([]uint8, 8192)
	eram = make([]uint8, 8192)
	wram = make([]uint8, 8192)
	zram = make([]uint8, 128)
	pc = 0x100
	sp = 0xFFFE
}

func (gb *GameBoy) loadGame() {
	f, _ := os.Open("tetris.gb")
	b := make([]byte, 32768)
	f.Read(b)

	for i := 0; i < len(b); i++ {
		rom[i] = b[i]
	}
}
