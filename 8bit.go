package main

/* 8bit LD */
/* opcode, asm command */

// 02 LD (BC),A
func LDaBCA() {
	bc := uint16(_b<<8) + uint16(_c)
	ram[bc] = _a
	pc++
}

// 06 LD B,d8
func LDBd8(d uint8) {
	_b = d
	pc += 2
}

// 0A LD A,(BC)
func LDAaBC() {
	bc := uint16(_b<<8) + uint16(_c)
	_a = ram[bc]
	pc++
}

// 0E LD C,d8
func LDCd8(d uint8) {
	_c = d
	pc += 2
}

// 12 LD (DE),A
func LDaDEA() {
	de := uint16(_d<<8) + uint16(_e)
	ram[de] = _a
	pc++
}

// 16 LD D,d8
func LDDd8(d uint8) {
	_d = d
	pc += 2
}

// 1A LD A,(DE)
func LDAaDE() {
	de := uint16(_d<<8) + uint16(_e)
	_a = ram[de]
	pc++
}

// 1E LD E,d8
func LDEd8(d uint8) {
	_e = d
	pc += 2
}

// 22 LDI (HL),A
func LDIaHLA() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _a
	ram[hl]++
	pc++
}

// 26 LD H,d8
func LDHd8(d uint8) {
	_h = d
	pc += 2
}

// 2A LDI A,(HL)
func LDIAaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_a = ram[hl]
	ram[hl]++
	pc++
}

// 2E LD L,d8
func LDLd8(d uint8) {
	_l = d
	pc += 2
}

// 32 LDD (HL),A
func LDDaHLA() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _a
	ram[hl]--
	pc++
}

// 36 LD (HL),d8
func LDaHLd8(d uint8) {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = d
	pc += 2
}

// 3A LDD A,(HL)
func LDDAaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_a = ram[hl]
	ram[hl]--
	pc++
}

// 3E LD A,d8
func LDAd8(d uint8) {
	_a = d
	pc += 2
}

// 40 LD B,B
func LDBB() {
	_b = _b
	pc++
}

// 41 LD B,C
func LDBC() {
	_b = _c
	pc++
}

// 42 LD B,D
func LDBD() {
	_b = _d
	pc++
}

// 43 LD B,E
func LDBE() {
	_b = _e
	pc++
}

// 44 LD B,H
func LDBH() {
	_b = _h
	pc++
}

// 45 LD B,L
func LDBL() {
	_b = _l
	pc++
}

// 46 LD B,(HL)
func LDBaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_b = ram[hl]
	pc++
}

// 47 LD B,A
func LDBA() {
	_b = _a
	pc++
}

// 48 LD C,B
func LDCB() {
	_c = _b
	pc++
}

// 49 LD C,C
func LDCC() {
	_c = _c
	pc++
}

// 4A LD C,D
func LDCD() {
	_c = _d
	pc++
}

// 4B LD C,E
func LDCE() {
	_c = _e
	pc++
}

// 4C LD C,H
func LDCH() {
	_c = _h
	pc++
}

// 4D LD C,L
func LDCL() {
	_c = _l
	pc++
}

// 4E LD C,(HL)
func LDCaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_c = ram[hl]
	pc++
}

// 4F LD C,A
func LDCA() {
	_c = _a
	pc++
}

// 50 LD D,B
func LDDB() {
	_d = _b
	pc++
}

// 51 LD D,C
func LDDC() {
	_d = _c
	pc++
}

// 52 LD D,D
func LDDD() {
	_d = _d
	pc++
}

// 53 LD D,E
func LDDE() {
	_d = _e
	pc++
}

// 54 LD D,H
func LDDH() {
	_d = _h
	pc++
}

// 55 LD D,L
func LDDL() {
	_d = _l
	pc++
}

// 56 LD D,(HL)
func LDDaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_d = ram[hl]
	pc++
}

// 57 LD D,A
func LDDA() {
	_d = _a
	pc++
}

