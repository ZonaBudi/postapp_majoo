package authservice_test

import (
	"context"
	"net/http"
	"postapp/internal/domain"
	"postapp/internal/port/mock"
	"postapp/internal/service/authservice"
	"postapp/pkg/model"
	"postapp/pkg/response"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type AuthTestSuite struct {
	suite.Suite
	Log                *zap.Logger
	User               *domain.User
	MockUserRepository *mock.MockUserRepository
}

func (suite *AuthTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	suite.Log, _ = zap.NewProduction()
	var id uint64 = 1
	suite.User = &domain.User{
		Base: model.Base{
			ID: &id,
		},
		Name:     "test",
		UserName: "test",
		Password: "e00cf25ad42683b3df678c61f42c6bda",
	}

	suite.MockUserRepository = mock.NewMockUserRepository(mockCtrl)
}

func (suite *AuthTestSuite) TestLogin() {
	cases := []struct {
		name       string
		wantResult *domain.Login
		err        error
		params     *domain.LoginRequest
	}{
		{
			name:       "failed_get_user",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)),
			params:     &domain.LoginRequest{},
		},
		{
			name:       "user_not_found",
			wantResult: nil,
			err:        response.NewError(http.StatusNotFound, response.WithMessageError(authservice.ErrUserNotFound)),
			params: &domain.LoginRequest{
				Username: "test",
				Password: "1234",
			},
		},
		{
			name:       "password_not_match",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(authservice.ErrCredentialNotMatch)),
			params: &domain.LoginRequest{
				Username: "test",
				Password: "1234",
			},
		},
		{
			name:       "failed_generate_token",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)),
			params: &domain.LoginRequest{
				Username: "test",
				Password: "admin1",
			},
		},
		{
			name: "login_success",
			wantResult: &domain.Login{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjF9.XXrW3Gv3zDbd9ydgTCb0A-rIQ0lSdi5QV_jpCWW22ew",
			},
			err: nil,
			params: &domain.LoginRequest{
				Username: "test",
				Password: "admin1",
			},
		},
	}
	for _, tCase := range cases {
		switch tCase.name {
		case "failed_get_user":
			suite.Run(tCase.name, func() {
				suite.MockUserRepository.EXPECT().
					FindOneByUsername(gomock.Any(), gomock.Any()).
					Return(nil, tCase.err).
					Times(1)
				authService := authservice.NewAuthService(suite.Log, suite.MockUserRepository)
				_, err := authService.Login(context.Background(), tCase.params)
				assert.Error(suite.T(), err)
			})
		case "user_not_found":
			suite.Run(tCase.name, func() {
				suite.MockUserRepository.EXPECT().
					FindOneByUsername(gomock.Any(), gomock.Any()).
					Return(nil, nil).
					Times(1)
				authService := authservice.NewAuthService(suite.Log, suite.MockUserRepository)
				_, err := authService.Login(context.Background(), tCase.params)
				assert.Error(suite.T(), err)
			})
		case "password_not_match":
			suite.Run(tCase.name, func() {
				suite.MockUserRepository.EXPECT().
					FindOneByUsername(gomock.Any(), gomock.Any()).
					Return(suite.User, nil).
					Times(1)
				authService := authservice.NewAuthService(suite.Log, suite.MockUserRepository)
				_, err := authService.Login(context.Background(), tCase.params)
				assert.Error(suite.T(), err)
			})
		case "failed_generate_token":
			suite.Run(tCase.name, func() {
				suite.User.ID = nil
				suite.MockUserRepository.EXPECT().
					FindOneByUsername(gomock.Any(), gomock.Any()).
					Return(suite.User, nil).
					Times(1)
				authService := authservice.NewAuthService(suite.Log, suite.MockUserRepository)
				_, err := authService.Login(context.Background(), tCase.params)
				assert.Error(suite.T(), err)
			})
		case "login_success":
			suite.Run(tCase.name, func() {
				var id uint64 = 1
				suite.User.ID = &id
				suite.MockUserRepository.EXPECT().
					FindOneByUsername(gomock.Any(), gomock.Any()).
					Return(suite.User, nil).
					Times(1)
				authService := authservice.NewAuthService(suite.Log, suite.MockUserRepository)
				_, err := authService.Login(context.Background(), tCase.params)
				assert.NoError(suite.T(), err)
			})
		}

	}
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
