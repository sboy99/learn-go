package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sboy99/learn-go/go-todo/infra/config"
)

// ------------------------------STRUCTS---------------------------------- //

type JwtAdapter struct{}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (a *JwtAdapter) Sign(payload map[string]interface{}) (string, error) {
	claims := a.getClaims(payload)
	token, err := a.signClaims(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ------------------------------PRIVATE_METHODS---------------------------------- //

func (a *JwtAdapter) getClaims(payload map[string]interface{}) jwt.MapClaims {
	claims := make(jwt.MapClaims)

	// map payload to claims
	for k, v := range payload {
		claims[k] = v
	}

	// add timestamp to claims
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// done
	return claims
}

func (a *JwtAdapter) signClaims(claims jwt.MapClaims) (string, error) {
	secret := a.getSecret("access_token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (a *JwtAdapter) getSecret(secretFor string) []byte {
	jwtSecret := config.JWT_SECRET
	switch secretFor {
	case "access_token":
		return []byte(jwtSecret)
	default:
		return []byte(jwtSecret)
	}
}
