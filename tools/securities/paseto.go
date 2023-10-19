package securities

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

// NewPasseto
func NewPaseto(payload map[string]string, signingKey string, secondsDuration int) (string, error) {

	seconds := time.Duration(secondsDuration) * time.Second

	token := paseto.NewToken()
	for key, value := range payload {
		token.SetString(key, value)
	}

	token.SetExpiration(time.Now().Add(seconds))
	token.SetNotBefore(time.Now())
	token.SetIssuedAt(time.Now())

	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(signingKey)
	if err != nil {
		return "", err
	}

	return token.V4Sign(secretKey, nil), nil
}

// ParseJWT verifies and parses token and returns its claims.
func ParsePaseto(signed string, publicKey string) (map[string]any, error) {
	key, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKey)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	token, err := parser.ParseV4Public(key, signed, nil)
	if err != nil {
		return nil, err
	}

	return token.Claims(), nil
}
