package httpadapter

import "net/http"

func APIKeyMiddleware(publicAPIKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")

			if apiKey != publicAPIKey {
				http.Error(w, "API key invalide", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
