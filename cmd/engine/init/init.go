package init

import (
	c "github.com/jsbento/chess-server/pkg/constants"
	"github.com/jsbento/chess-server/pkg/utils"
)

func AllInit() {
	InitSq120To64()
	InitFilesRanksBrd()
}

func InitSq120To64() {
	for i := 0; i < c.BRD_SQ_NUM; i++ {
		c.Sq120ToSq64[i] = 65
	}

	for i := 0; i < 64; i++ {
		c.Sq64ToSq120[i] = 120
	}

	sq64 := 0
	for rank := c.RANK_1; rank <= c.RANK_8; rank++ {
		for file := c.FILE_A; file <= c.FILE_H; file++ {
			sq := utils.Fr2Sq(file, rank)
			c.Sq64ToSq120[sq64] = sq
			c.Sq120ToSq64[int(sq)] = sq64
			sq64++
		}
	}
}

func InitFilesRanksBrd() {
	for i := 0; i < c.BRD_SQ_NUM; i++ {
		c.FilesBrd[i] = int(c.OFFBOARD)
		c.RanksBrd[i] = int(c.OFFBOARD)
	}

	for rank := c.RANK_1; rank <= c.RANK_8; rank++ {
		for file := c.FILE_A; file <= c.FILE_H; file++ {
			sq := utils.Fr2Sq(file, rank)
			c.FilesBrd[sq] = int(file)
			c.RanksBrd[sq] = int(rank)
		}
	}
}
