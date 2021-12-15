package response

import (
	"encoding/json"
	"net/http"
)

type AppSuccessOption func(*AppSuccess)

type AppSuccess struct {
	Data *interface{} `json:"data,omitempty"`
	Meta *interface{} `json:"meta,omitempty"`
}

func SuccessMeta(meta interface{}) AppSuccessOption {
	return func(h *AppSuccess) {
		h.Meta = &meta
	}
}
func SuccessData(data interface{}) AppSuccessOption {
	return func(h *AppSuccess) {
		h.Data = &data
	}
}

func Success(w http.ResponseWriter, status int, option ...AppSuccessOption) {
	appSuccess := new(AppSuccess)
	// Loop through each option
	for _, opt := range option {
		// Call the option giving the instantiated
		opt(appSuccess)
	}
	w.WriteHeader(status)
	data, err := json.Marshal(appSuccess)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
