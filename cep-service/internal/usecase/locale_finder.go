package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sandronister/mba-lab2/cep-service/internal/dto"
	"go.opentelemetry.io/otel"
)

const urlViacepApi = "https://viacep.com.br/ws/%s/json/"

type LocaleFinder struct {
	httpClient *http.Client
}

func NewLocaleFinder(httpClient *http.Client) *LocaleFinder {
	return &LocaleFinder{httpClient: httpClient}
}

func (l *LocaleFinder) Execute(ctx context.Context, cep string) (*dto.LocaleOutput, error) {
	tracer := otel.Tracer("cep-service")
	_, span := tracer.Start(ctx, "locale-finder-usecase")
	defer span.End()

	requestURL := fmt.Sprintf(urlViacepApi, cep)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := l.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	_ = res.Body.Close()

	var output dto.LocaleOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
