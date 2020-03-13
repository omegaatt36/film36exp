package auth

import (
	"fmt"
	"net/http"

	"film36exp/model"
	"film36exp/utility"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
	})

	return token.SignedString([]byte("secret"))
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			utility.ResponseWithJSON(w, http.StatusUnauthorized,
				utility.Response{Result: utility.ResFailed, Message: "not authorized"})
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					utility.ResponseWithJSON(w, http.StatusUnauthorized,
						utility.Response{Result: utility.ResFailed, Message: "not authorized"})
					return nil, fmt.Errorf("not authorization")
				}
				return []byte("secret"), nil
			})
			if !token.Valid {
				utility.ResponseWithJSON(w, http.StatusUnauthorized,
					utility.Response{Result: utility.ResFailed, Message: "not authorized"})
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
