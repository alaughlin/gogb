package main

/* Jumps/Calls */
/* opcode, asm command */

// 18 JR r8
func JRr8(r uint8) {
	pc += uint16(r)
	_m, _t = 2, 8
}

// 20 JR NZ,r8
func JRNZr8(r uint8) {
	if _f&(1<<7) == 0 {
		pc += uint16(r)
	} else {
		pc += 2
	}
	_m, _t = 2, 8
}

// 28 JR Z,r8
func JRZr8(r uint8) {
	if _f&(1<<7) > 0 {
		pc += uint16(r)
	} else {
		pc += 2
	}
	_m, _t = 2, 8
}

// 30 JR NC,r8
func JRNCr8(r uint8) {
	if _f&(1<<4) == 0 {
		pc += uint16(r)
	} else {
		pc += 2
	}
	_m, _t = 2, 8
}

// 38 JR C,r8
func JRCr8(r uint8) {
	if _f&(1<<4) > 0 {
		pc += uint16(r)
	} else {
		pc += 2
	}
	_m, _t = 2, 8
}

// C0 RET NZ
func RETNZ() {
	if _f&(1<<7) == 0 {
		bytes := pop(2)
		pc = uint16(bytes[0]<<8) + uint16(bytes[1])
	} else {
		pc++
	}
	_m, _t = 2, 8
}

// C8 RET Z
func RETZ() {
	if _f&(1<<7) > 0 {
		bytes := pop(2)
		pc = uint16(bytes[0]<<8) + uint16(bytes[1])
	} else {
		pc++
	}
	_m, _t = 2, 8
}

// D0 RET NC
func RETNC() {
	if _f&(1<<4) == 0 {
		bytes := pop(2)
		pc = uint16(bytes[0]<<8) + uint16(bytes[1])
	} else {
		pc++
	}
	_m, _t = 2, 8
}

// D8 RET C
func RETC() {
	if _f&(1<<4) > 0 {
		bytes := pop(2)
		pc = uint16(bytes[0]<<8) + uint16(bytes[1])
	} else {
		pc++
	}
	_m, _t = 2, 8
}
