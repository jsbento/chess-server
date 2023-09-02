package types

import (
	c "github.com/jsbento/chess-server/pkg/constants"
)

type PvEntry struct {
	PosKey uint64
	Move   int
}

type PvTable struct {
	PvEntries  []PvEntry
	NumEntries int
}

func NewPvTable(size int) (pvTable *PvTable) {
	numEntries := size - 2

	pvTable = &PvTable{
		PvEntries:  make([]PvEntry, size),
		NumEntries: numEntries,
	}
	pvTable.Clear()

	return
}

func (pv *PvTable) Clear() {
	for i := 0; i < pv.NumEntries; i++ {
		pv.PvEntries[i].PosKey = 0
		pv.PvEntries[i].Move = c.NOMOVE
	}
}
