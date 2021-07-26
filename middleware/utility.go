package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
)

func getToken(token string) string {
	if strings.Contains(token, "Bearer ") {
		return strings.Split(token, "Bearer ")[1]
	}
	return " ;P "
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

func contains(value string, slice []string) bool {
	for _, v := range slice {
		if strings.Contains(v, value) {
			return true
		}
	}
	return false
}
