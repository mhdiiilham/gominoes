package user

import "github.com/mhdiiilham/gominoes/pkg/password"

type manager struct {
	Repo Repository
	Pwd  password.Hasher
}

// NewManager create new repository
func NewManager(r Repository, pwd password.Hasher) *manager {
	return &manager{
		Repo: r,
		Pwd:  pwd,
	}
}

// Register new user
func (s *manager) Register(user User) (*User, error) {
	hashPassword := s.Pwd.Hash(user.Password)
	user.Password = hashPassword
	return s.Repo.Register(user)
}

// FindOne user
func (s *manager) FindOne(email string) (*User, error) {
	return s.Repo.FindOne(email)
}
