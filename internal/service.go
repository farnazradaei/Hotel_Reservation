package internal

import (
	storage "Hotel_reservation/storage"
	"errors"
	auth "Hotel_reservation/Auth"
)

type User struct {
	ID    int
	Email string
	Hash  string
}

func Signup(email, password string) error {

    exists, err := storage.UserExsist(email)
    if err != nil {
        return err
    }

    if exists {
        return errors.New("user already exists")
    }
    
    hash , err := auth.HashPassword(password)
    if err!= nil {
        return err
    }

    id := storage.NextID()
    
    user := storage.User {
        ID:       id,
        Email:    email,
        Hash:     hash,
    }
    if err := storage.UserSave(user); err != nil {
        return err
    }
    return  nil
}

func Login(email, password string) error {
	user, err := storage.GetUserByEmail(email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(password, user.Hash) {
		return errors.New("invalid credentials")
	}

	return nil
}