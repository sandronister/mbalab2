package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sandronister/mbalab2/weather_service/internal/dto"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
)

type CepUsacase struct {
	cepPath string
	http    types.IHttpService
}

func NewWeatherByCep(cepPath string, http types.IHttpService) *CepUsacase {
	return &CepUsacase{
		cepPath: cepPath,
		http:    http,
	}
}

func (m *CepUsacase) Get(cep string) (*dto.CepResultDTO, *types.HttpError) {
	url := fmt.Sprintf(m.cepPath, cep)
	var cepResult dto.CepResultDTO

	resp, err := m.http.Do(url)
	if err != nil {
		return nil, types.NewHttpError(http.StatusUnprocessableEntity, "invalid zipcode")
	}

	if len(resp) == 0 {
		return nil, types.NewHttpError(http.StatusNotFound, "zipcode not found")
	}

	json.Unmarshal(resp, &cepResult)
	if cepResult.Cep == "" {
		return nil, types.NewHttpError(http.StatusNotFound, "zipcode not found")
	}
	return &cepResult, nil
}
