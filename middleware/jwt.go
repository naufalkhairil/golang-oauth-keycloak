package middleware

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/naufalkhairil/golang-oauth-keycloak/config"
	"github.com/pkg/errors"
)

type JWTToken struct {
	Token  string
	Claims ExtendClaims
}

type ExtendClaims struct {
	StandardClaims jwt.StandardClaims
	Username       string `json:"Username"`
	Email          string `json:"Email"`
}

func (c ExtendClaims) Valid() error { return nil }

func GenerateJWT(email string) (*JWTToken, error) {
	keyDuration := config.GetJWTExpiredDuration() * time.Minute
	JWT_SIGNATURE_KEY := []byte(config.GetJWTSignKey())

	claims := ExtendClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "user",
			ExpiresAt: time.Now().Add(keyDuration).Unix(),
		},
		Email: email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to signed token")
	}

	return &JWTToken{
		Token:  signedToken,
		Claims: claims,
	}, nil
}
