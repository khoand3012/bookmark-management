package service

import (
	"crypto/rand"
	"math/big"
)

//go:generate mockery --name=GenPass --filename genpass.go --outpkg service_mocks
type GenPass interface {
	GeneratePassword(length int) (string, error)
}

type genPass struct{}

func NewGenPass() GenPass {
	return &genPass{}
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func (g *genPass) GeneratePassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}
	return string(password), nil
}
