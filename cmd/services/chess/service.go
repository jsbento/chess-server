package chess

import (
	e "github.com/jsbento/chess-server/cmd/engine"
	eT "github.com/jsbento/chess-server/cmd/engine/types"
	sT "github.com/jsbento/chess-server/cmd/server/types"
	t "github.com/jsbento/chess-server/cmd/services/chess/types"
	m "github.com/jsbento/chess-server/pkg/mongo"
)

const (
	ChessCol string = "chess"
)

type ChessService struct {
	m *m.Store
}

func NewChessService(config *sT.ServerConfig) (*ChessService, error) {
	store, err := m.NewStore(config.MongoHost, config.MongoDB)
	if err != nil {
		return nil, err
	}

	return &ChessService{
		m: store,
	}, nil
}

func (s *ChessService) Close() error {
	return s.m.Disconnect()
}

func (s *ChessService) EvaluatePositon(req *t.EvalPosReq) (int, error) {
	engine := e.NewEngine()
	if err := engine.ParseFEN(req.Fen); err != nil {
		return 0, err
	}

	return engine.EvalPosition(), nil
}

func (s *ChessService) SearchPosition(req *t.SearchPosReq) (string, error) {
	engine := e.NewEngine()
	if err := engine.ParseFEN(req.Fen); err != nil {
		return "", err
	}

	info := &eT.SearchInfo{}
	return engine.ParseGo(req.ToGoCmd(), info), nil
}
