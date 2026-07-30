package auth

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// UserIDFromJWT ≈ JWTAuthentication در DRF
// هدر Authorization را می‌گیرد، توکن را با secret چک می‌کند، user_id را برمی‌گرداند.
func UserIDFromJWT(authHeader, secret string) (string, error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("missing bearer token")
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	switch v := claims["user_id"].(type) {
	case float64:
		return strconv.FormatInt(int64(v), 10), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("user_id missing")
	}
}
