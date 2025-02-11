package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHandler struct {
	called bool
}

func (m *MockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.called = true
}

func TestCORSMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedCalled bool
	}{
		{
			name:           "OPTIONS request",
			method:         "OPTIONS",
			expectedCalled: false,
		},
		{
			name:           "POST request",
			method:         "POST",
			expectedCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := &MockHandler{}
			middleware := NewCORSMiddleware(mockHandler)

			req := httptest.NewRequest(tt.method, "/", nil)
			rr := httptest.NewRecorder()

			middleware.ServeHTTP(rr, req)

			// Check CORS headers
			headers := rr.Header()
			if headers.Get("Access-Control-Allow-Origin") != "*" {
				t.Error("Expected Access-Control-Allow-Origin header to be *")
			}
			if headers.Get("Access-Control-Allow-Methods") != "POST, OPTIONS" {
				t.Error("Expected Access-Control-Allow-Methods header to be POST, OPTIONS")
			}
			if headers.Get("Access-Control-Allow-Headers") != "Content-Type" {
				t.Error("Expected Access-Control-Allow-Headers header to be Content-Type")
			}

			if mockHandler.called != tt.expectedCalled {
				t.Errorf("Expected handler to be called: %v, got: %v", tt.expectedCalled, mockHandler.called)
			}
		})
	}
}
