package main

/* 8bit Load/Store/Move */
/* opcode, asm command */

// 02 LD (BC),A
func LDaBCA() {
	bc := uint16(_b<<8) + uint16(_c)
	mmu.write(bc, _a)
	pc++
	_m, _t = 2, 8
}

// 06 LD B,d8
func LDBd8(d uint8) {
	_b = d
	pc += 2
	_m, _t = 2, 8
}

// 0A LD A,(BC)
func LDAaBC() {
	bc := uint16(_b<<8) + uint16(_c)
	_a = mmu.read(bc)
	pc++
	_m, _t = 2, 8
}

// 0E LD C,d8
func LDCd8(d uint8) {
	_c = d
	pc += 2
	_m, _t = 2, 8
}

// 12 LD (DE),A
func LDaDEA() {
	de := uint16(_d<<8) + uint16(_e)
	mmu.write(de, _a)
	pc++
	_m, _t = 2, 8
}

// 16 LD D,d8
func LDDd8(d uint8) {
	_d = d
	pc += 2
	_m, _t = 2, 8
}

// 1A LD A,(DE)
func LDAaDE() {
	de := uint16(_d<<8) + uint16(_e)
	_a = mmu.read(de)
	pc++
	_m, _t = 2, 8
}

// 1E LD E,d8
func LDEd8(d uint8) {
	_e = d
	pc += 2
	_m, _t = 2, 8
}

// 22 LDI (HL),A
func LDIaHLA() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _a)
	mmu.incr(hl)
	pc++
	_m, _t = 2, 8
}

// 26 LD H,d8
func LDHd8(d uint8) {
	_h = d
	pc += 2
	_m, _t = 2, 8
}

// 2A LDI A,(HL)
func LDIAaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_a = mmu.read(hl)
	mmu.incr(hl)
	pc++
	_m, _t = 2, 8
}

// 2E LD L,d8
func LDLd8(d uint8) {
	_l = d
	pc += 2
	_m, _t = 2, 8
}

// 32 LDD (HL),A
func LDDaHLA() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _a)
	mmu.decr(hl)
	pc++
	_m, _t = 2, 8
}

// 36 LD (HL),d8
func LDaHLd8(d uint8) {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, d)
	pc += 2
	_m, _t = 3, 12
}

// 3A LDD A,(HL)
func LDDAaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_a = mmu.read(hl)
	mmu.decr(hl)
	pc++
	_m, _t = 2, 8
}

// 3E LD A,d8
func LDAd8(d uint8) {
	_a = d
	pc += 2
	_m, _t = 2, 8
}

// 40 LD B,B
func LDBB() {
	_b = _b
	pc++
	_m, _t = 1, 4
}

// 41 LD B,C
func LDBC() {
	_b = _c
	pc++
	_m, _t = 1, 4
}

// 42 LD B,D
func LDBD() {
	_b = _d
	pc++
	_m, _t = 1, 4
}

// 43 LD B,E
func LDBE() {
	_b = _e
	pc++
	_m, _t = 1, 4
}

// 44 LD B,H
func LDBH() {
	_b = _h
	pc++
	_m, _t = 1, 4
}

// 45 LD B,L
func LDBL() {
	_b = _l
	pc++
	_m, _t = 1, 4
}

// 46 LD B,(HL)
func LDBaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_b = mmu.read(hl)
	pc++
	_m, _t = 2, 8
}

// 47 LD B,A
func LDBA() {
	_b = _a
	pc++
	_m, _t = 1, 4
}

// 48 LD C,B
func LDCB() {
	_c = _b
	pc++
	_m, _t = 1, 4
}

// 49 LD C,C
func LDCC() {
	_c = _c
	pc++
	_m, _t = 1, 4
}

// 4A LD C,D
func LDCD() {
	_c = _d
	pc++
	_m, _t = 1, 4
}

// 4B LD C,E
func LDCE() {
	_c = _e
	pc++
	_m, _t = 1, 4
}

// 4C LD C,H
func LDCH() {
	_c = _h
	pc++
	_m, _t = 1, 4
}

// 4D LD C,L
func LDCL() {
	_c = _l
	pc++
	_m, _t = 1, 4
}

// 4E LD C,(HL)
func LDCaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_c = mmu.read(hl)
	pc++
	_m, _t = 2, 8
}

// 4F LD C,A
func LDCA() {
	_c = _a
	pc++
	_m, _t = 1, 4
}

// 50 LD D,B
func LDDB() {
	_d = _b
	pc++
	_m, _t = 1, 4
}

// 51 LD D,C
func LDDC() {
	_d = _c
	pc++
	_m, _t = 1, 4
}

// 52 LD D,D
func LDDD() {
	_d = _d
	pc++
	_m, _t = 1, 4
}

// 53 LD D,E
func LDDE() {
	_d = _e
	pc++
	_m, _t = 1, 4
}

// 54 LD D,H
func LDDH() {
	_d = _h
	pc++
	_m, _t = 1, 4
}

// 55 LD D,L
func LDDL() {
	_d = _l
	pc++
	_m, _t = 1, 4
}

// 56 LD D,(HL)
func LDDaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_d = mmu.read(hl)
	pc++
	_m, _t = 2, 8
}

// 57 LD D,A
func LDDA() {
	_d = _a
	pc++
	_m, _t = 1, 4
}

// 58 LD E,B
func LDEB() {
	_e = _b
	pc++
	_m, _t = 1, 4
}

// 59 LD E,C
func LDEC() {
	_e = _c
	pc++
	_m, _t = 1, 4
}

