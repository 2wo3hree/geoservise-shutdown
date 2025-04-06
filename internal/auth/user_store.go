package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

var users = make(map[string]string)
var mu sync.Mutex

func RegisterUser(username, password string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[username]; exists {
		return errors.New("такой юзер уже есть")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("ошибка при хешировании пароля пользователя")
	}

	users[username] = string(hashed)
	return nil
}

func AuthenticateUser(username, password string) error {
	mu.Lock()
	hashed, exists := users[username]
	mu.Unlock()

	if !exists {
		return errors.New("такого пользователя нету")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return errors.New("неверный пароль")
	}

	return nil
}
