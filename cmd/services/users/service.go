package users

import (
	"errors"
	"time"

	sT "github.com/jsbento/chess-server/cmd/server/types"
	t "github.com/jsbento/chess-server/cmd/services/users/types"
	"github.com/jsbento/chess-server/pkg/auth"
	m "github.com/jsbento/chess-server/pkg/mongo"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UsersColl string = "users"
)

type UserService struct {
	m *m.Store
}

func NewUserService(config *sT.ServerConfig) (*UserService, error) {
	store, err := m.NewStore(config.MongoHost, config.MongoDB)
	if err != nil {
		return nil, err
	}

	return &UserService{
		m: store,
	}, nil
}

func (s *UserService) Close() error {
	return s.m.Disconnect()
}

func (s *UserService) SignUp(req *t.CreateUserReq) (out *t.SignInResp, err error) {
	if user, err := s.GetUser(&t.GetUserReq{
		Username: &req.Username,
		Email:    &req.Email,
	}); err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	} else if user != nil {
		return nil, errors.New("username or email already exists")
	}

	err = s.CreateUser(req)
	if err != nil {
		return nil, err
	}

	user, err := s.GetUser(&t.GetUserReq{
		Username: &req.Username,
		Email:    &req.Email,
	})
	if err != nil {
		return nil, err
	}

	tkn, err := auth.GenToken(user.Id, user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	out = &t.SignInResp{
		User:  user,
		Token: tkn,
	}

	return
}

func (s *UserService) SignIn(req *t.SignInReq) (out *t.SignInResp, err error) {
	user, err := s.GetUser(&t.GetUserReq{
		Username: &req.Identifier,
		Email:    &req.Identifier,
	})
	if err != nil {
		return nil, err
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	tkn, err := auth.GenToken(user.Id, user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	out = &t.SignInResp{
		User:  user,
		Token: tkn,
	}

	return
}

func (s *UserService) CreateUser(req *t.CreateUserReq) error {
	col := s.m.Col(UsersColl)

	pwd, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	return s.m.Insert(col, &t.User{
		Id:        uuid.NewV4().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  pwd,
		CreatedAt: time.Now(),
	})
}

func (s *UserService) UpdateUser(user *t.UpdateUser) (out *t.User, err error) {
	col := s.m.Col(UsersColl)

	out = &t.User{}
	err = s.m.Update(col, m.M{"_id": user.Id}, user, out)

	return
}

func (s *UserService) GetUser(req *t.GetUserReq) (out *t.User, err error) {
	col := s.m.Col(UsersColl)

	q := m.M{}
	if req.Id != nil {
		q["_id"] = *req.Id
	}
	if req.Username != nil {
		q["username"] = *req.Username
	}
	if req.Email != nil {
		q["email"] = *req.Email
	}

	out = &t.User{}
	err = s.m.FindOne(col, q, out)

	return
}

func (s *UserService) SearchUsers(req *t.SearchUsersReq) (out []*t.User, err error) {
	col := s.m.Col(UsersColl)

	q := m.M{}
	if req.Ids != nil {
		q["_id"] = m.M{"$in": *req.Ids}
	}
	if req.Usernames != nil {
		q["username"] = m.M{"$in": *req.Usernames}
	}
	if req.Emails != nil {
		q["email"] = m.M{"$in": *req.Emails}
	}

	opts := options.Find()
	if req.Limit != 0 {
		opts.SetLimit(int64(req.Limit))
		opts.SetSkip(int64(req.Offset))
	}

	out = []*t.User{}
	err = s.m.Find(col, q, opts, &out)

	return
}

func (s *UserService) DeleteUser(id string) (out *t.User, err error) {
	col := s.m.Col(UsersColl)

	out = &t.User{}
	err = s.m.Delete(col, m.M{"_id": id}, out)

	return
}
