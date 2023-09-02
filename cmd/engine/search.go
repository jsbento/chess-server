package engine

import (
	"fmt"

	t "github.com/jsbento/chess-server/cmd/engine/types"
	c "github.com/jsbento/chess-server/pkg/constants"
	"github.com/jsbento/chess-server/pkg/utils"
)

func (e *Engine) PickNextMove(moveNum int, list *t.MoveList) {
	bestScore := 0
	bestNum := moveNum
	for i := moveNum; i < list.Count; i++ {
		if list.Moves[i].Score > bestScore {
			bestScore = list.Moves[i].Score
			bestNum = i
		}
	}
	temp := list.Moves[moveNum]
	list.Moves[moveNum] = list.Moves[bestNum]
	list.Moves[bestNum] = temp
}

func (e *Engine) IsRepetition() bool {
	for i := e.Board.HisPly - e.Board.FiftyMove; i < e.Board.HisPly-1; i++ {
		if e.Board.PosKey == e.Board.History[i].PosKey {
			return true
		}
	}
	return false
}

func (e *Engine) CheckUp() {

}

func (e *Engine) ClearForSearch(info *t.SearchInfo) {
	for i := 0; i < 13; i++ {
		for j := 0; j < c.BRD_SQ_NUM; j++ {
			e.Board.SearchHistory[i][j] = 0
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < c.MAX_DEPTH; j++ {
			e.Board.SearchKillers[i][j] = 0
		}
	}

	e.Board.PvTable.Clear()
	e.Board.Ply = 0

	info.StartTime = utils.GetTimeMs()
	info.Stopped = false
	info.Nodes = 0
	info.Fh = 0.0
	info.Fhf = 0.0
}

func (e *Engine) Quiescence(alpha, beta int, info *t.SearchInfo) int {
	info.Nodes++

	if e.IsRepetition() || e.Board.FiftyMove >= 100 {
		return 0
	}

	if e.Board.Ply > c.MAX_DEPTH-1 {
		return e.EvalPosition()
	}

	score := e.EvalPosition()

	if score >= beta {
		return beta
	}

	if score > alpha {
		alpha = score
	}

	list := t.NewMoveList()
	e.GenerateAllCaptures(list)

	legal := 0
	oldAlpha := alpha
	bestMove := c.NOMOVE
	score = -c.INFINITE

	for moveNum := 0; moveNum < list.Count; moveNum++ {
		e.PickNextMove(moveNum, list)
		if !e.MakeMove(list.Moves[moveNum].Move) {
			continue
		}

		legal++
		score = -e.Quiescence(-beta, -alpha, info)
		e.TakeMove()

		if score > alpha {
			if score >= beta {
				if legal == 1 {
					info.Fhf++
				}
				info.Fh++

				return beta
			}
			alpha = score
			bestMove = list.Moves[moveNum].Move
		}
	}

	if alpha != oldAlpha {
		e.StorePvMove(bestMove)
	}

	return alpha
}

func (e *Engine) AlphaBeta(alpha, beta, depth int, info *t.SearchInfo, doNull bool) int {
	if depth == 0 {
		info.Nodes++
		return e.Quiescence(alpha, beta, info)
	}
	info.Nodes++

	if e.IsRepetition() || e.Board.FiftyMove >= 100 {
		return 0
	}

	if e.Board.Ply > c.MAX_DEPTH-1 {
		return e.EvalPosition()
	}

	list := t.NewMoveList()
	e.GenerateAllMoves(list)

	legal := 0
	oldAlpha := alpha
	bestMove := c.NOMOVE
	score := -c.INFINITE

	pvMove := e.ProbePvTable()
	if pvMove != c.NOMOVE {
		for moveNum := 0; moveNum < list.Count; moveNum++ {
			if list.Moves[moveNum].Move == pvMove {
				list.Moves[moveNum].Score = 2000000
				break
			}
		}
	}

	for moveNum := 0; moveNum < list.Count; moveNum++ {
		e.PickNextMove(moveNum, list)
		if !e.MakeMove(list.Moves[moveNum].Move) {
			continue
		}

		legal++
		score = -e.AlphaBeta(-beta, -alpha, depth-1, info, true)
		e.TakeMove()

		if score > alpha {
			if score >= beta {
				if legal == 1 {
					info.Fhf++
				}
				info.Fh++

				if list.Moves[moveNum].Move&int(c.MFLAGCAP) == 0 {
					e.Board.SearchKillers[1][e.Board.Ply] = e.Board.SearchKillers[0][e.Board.Ply]
					e.Board.SearchKillers[0][e.Board.Ply] = list.Moves[moveNum].Move
				}

				return beta
			}
			alpha = score
			bestMove = list.Moves[moveNum].Move

			if list.Moves[moveNum].Move&int(c.MFLAGCAP) == 0 {
				e.Board.SearchHistory[e.Board.Pieces[utils.FromSq(bestMove)]][utils.ToSq(bestMove)] += depth
			}
		}
	}

	if legal == 0 {
		if e.IsSqAttacked(c.Square(e.Board.KingSq[e.Board.Side]), e.Board.Side^1) {
			return -c.MATE + e.Board.Ply
		} else {
			return 0
		}
	}

	if alpha != oldAlpha {
		e.StorePvMove(bestMove)
	}

	return alpha
}

func (e *Engine) SearchPosition(info *t.SearchInfo) {
	bestMove := c.NOMOVE
	bestScore := -c.INFINITE
	pvMoves := 0
	pvNum := 0

	e.ClearForSearch(info)

	for currDepth := 1; currDepth <= info.Depth; currDepth++ {
		bestScore = e.AlphaBeta(-c.INFINITE, c.INFINITE, currDepth, info, true)
		pvMoves = e.GetPvLine(currDepth)
		bestMove = e.Board.PvArray[0]
		fmt.Printf("Depth: %d, Score: %d, Nodes: %d, Move: %s\n", currDepth, bestScore, info.Nodes, utils.PrintMove(bestMove))
		fmt.Print("PV:")
		for pvNum = 0; pvNum < pvMoves; pvNum++ {
			fmt.Printf(" %s", utils.PrintMove(e.Board.PvArray[pvNum]))
		}
		fmt.Println()
		fmt.Printf("Ordering: %.2f\n", float64(info.Fhf)/float64(info.Fh))
	}
}
