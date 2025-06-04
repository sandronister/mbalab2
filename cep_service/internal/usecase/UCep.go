package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"githuc.com/sandronister/mbalab2/cep_service/internal/dto"
	"githuc.com/sandronister/mbalab2/cep_service/internal/infra/client_http/types"
)

type UCep struct {
	httpClient types.IHttpService
	urlService string
}

func NewUCep(httpClient types.IHttpService, urlService string) ICep {
	return &UCep{
		httpClient: httpClient,
		urlService: urlService,
	}
}

func (u *UCep) Do(ctx context.Context, cep string) (*dto.Result, *dto.HttpError) {

	err := u.Valid(cep)

	if err.StatusCode != 0 {
		return nil, dto.NewHttpError(err.StatusCode, err.Message)
	}

	url := fmt.Sprintf("%s/%s", u.urlService, cep)

	resp, errHttp := u.httpClient.Do(url)

	if errHttp != nil {
		return nil, dto.NewHttpError(http.StatusInternalServerError, fmt.Sprintf("Error fetching data from URL: %s", url))
	}

	var location dto.Result

	if err := json.Unmarshal(resp, &location); err != nil {
		return nil, dto.NewHttpError(http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling response body: %v", err))
	}

	return &location, nil

}
