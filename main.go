package main

import (
	"fmt"
	"time"
)

var (
	gb  *GameBoy
	mmu *MMU
)

func main() {
	gb = &GameBoy{}

	gb.reset()
	gb.loadGame()

	for {
		fmt.Println(mmu.read(pc))
		gb.execute()
		time.Sleep(10)
	}
}
