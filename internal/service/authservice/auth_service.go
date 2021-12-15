package authservice

import (
	"context"
	"errors"
	"net/http"
	"postapp/internal/domain"
	"postapp/internal/port"
	"postapp/pkg/response"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrCredentialNotMatch = errors.New("credential not match")
)

type authService struct {
	log       *zap.Logger
	userMysql port.UserRepository
}

func NewAuthService(log *zap.Logger, userMysql port.UserRepository) port.AuthService {
	return &authService{
		log:       log,
		userMysql: userMysql,
	}
}

func (_instance *authService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.Login, error) {
	user, err := _instance.userMysql.FindOneByUsername(ctx, req.Username)
	if err != nil {
		_instance.log.Error("failed get user : %v", zap.Error(err))
		return nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault))
	}
	if user.IsEmpty() {
		return nil, response.NewError(http.StatusNotFound, response.WithMessageError(ErrUserNotFound))
	}
	if !user.ComparePassword(req.HashedPassword()) {
		return nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(ErrCredentialNotMatch))
	}
	token, err := user.GenerateTokenAccess(viper.GetString("server.secret_access"))
	if err != nil {
		_instance.log.Error("failed generate token access", zap.Error(err))
		return nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault))
	}
	return &domain.Login{
		Token: *token,
	}, nil
}
