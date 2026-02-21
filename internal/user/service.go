package user

import (
	"Hotel_reservation/internal/auth"
	"errors"
)

type User struct {
	ID    int
	Email string
	Hash  string
}

type Service struct {
    Storage UserStorage
}

func NewService(storage UserStorage) *Service {
    return &Service{Storage: storage}
}

func (s *Service)Signup(email, password string) error {

    exists, err := s.Storage.UserExsist(email)
    if err != nil {
        return err
    }

    if exists {
        return errors.New("user already exists")
    }
    
    hash , err := HashPassword(password)
    if err!= nil {
        return err
    }
    
    user := User {
        Email:    email,
        Hash:     hash,
    }
    if err := s.Storage.UserSave(user); err != nil {
        return err
    }
    return  nil
}

func (s *Service)Login(email, password string) (string, error) {
	user, err := s.Storage.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !CheckPasswordHash(password, user.Hash) {
		return "", errors.New("invalid credentials")
	}

	return auth.GenerateToken(user.ID, user.Email)
}