package cell

type Cell struct {
	alive         bool
	numNeighbours int64
}

func updateCell(cell Cell) {
	if cell.numNeighbours < 2 {
		cell.alive = false
	} else if cell.alive {
		if cell.numNeighbours == 2 || cell.numNeighbours == 3 {
			cell.alive = true
		} else if cell.numNeighbours >= 4 {
			cell.alive = false
		}
	} else if cell.numNeighbours == 3 {
		cell.alive = true
	}
}
