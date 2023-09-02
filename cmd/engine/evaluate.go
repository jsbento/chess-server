package engine

import (
	c "github.com/jsbento/chess-server/pkg/constants"
	"github.com/jsbento/chess-server/pkg/utils"
)

var PawnTable [64]int = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	10, 10, 0, -10, -10, 0, 10, 10,
	5, 0, 0, 5, 5, 0, 0, 5,
	0, 0, 10, 20, 20, 10, 0, 0,
	5, 5, 5, 10, 10, 5, 5, 5,
	10, 10, 10, 20, 20, 10, 10, 10,
	20, 20, 20, 30, 30, 20, 20, 20,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var KnightTable [64]int = [64]int{
	0, -10, 0, 0, 0, 0, -10, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 5, 0,
	5, 10, 15, 20, 20, 15, 10, 5,
	5, 10, 10, 20, 20, 10, 10, 5,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}
var BishopTable [64]int = [64]int{
	0, 0, -10, 0, 0, -10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var RookTable [64]int = [64]int{
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	25, 25, 25, 25, 25, 25, 25, 25,
	0, 0, 5, 10, 10, 5, 0, 0,
}

var KingEndgame [64]int = [64]int{
	-50, -10, 0, 0, 0, 0, -10, -50,
	-10, 0, 10, 10, 10, 10, 0, -10,
	0, 10, 15, 15, 15, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 15, 15, 15, 10, 0,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-50, -10, 0, 0, 0, 0, -10, -50,
}

var KingTable [64]int = [64]int{
	0, 5, 5, -10, -10, 0, 10, 5,
	-30, -30, -30, -30, -30, -30, -30, -30,
	-50, -50, -50, -50, -50, -50, -50, -50,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
}

var _Mirror64 [64]int = [64]int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7,
}

func Mirror64(sq int) int {
	return _Mirror64[sq]
}

func (e *Engine) EvalPosition() int {
	score := e.Board.Material[c.WHITE] - e.Board.Material[c.BLACK]

	pce := c.WP
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score += PawnTable[utils.Sq64(sq)]
	}

	pce = c.BP
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score -= PawnTable[Mirror64(utils.Sq64(sq))]
	}

	pce = c.WN
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score += KnightTable[utils.Sq64(sq)]
	}

	pce = c.BN
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score -= KnightTable[Mirror64(utils.Sq64(sq))]
	}

	pce = c.WB
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score += BishopTable[utils.Sq64(sq)]
	}

	pce = c.BB
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score -= BishopTable[Mirror64(utils.Sq64(sq))]
	}

	pce = c.WR
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score += RookTable[utils.Sq64(sq)]
	}

	pce = c.BR
	for i := 0; i < e.Board.PceNum[pce]; i++ {
		sq := e.Board.Plist[pce][i]
		score -= RookTable[Mirror64(utils.Sq64(sq))]
	}

	if e.Board.Side == c.WHITE {
		return score
	} else {
		return -score
	}
}
