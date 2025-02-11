package handler

import (
	"bytes"
	"encoding/json"
	"gymshark/packcalculator/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockPackService struct{}

func (m *MockPackService) CalculatePacks(orderAmount int) model.PackResponse {
	return model.PackResponse{
		OrderAmount: orderAmount,
		Packs:       map[int]int{250: 1},
		TotalItems:  250,
	}
}

func TestPackHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		requestBody    interface{}
		expectedStatus int
	}{
		{
			name:           "Valid POST request",
			method:         http.MethodPost,
			requestBody:    model.OrderRequest{OrderAmount: 250},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			requestBody:    nil,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid request body",
			method:         http.MethodPost,
			requestBody:    "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Negative order amount",
			method:         http.MethodPost,
			requestBody:    model.OrderRequest{OrderAmount: -1},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewPackHandler(&MockPackService{})

			var body []byte
			if tt.requestBody != nil {
				body, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest(tt.method, "/calculate", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}
		})
	}
}
