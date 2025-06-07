package http

import (
	"net/http"
	"testing"

	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
	"github.com/stretchr/testify/assert"
)

type mockRequestService struct{}

func (m *mockRequestService) Get(url string) (*http.Request, error) {
	return nil, nil
}

func TestNewService(t *testing.T) {
	mockReq := &mockRequestService{}
	svc := NewService(mockReq)
	assert.NotNil(t, svc, "NewService should not return nil")

	// Optionally check interface implementation
	_, ok := svc.(types.IHttpService)
	assert.True(t, ok, "Returned service should implement types.IHttpService")
}
