package main

import "os"

var (
	_a, _b, _c, _d, _e, _f, _h, _l, _m, _t uint8   // cpu registers
	m, t                                   uint8   // clock
	pc                                     uint16  // program counter
	sp                                     uint16  // stack pointer
	rom                                    []uint8 // 32k rom (0x0000-0x7FFF)
	vram                                   []uint8 // 8k vram (0x8000-0x9FFF)
	eram                                   []uint8 // 8k eram (0xA000-0xBFFF)
	wram                                   []uint8 // 8k wram (0xC000-0xDFFF)
	zram                                   []uint8 // 128b zram (0xFF80-0xFFFE)
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

func (gb *GameBoy) execute() {
	opcode := mmu.read(pc)
	executeOpcode(opcode)
}

func executeOpcode(opcode uint8) {
	switch opcode {
	case 0x00:
		NOP()
	case 0x01:
		LDBCd16(uint16(mmu.read(pc+1)<<8) + uint16(mmu.read(pc+2)))
	case 0x04:
		INCB()
	case 0x08:
		LDa16SP(uint16(mmu.read(pc+1)<<8) + uint16(mmu.read(pc+2)))
	case 0x11:
		LDDEd16(uint16(mmu.read(pc+1)<<8) + uint16(mmu.read(pc+2)))
	case 0x18:
		JRr8(mmu.read(pc + 1))
	case 0x20:
		JRNZr8(mmu.read(pc + 1))
	case 0x21:
		LDHLd16(uint16(mmu.read(pc+1)<<8) + uint16(mmu.read(pc+2)))
	case 0x22:
		LDIaHLA()
	case 0x26:
		LDHd8(mmu.read(pc + 1))
	case 0x28:
		JRZr8(mmu.read(pc + 1))
	case 0x2A:
		LDIAaHL()
	case 0x2E:
		LDLd8(mmu.read(pc + 1))
	case 0x30:
		JRNCr8(mmu.read(pc + 1))
	case 0x31:
		LDSPd16(uint16(mmu.read(pc+1)<<8) + uint16(mmu.read(pc+2)))
	case 0x32:
		LDDaHLA()
	case 0x36:
		LDaHLd8(mmu.read(pc + 1))
	case 0x38:
		JRCr8(mmu.read(pc + 1))
	case 0x3A:
		LDDAaHL()
	case 0x3E:
		LDAd8(mmu.read(pc + 1))
	case 0x40:
		LDBB()
	case 0x41:
		LDBC()
	case 0x42:
		LDBD()
	case 0x43:
		LDBE()
	case 0x44:
		LDBH()
	case 0x45:
		LDBL()
	case 0x46:
		LDBaHL()
	case 0x47:
		LDBA()
	case 0x48:
		LDCB()
	case 0x49:
		LDCC()
	case 0x4A:
		LDCD()
	case 0x4B:
		LDCE()
	case 0x4C:
		LDCH()
	case 0x4D:
		LDCL()
	case 0x4E:
		LDCaHL()
	case 0x4F:
		LDCA()
	case 0x50:
		LDDB()
	case 0x51:
		LDDC()
	case 0x52:
		LDDD()
	case 0x53:
		LDDE()
	case 0x54:
		LDDH()
	case 0x55:
		LDDL()
	case 0x56:
		LDDaHL()
	case 0x57:
		LDDA()
	case 0x58:
		LDEB()
	case 0x59:
		LDEB()
	case 0x5A:
		LDED()
	case 0x5B:
		LDEE()
	case 0x5C:
		LDEH()
	case 0x5D:
		LDEL()
	case 0x5E:
		LDEaHL()
	case 0x5F:
		LDEA()
	case 0x60:
		LDHB()
	case 0x61:
		LDHC()
	case 0x62:
		LDHD()
	case 0x63:
		LDHE()
	case 0x64:
		LDHH()
	case 0x65:
		LDHL()
	case 0x66:
		LDHaHL()
	case 0x67:
		LDHA()
	case 0x68:
		LDLB()
	case 0x69:
		LDLC()
	case 0x6A:
		LDLD()
	case 0x6B:
		LDLE()
	case 0x6C:
		LDLH()
	case 0x6D:
		LDLL()
	case 0x6E:
		LDLaHL()
	case 0x6F:
		LDLA()
	case 0x70:
		LDaHLB()
	case 0x71:
		LDaHLC()
	case 0x72:
		LDaHLD()
	case 0x73:
		LDaHLE()
	case 0x74:
		LDaHLH()
	case 0x75:
		LDaHLL()
	case 0x77:
		LDaHLA()
	case 0x78:
		LDAB()
	case 0x79:
		LDAC()
	case 0x7A:
		LDAD()
	case 0x7B:
		LDAE()
	case 0x7C:
		LDAH()
	case 0x7D:
		LDAL()
	case 0x7E:
		LDAaHL()
	case 0x7F:
		LDAA()
	case 0xC0:
		RETNZ()
	case 0xC1:
		POPBC()
	case 0xC2:
		JPBZa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xC3:
		JPa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xC4:
		CALLNZa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xC5:
		PUSHBC()
	case 0xC7:
		RST00H()
	case 0xC8:
		RETZ()
	case 0xC9:
		RET()
	case 0xCA:
		JPZa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xCC:
		CALLZa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xCD:
		CALLa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xCF:
		RST08()
	case 0xD0:
		RETNC()
	case 0xD1:
		POPDE()
	case 0xD2:
		JPNCa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xD4:
		CALLNCa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xD5:
		PUSHDE()
	case 0xD7:
		RST10H()
	case 0xD8:
		RETC()
	case 0xD9:
		RETI()
	case 0xDA:
		JPCa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xDC:
		CALLCa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xDF:
		RST18H()
	case 0xE0:
		LDHa8A(mmu.read(pc + 1))
	case 0xE1:
		POPHL()
	case 0xE2:
		LDaCA()
	case 0xE5:
		PUSHHL()
	case 0xE7:
		RST20H()
	case 0xE9:
		JPHL()
	case 0xEA:
		LDa16A(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xEF:
		RST28H()
	case 0xF0:
		LDHAa8(mmu.read(pc + 1))
	case 0xF1:
		POPAF()
	case 0xF2:
		LDAmC()
	case 0xF5:
		PUSHAF()
	case 0xF7:
		RST30H()
	case 0xF8:
		LDHLSPpr8(mmu.read(pc + 1))
	case 0xF9:
		LDSPHL()
	case 0xFA:
		LDAa16(uint16(mmu.read(pc+2)<<8) + uint16(mmu.read(pc+1)))
	case 0xFF:
		RST38H()
	}
}
