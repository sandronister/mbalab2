package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock IHttpService implementation for constructor test
type mockHttpService struct{}

func (m *mockHttpService) Do(url string) ([]byte, error) {
	return nil, nil
}

func TestNewWeatherByCep(t *testing.T) {
	mockHttp := &mockHttpService{}
	cepPath := "http://example.com/%s"
	usecase := NewWeatherByCep(cepPath, mockHttp)

	assert.NotNil(t, usecase, "NewWeatherByCep should not return nil")
	assert.Equal(t, cepPath, usecase.cepPath, "cepPath should be set correctly")
	assert.Equal(t, mockHttp, usecase.http, "http service should be set correctly")
}
