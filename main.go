package main

import (
	"fmt"

	e "github.com/jsbento/chess-server/cmd/engine"
	bb "github.com/jsbento/chess-server/cmd/engine/bitboards"
	i "github.com/jsbento/chess-server/cmd/engine/init"
	c "github.com/jsbento/chess-server/pkg/constants"
)

const (
	START_FEN string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	FEN_1     string = "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	FEN_2     string = "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"
	FEN_3     string = "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"
	FEN_4     string = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 0 1"
)

func main() {
	i.AllInit()

	engine := e.NewEngine()

	engine.ParseFEN(FEN_4)
	bb.PrintBitBoard(engine.Board.Pawns[c.WHITE])
	fmt.Println()
	bb.PrintBitBoard(engine.Board.Pawns[c.BLACK])
	fmt.Println()
	bb.PrintBitBoard(engine.Board.Pawns[c.BOTH])
}
