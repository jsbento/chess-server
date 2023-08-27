package types

import (
	c "github.com/jsbento/chess-server/pkg/constants"
)

type Board struct {
	Pieces [c.BRD_SQ_NUM]c.Piece
	Pawns  [3]uint64
	KingSq [2]int

	Side      c.Side
	EnPas     int
	FiftyMove int

	Ply    int
	HisPly int

	PosKey uint64

	PceNum [13]int
	BigPce [3]int
	MajPce [3]int
	MinPce [3]int
}

func NewBoard() *Board {
	return &Board{
		Pieces:    [c.BRD_SQ_NUM]c.Piece{},
		Pawns:     [3]uint64{},
		KingSq:    [2]int{},
		Side:      c.BOTH,
		EnPas:     c.SQ_NONE,
		FiftyMove: 0,
		Ply:       0,
		HisPly:    0,
		PosKey:    0,
		PceNum:    [13]int{},
		BigPce:    [3]int{},
		MajPce:    [3]int{},
		MinPce:    [3]int{},
	}
}
