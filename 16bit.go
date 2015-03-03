package main

/* 16bit Load/Store/Move */
/* opcode, asm command */

// 01 LD BC,d16
func LDBCd16(d uint16) {
	_b = uint8(d & 0x00FF)
	_c = uint8(d >> 8)
	pc += 3
	_m, _t = 3, 12
}

// 08 LD (a16), SP
func LDa16SP(a uint16) {
	sp = a
	pc += 3
	_m, _t = 5, 20
}

// 11 LD DE,d16
func LDDEd16(d uint16) {
	_d = uint8(d & 0x00FF)
	_e = uint8(d >> 8)
	pc += 3
	_m, _t = 3, 12
}

// 21 LD HL,d16
func LDHLd16(d uint16) {
	_h = uint8(d & 0x00FF)
	_l = uint8(d >> 8)
	pc += 3
	_m, _t = 3, 12
}

// 31 LD SP,d16
func LDSPd16(d uint16) {
	sp = uint16((d&0x00FF)<<8) + uint16(d>>8)
	pc += 3
	_m, _t = 3, 12
}

// C1 POP BC
func POPBC() {
	vals := pop(2)
	_b, _c = vals[0], vals[1]
	pc++
	_m, _t = 3, 12
}

// C5 PUSH BC
func PUSHBC() {
	push(_b, _c)
	pc++
	_m, _t = 4, 16
}

// D1 POP DE
func POPDE() {
	vals := pop(2)
	_d, _e = vals[0], vals[1]
	pc++
	_m, _t = 3, 12
}

// D5 PUSH DE
func PUSHDE() {
	push(_d, _e)
	pc++
	_m, _t = 4, 16
}

// E1 POP HL
func POPHL() {
	vals := pop(2)
	_h, _l = vals[0], vals[1]
	pc++
	_m, _t = 3, 12
}

// E5 PUSH HL {
func PUSHHL() {
	push(_h, _l)
	pc++
	_m, _t = 4, 16
}

// F1 POP AF
func POPAF() {
	vals := pop(2)
	_a, _f = vals[0], vals[1]
	pc++
	_m, _t = 3, 12
}

// F5 PUSH AF
func PUSHAF() {
	push(_a, _f)
	pc++
	_m, _t = 4, 16
}

// F8 LD HL,SP+r8
// TODO
func LDHLSPpr8(r uint8) {
	pc += 2
	_m, _t = 3, 12
}

// F9 LD SP,HL
func LDSPHL() {
	hl := uint16(_h<<8) + uint16(_l)
	sp = hl
	pc++
	_m, _t = 2, 8
}

/* 16bit Arithmetic/Logical */
/* opcode, asm command */

// 03 INC BC
func INCBC() {
	_b++
	_c++
	pc++
	_m, _t = 2, 8
}

// 13 INC DE
func INCDE() {
	_d++
	_e++
	pc++
	_m, _t = 2, 8
}

// 23 INC HL
func INCHL() {
	_h++
	_l++
	pc++
	_m, _t = 2, 8
}

// 33 INC SP
func INCSP() {
	sp++
	_m, _t = 2, 8
}
