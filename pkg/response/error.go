package response

import (
	"encoding/json"
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	ErrBadRequest = errors.New("your request is in a bad format")
	ErrInput      = errors.New("data is invalid")
	ErrDefault    = errors.New("server busy, please try again later")
)

type AppErrorOption func(*AppError)

// APPError is the default error struct containing detailed information about the error
type AppError struct {
	// HTTP Status code to be set in response
	Status int `json:"-"`
	// Message is the error message that may be displayed to end users
	Message *string `json:"message,omitempty"`
	// Meta is the error detail detail data
	Meta *interface{} `json:"meta,omitempty"`
}

// New generates an application error
func NewError(status int, opts ...AppErrorOption) *AppError {
	err := new(AppError)
	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(err)
	}
	err.Status = status
	return err
}

// Error returns the error message.
func (e AppError) Error() string {
	return *e.Message
}

func WithMessageError(message error) AppErrorOption {
	return func(h *AppError) {
		err := message.Error()
		h.Message = &err
	}
}

func WithMetaError(meta interface{}) AppErrorOption {
	return func(h *AppError) {
		h.Meta = &meta
	}
}

// Response writes an error response to client
func Error(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *AppError:
		w.WriteHeader(e.Status)
		data, err := json.Marshal(err)
		if err != nil {
			panic(err)
		}
		w.Write(data)
		return
	case validation.Errors:
		errNew := NewError(http.StatusBadRequest, WithMessageError(ErrInput), WithMetaError(err))
		w.WriteHeader(http.StatusBadRequest)
		data, err := json.Marshal(errNew)
		if err != nil {
			panic(err)
		}
		w.Write(data)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
