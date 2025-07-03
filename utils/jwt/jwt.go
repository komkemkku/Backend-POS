package jwt

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"Backend-POS/model"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(raw string) (map[string]any, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (
		interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token singing method")
		}
		secret := []byte(os.Getenv("TOKEN_SECRET"))
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}

func GenerateTokenStaff(ctx context.Context, staff *model.Staff) (string, error) {
	tokenDurationStr := os.Getenv("TOKEN_DURATION")
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"sub": jwt.MapClaims{
			"id":            staff.ID,
			"username":      staff.UserName,
			"password_hash": staff.PasswordHash,
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(tokenDuration).Unix(),
	})

	secret := []byte(os.Getenv("TOKEN_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	return tokenString, nil
}

// func GenerateTokenAdmin(ctx context.Context, admin *model.Admins) (string, error) {
// 	godotenv.Load()
// 	tokenDurationStr := os.Getenv("TOKEN_DURATION")
// 	tokenDuration, err := time.ParseDuration(tokenDurationStr)
// 	if err != nil {
// 		log.Printf("[error]: %v", err)
// 		return "", err
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

// 		"sub": jwt.MapClaims{
// 			"id":       admin.ID,
// 			"name":     admin.Name,
// 			"password": admin.Password,
// 		},
// 		"nbf": time.Now().Unix(),
// 		"exp": time.Now().Add(tokenDuration).Unix(),
// 	})

// 	secret := []byte(os.Getenv("TOKEN_SECRET"))
// 	tokenString, err := token.SignedString(secret)
// 	if err != nil {
// 		log.Printf("[error]: %v", err)
// 		return "", err
// 	}
// 	return tokenString, nil
// }
