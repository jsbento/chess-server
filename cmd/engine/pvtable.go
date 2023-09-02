package engine

import (
	c "github.com/jsbento/chess-server/pkg/constants"
)

func (e *Engine) StorePvMove(move int) {
	idx := e.Board.PosKey % uint64(e.Board.PvTable.NumEntries)

	e.Board.PvTable.PvEntries[idx].PosKey = e.Board.PosKey
	e.Board.PvTable.PvEntries[idx].Move = move
}

func (e *Engine) ProbePvTable() int {
	idx := e.Board.PosKey % uint64(e.Board.PvTable.NumEntries)

	if e.Board.PvTable.PvEntries[idx].PosKey == e.Board.PosKey {
		return e.Board.PvTable.PvEntries[idx].Move
	}

	return c.NOMOVE
}

func (e *Engine) GetPvLine(depth int) int {
	move := e.ProbePvTable()
	count := 0

	for move != c.NOMOVE && count < depth {
		if e.MoveExists(move) {
			e.MakeMove(move)
			e.Board.PvArray[count] = move
			count++
		} else {
			break
		}
		move = e.ProbePvTable()
	}

	for e.Board.Ply > 0 {
		e.TakeMove()
	}

	return count
}
