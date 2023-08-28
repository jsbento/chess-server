package engine

import (
	t "github.com/jsbento/chess-server/cmd/engine/types"
)

type Engine struct {
	Board      *t.Board
	SetMask    [64]uint64
	ClearMask  [64]uint64
	PieceKeys  [13][120]uint64
	SideKey    uint64
	EnPasKey   uint64
	CastleKeys [16]uint64
}

func NewEngine() (e *Engine) {
	e = &Engine{
		Board:      t.NewBoard(),
		SetMask:    [64]uint64{},
		ClearMask:  [64]uint64{},
		PieceKeys:  [13][120]uint64{},
		SideKey:    0,
		EnPasKey:   0,
		CastleKeys: [16]uint64{},
	}
	e.InitBitmasks()
	e.InitHashKeys()
	return e
}

func (e *Engine) InitBitmasks() {
	for i := 0; i < 64; i++ {
		e.SetMask[i] = 0
		e.ClearMask[i] = 0
	}

	for i := 0; i < 64; i++ {
		e.SetMask[i] |= (uint64(1) << uint64(i))
		e.ClearMask[i] = ^e.SetMask[i]
	}
}

func (e *Engine) SetBit(bb *uint64, sq int) {
	*bb |= e.SetMask[sq]
}

func (e *Engine) ClearBit(bb *uint64, sq int) {
	*bb &= e.ClearMask[sq]
}
