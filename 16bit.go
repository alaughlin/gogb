package main

/* 16bit LD */
/* opcode, asm command */

// 01 LD BC,d16
func LDBCd16(d uint16) {
	_b = uint8(d & 0x00FF)
	_c = uint8(d >> 8)
	pc += 3
}

// 08 LD (a16), SP
func LDa16SP(a uint16) {
	sp = a
	pc += 3
}

// 11 LD DE,d16
func LDDEd16(d uint16) {
	_d = uint8(d & 0x00FF)
	_e = uint8(d >> 8)
	pc += 3
}

// 21 LD HL,d16
func LDHLd16(d uint16) {
	_h = uint8(d & 0x00FF)
	_l = uint8(d >> 8)
	pc += 3
}

// 31 LD SP,d16
func LDSPd16(d uint16) {
	sp = uint16((d&0x00FF)<<8) + uint16(d>>8)
	pc += 3
}

// C1 POP BC
func POPBC() {
	vals := pop(2)
	_b, _c = vals[0], vals[1]
	pc++
}

// C5 PUSH BC
func PUSHBC() {
	push(_b, _c)
	pc++
}

// D1 POP DE
func POPDE() {
	vals := pop(2)
	_d, _e = vals[0], vals[1]
	pc++
}

// D5 PUSH DE
func PUSHDE() {
	push(_d, _e)
	pc++
}

// E1 POP HL
func POPHL() {
	vals := pop(2)
	_h, _l = vals[0], vals[1]
	pc++
}

// E5 PUSH HL {
func PUSHHL() {
	push(_h, _l)
	pc++
}

// F1 POP AF
func POPAF() {
	vals := pop(2)
	_a, _f = vals[0], vals[1]
	pc++
}

// F5 PUSH AF
func PUSHAF() {
	push(_a, _f)
	pc++
}

// F8 LD HL,SP+r8
// TODO
func LDHLSPpr8(r uint8) {
	pc += 2
}

// F9 LD SP,HL
func LDSPHL() {
	hl := uint16(_h<<8) + uint16(_l)
	sp = hl
	pc++
}
