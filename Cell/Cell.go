package cell

type Cell struct {
	alive         bool
	numNeighbours int64
}

func UpdateCell(cell *Cell) int64 {
	if cell.numNeighbours == 3 && cell.alive {
		cell.alive = true
		return 0
	} else if cell.numNeighbours == 3 {
		cell.alive = true
		return 1
	} else if cell.alive {
		if cell.numNeighbours == 2 {
			cell.alive = true
			return 0
		} else if cell.numNeighbours >= 4 || cell.numNeighbours < 2 {
			cell.alive = false
			return -1
		}
	} else {
		cell.alive = false
	}
	return 0
}

func CreateCell(neighbours int64) Cell {
	return Cell{alive: true, numNeighbours: neighbours}
}

type CellGrid struct {
	cells []Cell
}

func UpdateGrid(grid *CellGrid) {
	grid.cells[0] = CreateCell(0)

}
