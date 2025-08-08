package token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	publicKey *ecdsa.PublicKey
}

// ! Buradaki degerler supabase'den
func createECDSAPublicKey() *ecdsa.PublicKey {
	xBytes, _ := base64.RawURLEncoding.DecodeString("4Z-_yXVvMDUBLwyZwDhX5dWXKLBHmuoZDK4N06JCGdc")
	yBytes, _ := base64.RawURLEncoding.DecodeString("DpQAGT0bG7ZvzKvmHojlmG7NAHEWD3NaKDHQxA9p_7c")

	x := new(big.Int).SetBytes(xBytes)
	y := new(big.Int).SetBytes(yBytes)

	publicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	return publicKey
}

func NewManager() Manager {
	return Manager{
		publicKey: createECDSAPublicKey(),
	}
}

func (m *Manager) VerifySupabaseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.publicKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("wrong claims")
	}

	return claims, err
}
