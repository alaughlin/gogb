package main

import "fmt"

var (
	gb  *GameBoy
	mmu *MMU
)

func main() {
	gb = &GameBoy{}

	gb.reset()
	gb.loadGame()

	for i := 0; i < 100; i++ {
		fmt.Println(mmu.read(pc))
		gb.execute()
	}
}
