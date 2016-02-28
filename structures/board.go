package structures

import "fmt"
import "math/rand"
import "time"

type Board struct {
	Rows    int
	Columns int
	Cells   [][]Cell
}

func createCells(Rows, Columns int) [][]Cell {
	cells := make([][]Cell, Rows)
	for i := 0; i < Rows; i++ {
		cells[i] = make([]Cell, Columns)
	}
	return cells
}
func NewBoard(Rows, Columns int) Board {
	var board Board
	board.Rows = Rows
	board.Columns = Columns
	board.Cells = createCells(Rows, Columns)
	return board
}
func (board *Board) HackerEmblemSeed() {
	board.Cells[0][2].SetAlive(true)
	board.Cells[1][0].SetAlive(true)
	board.Cells[1][2].SetAlive(true)
	board.Cells[2][1].SetAlive(true)
	board.Cells[2][2].SetAlive(true)
}
func (board *Board) Seed() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	times := 0
	for times == 0 {
		times = r.Intn(board.Rows * board.Columns)
	}

	for t := 0; t < times; t++ {
		i, j := rand.Intn(board.Rows), rand.Intn(board.Columns)

		board.Cells[i][j].SetAlive(true)
	}
}
func (board Board) GetNeighbours(r, c int) []Cell {
	var cells []Cell
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			newR, newC := r+i, c+j
			if newR >= board.Rows {
				newR = 0
			} else if newR < 0 {
				newR = board.Rows - 1
			}
			if newC >= board.Columns {
				newC = 0
			} else if newC < 0 {
				newC = board.Columns - 1
			}
			if !(newR == r && newC == c) {
				cells = append(cells, board.Cells[newR][newC])
			}
		}
	}
	return cells
}
func (board Board) Step() Board {
	var newCells = createCells(board.Rows, board.Columns)

	for i := 0; i < board.Rows; i++ {
		for j := 0; j < board.Columns; j++ {
			newCells[i][j] = board.Cells[i][j].Step2(board.GetNeighbours(i, j))
		}
	}

	board.Cells = newCells
	return board
}
func (b Board) String() string {
	for _, v := range b.Cells {
		for _, c := range v {
			fmt.Printf("%v", c)
		}
		fmt.Println("")
	}
	return ""
}
