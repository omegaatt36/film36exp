package auth

import (
	"net/http"
	"strings"

	"film36exp/model"
	"film36exp/utility"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret")

func GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
	})

	return token.SignedString(secret)
}

// TokenMiddleware is a middleware for some authenticated func
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) == 0 {
			utility.ResponseWithJSON(w, http.StatusUnauthorized, utility.Response{Result: utility.ResFailed, Message: "missing Authorization header"})
			return
		}
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
		token, err := verifyToken(tokenStr)
		if err != nil {
			utility.ResponseWithJSON(w, http.StatusUnauthorized, utility.Response{Result: utility.ResFailed, Message: "not authorized"})
		}
		if !token.Valid {
			utility.ResponseWithJSON(w, http.StatusUnauthorized, utility.Response{Result: utility.ResFailed, Message: "not authorized"})
		} else {
			r.Header.Set("userName", token.Claims.(jwt.MapClaims)["username"].(string))
			next.ServeHTTP(w, r)
		}

	})
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, err
}
