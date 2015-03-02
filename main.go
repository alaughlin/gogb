package main

var (
	gb  *GameBoy
	mmu *MMU
)

func main() {
	gb = &GameBoy{}

	gb.reset()
	gb.loadGame()
}
