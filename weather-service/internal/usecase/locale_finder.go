package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sandronister/mba-lab2/weather-service/internal/dto"
	"go.opentelemetry.io/otel"
)

const urlLocaleApi = "http://%s%s/"

type LocaleFinder struct {
	httpClient *http.Client
}

func NewLocaleFinder(httpClient *http.Client) *LocaleFinder {
	return &LocaleFinder{httpClient: httpClient}
}

func (l *LocaleFinder) Execute(ctx context.Context, cep string) (interface{}, error) {
	tracer := otel.Tracer("weather-service")
	_, span := tracer.Start(ctx, "locale-finder-usecase")
	defer span.End()

	input := &dto.LocaleInput{
		Cep: cep,
	}
	inputJson, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	cepServiceServerName := os.Getenv("CEP_SERVICE_SERVER_NAME")
	if cepServiceServerName == "" {
		cepServiceServerName = "localhost"
	}
	cepServiceServerPort := os.Getenv("CEP_SERVICE_SERVER_PORT")
	if cepServiceServerPort == "" {
		cepServiceServerPort = ":8080"
	}

	requestURL := fmt.Sprintf(urlLocaleApi, cepServiceServerName, cepServiceServerPort)

	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(inputJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")
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
