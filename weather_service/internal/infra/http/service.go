package http

import (
	"io"
	"net/http"

	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
)

type service struct {
	request types.IRequestService
}

func NewService(request types.IRequestService) types.IHttpService {
	return &service{
		request: request,
	}
}

func (s *service) Do(url string) ([]byte, error) {
	req, err := s.request.Get(url)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, types.NewHttpError(res.StatusCode, "error on request")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
