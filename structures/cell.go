package structures

type Cell struct {
	alive bool
}

func (c *Cell) IsAlive() bool {
	return c.alive
}

func (c *Cell) SetAlive(alive bool) {
	c.alive = alive
}

func (c Cell) Step2(neighbours []Cell) Cell {
	aliveNeighbours := 0

	for _, n := range neighbours {
		if n.IsAlive() {
			aliveNeighbours++
		}
	}
	newCell := c.Copy()
	if c.IsAlive() {
		if aliveNeighbours > 3 || aliveNeighbours < 2 {
			newCell.SetAlive(false)
		}
	} else {
		if aliveNeighbours == 3 {
			newCell.SetAlive(true)
		}
	}
	return newCell
}
func (c Cell) String() string {
	if c.IsAlive() {
		return " # "
	} else {
		return "   "
	}
}

func (c Cell) Copy() Cell {
	var newCell Cell

	newCell.SetAlive(c.IsAlive())

	return newCell
}
