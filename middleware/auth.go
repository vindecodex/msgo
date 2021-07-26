package middleware

import (
	"net/http"

	"github.com/vindecodex/msgo/domain"
	"github.com/vindecodex/msgo/logger"
)

type AuthMiddleware struct {
	Repo domain.UserRepository
}

func (m AuthMiddleware) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if contains(url, PUBLIC_URL) {
			next.ServeHTTP(w, r)
		} else {
			reqToken := r.Header.Get("Authorization")
			if reqToken == "" {
				logger.Error("Unauthorized Access")
				writeResponse(w, http.StatusUnauthorized, "Unauthorized Access")
			} else {
				token := getToken(reqToken)
				err := m.Repo.Authorize(token, url)
				if err != nil {
					logger.Error(err.AsMessage())
					writeResponse(w, err.Code, err.AsResponse())
				} else {
					next.ServeHTTP(w, r)
				}

			}

		}

	})
}
