package types

import "errors"

type GetUserReq struct {
	Id       *string `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

type SearchUsersReq struct {
	Ids       *[]string `json:"ids"`
	Usernames *[]string `json:"usernames"`
	Emails    *[]string `json:"emails"`
	Limit     int       `json:"limit"`
	Offset    int       `json:"offset"`
}

type SignInReq struct {
	Identifier string `json:"identifier"` // username or email
	Password   string `json:"password"`
}

func (s SignInReq) Validate() error {
	if s.Identifier == "" {
		return errors.New("username or email is required")
	}
	if s.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type SignInResp struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s CreateUserReq) Validate() error {
	if s.Username == "" {
		return errors.New("username is required")
	}
	if s.Email == "" {
		return errors.New("email is required")
	}
	if s.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
