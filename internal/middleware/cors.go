package middleware

import "net/http"

type CORSMiddleware struct {
	handler http.Handler
}

func NewCORSMiddleware(handler http.Handler) *CORSMiddleware {
	return &CORSMiddleware{handler: handler}
}

func (m *CORSMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	m.handler.ServeHTTP(w, r)
}