// 58 LD E,B
func LDEB() {
	_e = _b
	pc++
}

// 59 LD E,C
func LDEC() {
	_e = _c
	pc++
}

// 5A LD E,E
func LDED() {
	_e = _d
	pc++
}

// 5B LD E,E
func LDEE() {
	_e = _e
	pc++
}

// 5C LD E,H
func LDEH() {
	_e = _h
	pc++
}

// 5D LD E,L
func LDEL() {
	_e = _l
	pc++
}

// 5E LD E,(HL)
func LDEaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_e = ram[hl]
	pc++
}

// 5F LD E,A
func LDEA() {
	_e = _a
	pc++
}

// 60 LD H,B
func LDHB() {
	_h = _b
	pc++
}

// 61 LD H,C
func LDHC() {
	_h = _c
	pc++
}

// 62 LD H,D
func LDHD() {
	_h = _d
	pc++
}

// 63 LD H,E
func LDHE() {
	_h = _e
	pc++
}

// 64 LD H,H
func LDHH() {
	_h = _h
	pc++
}

// 65 LD H,L
func LDHL() {
	_h = _l
	pc++
}

// 66 LD H,(HL)
func LDHaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_h = ram[hl]
	pc++
}

// 67 LD H,A
func LDHA() {
	_h = _a
	pc++
}

// 68 LD L,B
func LDLB() {
	_l = _b
	pc++
}

// 69 LD L,C
func LDLC() {
	_l = _c
	pc++
}

// 6A LD L,D
func LDLD() {
	_l = _d
	pc++
}

// 6B LD L,E
func LDLE() {
	_l = _e
	pc++
}

// 6C LD L,H
func LDLH() {
	_l = _h
	pc++
}

// 6D LD L,L
func LDLL() {
	_l = _l
	pc++
}

// 6E LD L,(HL)
func LDLaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_l = ram[hl]
	pc++
}

// 6F LD L,A
func LDLA() {
	_l = _a
	pc++
}

// 70 LD (HL),B
func LDaHLB() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _b
	pc++
}

// 71 LD (HL),C
func LDaHLC() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _c
}

// 72 LD (HL),D
func LDaHLD() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _d
	pc++
}

// 73 LD (HL),E
func LDaHLE() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _e
	pc++
}

// 74 LD (HL),H
func LDaHLH() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _h
	pc++
}

// 75 LD (HL),L
func LDaHLL() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _l
	pc++
}

// 77 LD (HL),A
func LDaHLA() {
	hl := uint16(_h<<8) + uint16(_l)
	ram[hl] = _a
	pc++
}

// 78 LD A,B
func LDAB() {
	_a = _b
	pc++
}

// 79 LD A,C
func LDAC() {
	_a = _c
	pc++
}

// 7A LD A,D
func LDAD() {
	_a = _d
	pc++
}

// 7B LD A,E
func LDAE() {
	_a = _e
	pc++
}

// 7C LD A,H
func LDAH() {
	_a = _h
	pc++
}

// 7D LD A,L
func LDAL() {
	_a = _l
	pc++
}

// 7E LD A,(HL)
func LDAaHL() {
	hl := uint16(_h<<8) + uint16(_l)
	_a = ram[hl]
	pc++
}

// 7F LD A,A
func LDAA() {
	_a = _a
	pc++
}

// E0 LDH (a8),A
func LDHa8A(a uint8) {
	ram[a] = _a
	pc += 2
}

// E2 LD (C), A
func LDaCA() {
	ram[_c] = _a
	pc++
}

// EA LD (a16),A
func LDa16A(a uint16) {
	ram[a] = _a
	pc += 3
}

// F0 LDH A,(a8)
func LDHAa8(a uint8) {
	_a = ram[a]
	pc += 2
}

// F2 LD A,(C)
func LDAmC() {
	_a = ram[_c]
	pc++
}

// FA LD A,(a16)
func LDAa16(a uint16) {
	_a = ram[a]
	pc += 3
}