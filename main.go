package main

import "github.com/gafr/gonway/structures"
import "fmt"
import "time"

func main() {

	board := structures.NewBoard(80, 80)
	//board.HackerEmblemSeed()
	board.Seed()
	for {
		board = board.Step()
		fmt.Printf("%v", board)
		fmt.Printf("\033[0;0H")
		Sleep()
		//fmt.Print("\033[2J")
	}
}

func Sleep() {
	time.Sleep(time.Duration(100) * time.Millisecond)
}
