package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CheckAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		key := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(r.Header.Get("Authorization")[7:], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(key), nil
		})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expAt, ok := claims["exp"].(int64)
			if !ok || expAt < time.Now().Unix() {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			handler.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func UserFromHeader(r *http.Request) string {
	key := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(r.Header.Get("Authorization")[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("invalid token")
		}

		return []byte(key), nil
	})
	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expAt, ok := claims["exp"].(int64)
		if !ok || expAt < time.Now().Unix() {
			return ""
		}

		return claims["id"].(string)
	}

	return ""
}