// 5A LD E,E
func LDED() {
	_e = _d
	pc++
	_m, _t = 1, 4
}

// 5B LD E,E
func LDEE() {
	_e = _e
	pc++
	_m, _t = 1, 4
}

// 5C LD E,H
func LDEH() {
	_e = _h
	pc++
	_m, _t = 1, 4
}

// 5D LD E,L
func LDEL() {
	_e = _l
	pc++
	_m, _t = 1, 4
}

// 5E LD E,(HL)
func LDEaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_e = mmu.read(hl)
	pc++
	_m, _t = 2, 8
}

// 5F LD E,A
func LDEA() {
	_e = _a
	pc++
	_m, _t = 1, 4
}

// 60 LD H,B
func LDHB() {
	_h = _b
	pc++
	_m, _t = 1, 4
}

// 61 LD H,C
func LDHC() {
	_h = _c
	pc++
	_m, _t = 1, 4
}

// 62 LD H,D
func LDHD() {
	_h = _d
	pc++
	_m, _t = 1, 4
}

// 63 LD H,E
func LDHE() {
	_h = _e
	pc++
	_m, _t = 1, 4
}

// 64 LD H,H
func LDHH() {
	_h = _h
	pc++
	_m, _t = 1, 4
}

// 65 LD H,L
func LDHL() {
	_h = _l
	pc++
	_m, _t = 1, 4
}

// 66 LD H,(HL)
func LDHaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_h = mmu.read(hl)
	pc++
	_m, _t = 2, 8
}

// 67 LD H,A
func LDHA() {
	_h = _a
	pc++
	_m, _t = 1, 4
}

// 68 LD L,B
func LDLB() {
	_l = _b
	pc++
	_m, _t = 1, 4
}

// 69 LD L,C
func LDLC() {
	_l = _c
	pc++
	_m, _t = 1, 4
}

// 6A LD L,D
func LDLD() {
	_l = _d
	pc++
	_m, _t = 1, 4
}

// 6B LD L,E
func LDLE() {
	_l = _e
	pc++
	_m, _t = 1, 4
}

// 6C LD L,H
func LDLH() {
	_l = _h
	pc++
	_m, _t = 1, 4
}

// 6D LD L,L
func LDLL() {
	_l = _l
	pc++
	_m, _t = 1, 4
}

// 6E LD L,(HL)
func LDLaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_l = mmu.read(hl)
	pc++
	_m, _t = 2, 8
}

// 6F LD L,A
func LDLA() {
	_l = _a
	pc++
	_m, _t = 1, 4
}

// 70 LD (HL),B
func LDaHLB() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _b)
	pc++
	_m, _t = 2, 8
}

// 71 LD (HL),C
func LDaHLC() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _c)
	pc++
	_m, _t = 2, 8
}

// 72 LD (HL),D
func LDaHLD() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _d)
	pc++
	_m, _t = 2, 8
}

// 73 LD (HL),E
func LDaHLE() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _e)
	pc++
	_m, _t = 2, 8
}

// 74 LD (HL),H
func LDaHLH() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _h)
	pc++
	_m, _t = 2, 8
}

// 75 LD (HL),L
func LDaHLL() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _l)
	pc++
	_m, _t = 2, 8
}

// 77 LD (HL),A
func LDaHLA() {
	hl := uint16(_h<<8) + uint16(_l)
	mmu.write(hl, _a)
	pc++
	_m, _t = 2, 8
}

// 78 LD A,B
func LDAB() {
	_a = _b
	pc++
	_m, _t = 1, 4
}

// 79 LD A,C
func LDAC() {
	_a = _c
	pc++
	_m, _t = 1, 4
}

// 7A LD A,D
func LDAD() {
	_a = _d
	pc++
	_m, _t = 1, 4
}

// 7B LD A,E
func LDAE() {
	_a = _e
	pc++
	_m, _t = 1, 4
}

// 7C LD A,H
func LDAH() {
	_a = _h
	pc++
	_m, _t = 1, 4
}

// 7D LD A,L
func LDAL() {
	_a = _l
	pc++
	_m, _t = 1, 4
}

// 7E LD A,(HL)
func LDAaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_a = mmu.read(hl)
	pc++
	_m, _t = 2, 8
}

// 7F LD A,A
func LDAA() {
	_a = _a
	pc++
	_m, _t = 1, 4
}

// E0 LDH (a8),A
func LDHa8A(a uint8) {
	mmu.write(0xFF00+uint16(a), _a)
	pc += 2
	_m, _t = 3, 12
}

// E2 LD (C), A
func LDaCA() {
	mmu.write(0xFF00+uint16(_c), _a)
	pc++
	_m, _t = 2, 8
}

// EA LD (a16),A
func LDa16A(a uint16) {
	mmu.write(a, _a)
	pc += 3
	_m, _t = 4, 16
}

// F0 LDH A,(a8)
func LDHAa8(a uint8) {
	_a = mmu.read(0xFF00 + uint16(a))
	pc += 2
	_m, _t = 3, 12
}

// F2 LD A,(C)
func LDAmC() {
	_a = mmu.read(0xFF00 + uint16(_c))
	pc++
	_m, _t = 2, 8
}

// FA LD A,(a16)
func LDAa16(a uint16) {
	_a = mmu.read(a)
	pc += 3
	_m, _t = 4, 16
}

/* 8bit Arithmetic/Logical */
/* opcode, asm command */

// 04 INC B
// TODO
func INCB() {
	_b++
}
