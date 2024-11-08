package cell

import (
	"testing"
)

func TestUpdateCell(t *testing.T) {
	t.Log("testing UpdateCell")

	t.Run("Alive not enough neighbours", func(t *testing.T) {
		c := Cell{alive: true, numNeighbours: 0}
		r := UpdateCell(&c)
		if c.alive {
			t.Error("Cell alive; want died")
		}
		if r != -1 {
			t.Errorf("Returned %d; want -1", r)
		}
	})
	t.Run("Alive enough neighbours", func(t *testing.T) {
		c := Cell{alive: true, numNeighbours: 3}
		r := UpdateCell(&c)
		if !c.alive {
			t.Error("Cell died; want alive")
		}
		if r != 0 {
			t.Errorf("Returned %d; want 0", r)
		}
	})
	t.Run("Alive too many neighbours", func(t *testing.T) {
		c := Cell{alive: true, numNeighbours: 6}
		r := UpdateCell(&c)
		if c.alive {
			t.Error("Cell alive; want died")
		}
		if r != -1 {
			t.Errorf("Returned %d; want -1", r)
		}
	})
	t.Run("dead not 3 neighbours", func(t *testing.T) {
		c := Cell{alive: false, numNeighbours: 2}
		r := UpdateCell(&c)
		if c.alive {
			t.Error("Cell alive; want dead")
		}
		if r != 0 {
			t.Errorf("Returned %d; want 0", r)
		}
	})
	t.Run("dead 3 neighbours", func(t *testing.T) {
		c := Cell{alive: false, numNeighbours: 3}
		r := UpdateCell(&c)
		if !c.alive {
			t.Error("Cell died; want alive")
		}
		if r != 1 {
			t.Errorf("Returned %d; want 1", r)
		}
	})
}
