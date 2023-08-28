package engine

import (
	d "github.com/jsbento/chess-server/cmd/engine/data"
	c "github.com/jsbento/chess-server/pkg/constants"
)

func IsKn(pce int) bool {
	return d.PieceKnight[pce]
}

func IsKi(pce int) bool {
	return d.PieceKing[pce]
}

func IsRQ(pce int) bool {
	return d.PieceRookQueen[pce]
}

func IsBQ(pce int) bool {
	return d.PieceBishopQueen[pce]
}

func (e *Engine) IsSqAttacked(sq c.Square, side c.Side) bool {
	if side == c.WHITE {
		if e.Board.Pieces[sq-11] == c.WP || e.Board.Pieces[sq-9] == c.WP {
			return true
		}
	} else {
		if e.Board.Pieces[sq+11] == c.BP || e.Board.Pieces[sq+9] == c.BP {
			return true
		}
	}

	for i := 0; i < 8; i++ {
		pce := e.Board.Pieces[sq+c.Square(c.KnDir[i])]
		if pce != c.Piece(c.OFFBOARD) && IsKn(int(pce)) && d.PieceCol[int(pce)] == side {
			return true
		}
	}

	for i := 0; i < 4; i++ {
		dir := c.RkDir[i]
		tSq := sq + c.Square(dir)
		pce := e.Board.Pieces[tSq]
		for pce != c.Piece(c.OFFBOARD) {
			if pce != c.Piece(c.EMPTY) {
				if IsRQ(int(pce)) && d.PieceCol[int(pce)] == side {
					return true
				}
				break
			}
			tSq += c.Square(dir)
			pce = e.Board.Pieces[tSq]
		}
	}

	for i := 0; i < 4; i++ {
		dir := c.BiDir[i]
		tSq := sq + c.Square(dir)
		pce := e.Board.Pieces[tSq]
		for pce != c.Piece(c.OFFBOARD) {
			if pce != c.Piece(c.EMPTY) {
				if IsBQ(int(pce)) && d.PieceCol[int(pce)] == side {
					return true
				}
				break
			}
			tSq += c.Square(dir)
			pce = e.Board.Pieces[tSq]
		}
	}

	for i := 0; i < 8; i++ {
		pce := e.Board.Pieces[sq+c.Square(c.KiDir[i])]
		if pce != c.Piece(c.OFFBOARD) && IsKi(int(pce)) && d.PieceCol[int(pce)] == side {
			return true
		}
	}

	return false
}
