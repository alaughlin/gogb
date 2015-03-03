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
		pc = uint16(bytes[1]<<8) + uint16(bytes[0])
	} else {
		pc++
	}
	_m, _t = 2, 8
}

// C2 JP NZ,a16
func JPBZa16(a uint16) {
	if _f&(1<<7) == 0 {
		pc = a
	} else {
		pc++
	}
	_m, _t = 3, 12
}

// C3 JP a16
func JPa16(a uint16) {
	pc = a
	_m, _t = 3, 12
}

// C4 CALL NZ,a16
func CALLNZa16(a uint16) {
	if _f&(1<<7) == 0 {
		push(uint8(pc+3)>>8, uint8(pc+3))
		pc = a
	} else {
		pc += 3
	}
	_m, _t = 3, 12
}

// C7 RST 00H
// TODO
func RST00H() {}

// C8 RET Z
func RETZ() {
	if _f&(1<<7) > 0 {
		bytes := pop(2)
		pc = uint16(bytes[1]<<8) + uint16(bytes[0])
	} else {
		pc++
	}
	_m, _t = 2, 8
}

// C9 RET
func RET() {
	bytes := pop(2)
	pc = uint16(bytes[1]<<8) + uint16(bytes[0])
	_m, _t = 2, 8
}

// CA JP Z,a16
func JPZa16(a uint16) {
	if _f&(1<<7) > 0 {
		pc = a
	} else {
		pc += 3
	}
	_m, _t = 3, 12
}

// CC CALL Z,a16
func CALLZa16(a uint16) {
	if _f&(1<<7) > 0 {
		push(uint8(pc+3)>>8, uint8(pc+3))
		pc = a
	} else {
		pc += 3
	}
	_m, _t = 3, 12
}

// CD CALL a16
func CALLa16(a uint16) {
	push(uint8(pc+3)>>8, uint8(pc+3))
	pc = a
	_m, _t = 3, 12
}

// CF RST 08H
// TODO
func RST08() {}

// D0 RET NC
func RETNC() {
	if _f&(1<<4) == 0 {
		bytes := pop(2)
		pc = uint16(bytes[1]<<8) + uint16(bytes[0])
	} else {
		pc++
	}
	_m, _t = 2, 8
}

// D2 JP NC,a16
func JPNCa16(a uint16) {
	if _f&(1<<4) == 0 {
		pc = a
	} else {
		pc++
	}
	_m, _t = 3, 12
}

// D4 CALL NC,a16
func CALLNCa16(a uint16) {
	if _f&(1<<4) == 0 {
		push(uint8(pc + 3))
		pc = a
	} else {
		pc += 3
	}
	_m, _t = 3, 12
}

// D7 RST 10H
// TODO
func RST10H() {}

// D8 RET C
func RETC() {
	if _f&(1<<4) > 0 {
		bytes := pop(2)
		pc = uint16(bytes[1]<<8) + uint16(bytes[0])
	} else {
		pc++
	}
	_m, _t = 2, 8
}

// D9 RETI
// TODO
func RETI() {
	bytes := pop(2)
	pc = uint16(bytes[1]<<8) + uint16(bytes[0])
	// enable interrupts
	_m, _t = 2, 8
}

// DA JP C,a16
func JPCa16(a uint16) {
	if _f&(1<<4) > 0 {
		pc = a
	} else {
		pc++
	}
	_m, _t = 3, 12
}

// DC CALL C,a16
func CALLCa16(a uint16) {
	if _f&(1<<4) > 0 {
		push(uint8(pc + 3))
		pc = a
	} else {
		pc += 3
	}
	_m, _t = 3, 12
}

// DF RST 18H
// TODO
func RST18H() {}

// E7 RST 20H
func RST20H() {}

// E9 JP (HL)
func JPHL() {
	pc = uint16(_h<<8) + uint16(_l)
	_m, _t = 1, 4
}

// EF RST 28H
// TODO
func RST28H() {}

// F7 RST 30H
// TODO
func RST30H() {}

// FF RST 38H
// TODO
func RST38H() {}
