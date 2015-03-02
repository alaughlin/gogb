package main

/* misc & control */

// 00 NOP
func NOP() {
	pc++
	_m, _t = 1, 4
}

// 10 STOP 0
// TODO
func STOP0() {
	// halt cpu and lcd until button is pressed
	pc += 2
	_m, _t = 1, 4
}

// 76 HALT
// TODO
func HALT() {
	// halt cpu until interrupt is received
	pc++
	_m, _t = 1, 4
}

// CB PREFIX CB
// TODO
func PREFIXCB() {
	pc++
	_m, _t = 1, 4
}

// DI
// TODO
func DI() {
	// disable interrupts after next instruction
	pc++
	_m, _t = 1, 4
}

// EI
// TODO
func EI() {
	// enable interrupts after next instruction
	pc++
	_m, _t = 1, 4
}
