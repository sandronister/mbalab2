// go
package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeatherUseCase(t *testing.T) {
	mockHttp := &mockHttpService{}
	path := "http://example.com/%s"
	usecase := NewWeatherUseCase(path, mockHttp)

	assert.NotNil(t, usecase, "NewWeatherUseCase should not return nil")
	assert.Equal(t, path, usecase.path, "path should be set correctly")
	assert.Equal(t, mockHttp, usecase.http, "http service should be set correctly")
}
