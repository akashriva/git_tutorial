package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define your secret key here
var (
	secretKey = []byte(os.Getenv("JWT_SECRET"))
)

// GenerateToken generates a new JWT token
func GenerateToken(Id string, Email string, Role string) (string, error) {
	fmt.Printf("Type of str: %T\n", secretKey)
	claims := jwt.MapClaims{
		"user_id": Id,
		"email":   Email,
		"role":    Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
