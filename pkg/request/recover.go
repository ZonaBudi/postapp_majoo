package request

import (
	"log"
	"net/http"
	"postapp/pkg/response"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				log.Print(rvr)
				response.Error(w, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)))
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
