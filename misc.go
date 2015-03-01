package main

/* misc & control */

// 00 NOP
func NOP() {
	pc++
}

// 10 STOP 0
// TODO
func STOP0() {
	// halt cpu and lcd until button is pressed
	pc += 2
}

// 76 HALT
// TODO
func HALT() {
	// halt cpu until interrupt is received
	pc++
}

// CB PREFIX CB
// TODO
func PREFIXCB() {
	pc++
}

// DI
// TODO
func DI() {
	// disable interrupts after next instruction
	pc++
}

// EI
// TODO
func EI() {
	// enable interrupts after next instruction
	pc++
}
