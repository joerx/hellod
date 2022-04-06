package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	bytes, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Header().Add("Content-type", "application/json")
	w.Write(bytes)
}

func OK(w http.ResponseWriter, payload interface{}) {
	JSON(w, http.StatusOK, payload)
}

func InternalServerError(w http.ResponseWriter, err error) {
	JSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("%s", err)})
}
