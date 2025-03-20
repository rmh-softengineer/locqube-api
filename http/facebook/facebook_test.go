package facebook

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rmh-softengineer/locqube/api/model"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		name          string
		accessToken   string
		mockHandler   http.Handler
		expectedToken *string
		expectedError error
	}{
		{
			name:        "successful return when is a valid Token",
			accessToken: "valid-token",
			mockHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(model.ValidationTokenResponse{
					Data: model.ValidationTokenData{
						IsValid: true,
						UserID:  "12345",
					},
				})
			}),
			expectedToken: stringPtr("mock-jwt-token-for-12345"),
			expectedError: nil,
		},
		{
			name:        "successful return when is a non-valid Token",
			accessToken: "invalid-token",
			mockHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(model.ValidationTokenResponse{
					Data: model.ValidationTokenData{
						IsValid: false,
					},
				})
			}),
			expectedToken: nil,
			expectedError: errors.New("invalid facebook token"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockClient := http.Client{Transport: MockRoundTripper(tt.mockHandler)}

			client := NewClient(mockClient, "appID", "appSecret")

			// When
			token, err := client.Login(tt.accessToken)

			// Then
			assert.Equal(t, tt.expectedToken, token)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestPost(t *testing.T) {
	tests := []struct {
		name          string
		post          model.Post
		accessToken   string
		mockHandler   http.Handler
		expectedError error
	}{
		{
			name: "successful return with a valid post",
			post: model.Post{
				Message: "Hello, World!",
			},
			accessToken: "valid-token",
			mockHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}),
			expectedError: nil,
		},
		{
			name: "successful return with a non-valid post",
			post: model.Post{
				Message: "",
			},
			accessToken: "valid-token",
			mockHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			}),
			expectedError: errors.New("failed to share a post"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockClient := http.Client{Transport: MockRoundTripper(tt.mockHandler)}

			client := NewClient(mockClient, "appID", "appSecret")

			// When
			err := client.Post(tt.post, tt.accessToken)

			// Then
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

// MockRoundTripper is a mock implementation of an HTTP RoundTripper
func MockRoundTripper(handler http.Handler) http.RoundTripper {
	return &mockTransport{handler: handler}
}

type mockTransport struct {
	handler http.Handler
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	m.handler.ServeHTTP(rec, req)

	response := rec.Result()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("error")
	}

	return response, nil
}

func stringPtr(s string) *string {
	return &s
}
