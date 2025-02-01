package pkg

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]string{"error": message}
	JSONResponse(w, statusCode, response)
}

func ParseJSONRequest(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return err
	}
	return nil
}
