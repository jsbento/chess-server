package games

import (
	"time"

	t "github.com/jsbento/chess-server/cmd/server/services/games/types"
	sT "github.com/jsbento/chess-server/cmd/server/types"
	m "github.com/jsbento/chess-server/pkg/mongo"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	GamesColl string = "games"
)

type GameService struct {
	m *m.Store
}

func NewGameService(config *sT.ServerConfig) (*GameService, error) {
	store, err := m.NewStore(config.MongoHost, config.MongoDB)
	if err != nil {
		return nil, err
	}

	return &GameService{
		m: store,
	}, nil
}

func (s *GameService) Close() error {
	return s.m.Disconnect()
}

func (s *GameService) CreateGame(game *t.Game) error {
	col := s.m.Col(GamesColl)

	game.Id = uuid.NewV4().String()
	game.Date = time.Now()

	return s.m.Insert(col, game)
}

func (s *GameService) UpdateGame(game *t.UpdateGame) (out *t.Game, err error) {
	col := s.m.Col(GamesColl)

	out = &t.Game{}
	err = s.m.Update(col, m.M{"_id": game.Id}, game, out)

	return
}

func (s *GameService) GetGame(id string) (out *t.Game, err error) {
	col := s.m.Col(GamesColl)

	out = &t.Game{}
	err = s.m.FindOne(col, m.M{"_id": id}, out)

	return
}

func (s *GameService) SearchGames(req *t.SearchGamesReq) (out []*t.Game, err error) {
	col := s.m.Col(GamesColl)

	q := m.M{}
	if req.Ids != nil {
		q["_id"] = m.M{"$in": *req.Ids}
	}
	if req.PlayerId != nil {
		q["pId"] = *req.PlayerId
	}
	if req.PlayerIds != nil {
		q["pId"] = m.M{"$in": *req.PlayerIds}
	}
	if req.Result != nil {
		q["res"] = *req.Result
	}

	opts := options.Find()
	if req.Limit != 0 {
		opts.SetLimit(int64(req.Limit))
		opts.SetSkip(int64(req.Offset))
	}

	out = []*t.Game{}
	err = s.m.Find(col, q, opts, &out)

	return
}

func (s *GameService) DeleteGame(id string) (out *t.Game, err error) {
	col := s.m.Col(GamesColl)

	out = &t.Game{}
	err = s.m.Delete(col, m.M{"_id": id}, out)

	return
}
