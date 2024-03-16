package bcrypt

import (
	lib_bcrypt "golang.org/x/crypto/bcrypt"
)

type Interface interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
}

type bcrypt struct {
	cost int
}

func Init() Interface {
	return &bcrypt{
		cost: 10,
	}
}

func (b *bcrypt) GenerateFromPassword(password string) (string, error) {
	bytePassword, err := lib_bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func (b *bcrypt) CompareHashAndPassword(hashedPassword string, password string) error {
	err := lib_bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
