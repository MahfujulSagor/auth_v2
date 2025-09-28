package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new NewJWTMaker
// The secret key must be at least 32 characters long
func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < 32 {
		return nil, fmt.Errorf("invalid key size: must be at least 32 characters")
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

// CreateToken creates a new token for a specific user
// and a specific duration (e.g., 15 minutes, 1 hour, etc.)
// It returns the token string and the user claims
// If there's an error, it returns an empty string and the error
func (maker *JWTMaker) CreateToken(id int64, username string, email string, duration time.Duration) (string, *UserClaims, error) {
	// Create the user claims with the provided information and duration
	claims, err := NewUserClaims(id, username, email, duration)
	if err != nil {
		return "", nil, err
	}

	// Create a new JWT token with the custom claims
	// Use the HS256 signing method (HMAC with SHA-256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret key to get the complete encoded token
	signedToken, err := token.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, claims, nil
}

// VerifyToken checks if the token is valid or not
// It returns the user claims if the token is valid
// If the token is invalid or expired, it returns an error
func (maker *JWTMaker) VerifyToken(tokenStr string) (*UserClaims, error) {
	// Parse the token with the provided token string and a function
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC and specifically HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte(maker.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Assert that the token claims are of type *UserClaims
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
