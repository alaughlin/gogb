package main

var (
	_a, _b, _c, _d, _e, _f, _h, _l uint8  // cpu registers
	pc                             uint16 // program counter
	sp                             uint16 // stack pointer
	ram                            []uint8
	vram                           []uint8
)

type GameBoy struct{}

func (gb *GameBoy) reset() {
	ram = make([]uint8, 8192)
	pc = 0x100
	sp = 0xFFFE
}
