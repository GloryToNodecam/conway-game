package cell

import (
	"testing"
)

func testUpdateCell(t *testing.T) {
	t.Log("testing UpdateCell")
	c := Cell{alive: true, numNeighbours: 6}
	UpdateCell(c)
	if c.alive {
		t.Error("Cell stayed alive; want died")
	}
}
