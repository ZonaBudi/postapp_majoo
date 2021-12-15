package midleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"postapp/internal/domain"
	"postapp/pkg/model"
	"postapp/pkg/response"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var (
	ErrUnauthorized = errors.New("you are not authorized")
)

type Key string

func Authenticate() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			if len(authHeader) != 2 {
				response.Error(w, response.NewError(http.StatusUnauthorized, response.WithMessageError(ErrUnauthorized)))
				return
			}
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(viper.GetString("server.secret_access")), nil
			})
			if err != nil {
				response.Error(w, response.NewError(http.StatusUnauthorized, response.WithMessageError(ErrUnauthorized)))
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), Key("user"), claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
		return http.HandlerFunc(fn)
	}
}

func User(c context.Context) *domain.User {
	user := c.Value(Key("user")).(jwt.MapClaims)
	id := uint64(user["uid"].(float64))
	return &domain.User{
		Base: model.Base{
			ID: &id,
		},
	}
}
