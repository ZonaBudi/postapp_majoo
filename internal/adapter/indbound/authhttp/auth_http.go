package authhttp

import (
	"encoding/json"
	"net/http"
	"postapp/internal/domain"
	"postapp/internal/port"
	"postapp/pkg/request"
	"postapp/pkg/response"

	"github.com/go-chi/chi/v5"
)

type authHandlerHttp struct {
	app         *chi.Mux
	authService port.AuthService
}

func NewAuthHttp(app *chi.Mux, authService port.AuthService) {
	api := authHandlerHttp{
		app:         app,
		authService: authService,
	}
	api.app.With(request.RequestMiddleware).Post("/login", api.DoLogin)
}

// Login godoc
// @Summary      Login
// @Description  get credential
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user  body      domain.LoginRequest  true  "Login Account"
// @Success      200  {object}  response.AppSuccess
// @Failure      default {object}  response.AppError
// @Router       /login [post]
func (_instance *authHandlerHttp) DoLogin(w http.ResponseWriter, r *http.Request) {
	body := new(domain.LoginRequest)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		response.Error(w, response.NewError(http.StatusBadRequest, response.WithMessageError(response.ErrBadRequest)))
		return
	}
	if err := body.LoginValidation(); err != nil {
		response.Error(w, err)
		return
	}
	result, err := _instance.authService.Login(r.Context(), body)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, http.StatusOK, response.SuccessData(result))
}
