package types

import (
	c "github.com/jsbento/chess-server/pkg/constants"
)

type Move struct {
	Move  int
	Score int
}

type Undo struct {
	Move       int
	CastlePerm c.CastlePerm
	EnPas      c.Square
	FiftyMove  int
	PosKey     uint64
}

type Board struct {
	Pieces [c.BRD_SQ_NUM]c.Piece
	Pawns  [3]uint64
	KingSq [2]int

	Side       c.Side
	EnPas      c.Square
	FiftyMove  int
	CastlePerm c.CastlePerm

	Ply    int
	HisPly int

	PosKey uint64

	PceNum   [13]int
	BigPce   [2]int
	MajPce   [2]int
	MinPce   [2]int
	Material [2]int

	History [c.MAX_GAME_MOVES]*Undo

	Plist [13][10]c.Square
}

func NewBoard() *Board {
	return &Board{
		Pieces:     [c.BRD_SQ_NUM]c.Piece{},
		Pawns:      [3]uint64{},
		KingSq:     [2]int{},
		Side:       c.BOTH,
		EnPas:      c.NO_SQ,
		FiftyMove:  0,
		CastlePerm: 0,
		Ply:        0,
		HisPly:     0,
		PosKey:     0,
		PceNum:     [13]int{},
		BigPce:     [2]int{},
		MajPce:     [2]int{},
		MinPce:     [2]int{},
		Material:   [2]int{},
		History:    [c.MAX_GAME_MOVES]*Undo{},
		Plist:      [13][10]c.Square{},
	}
}
