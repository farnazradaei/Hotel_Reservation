package internal

import (
	auth "Hotel_reservation/Auth"
	storage "Hotel_reservation/storage"
	"errors"
    "github.com/golang-jwt/jwt/v5"
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

func Login(email, password string) (string, error) {
	user, err := storage.GetUserByEmail(email)
    if err != nil {
		return "", err
	}

	// if user == nil {
	// 	return "", errors.New("invalid credentials")
	// }

	if !auth.CheckPasswordHash(password, user.Hash) {
		return "", errors.New("invalid credentials")
	}


    var jwtSecret = []byte("my-super-secret")

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), 
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}