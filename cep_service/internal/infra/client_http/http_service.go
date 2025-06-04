package clienthttp

import (
	"io"
	"net/http"

	"githuc.com/sandronister/mbalab2/cep_service/internal/dto"
	"githuc.com/sandronister/mbalab2/cep_service/internal/infra/client_http/types"
)

type httpService struct {
	requestService types.IRequestService
}

func NewHttpService(requestService types.IRequestService) types.IHttpService {
	return &httpService{
		requestService: requestService,
	}
}

func (h *httpService) Do(url string) ([]byte, error) {
	req, err := h.requestService.Get(url)
	if err != nil {
		return nil, dto.NewHttpError(http.StatusInternalServerError, "Error creating HTTP request: "+err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, dto.NewHttpError(http.StatusInternalServerError, "Error making HTTP request: "+err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, dto.NewHttpError(resp.StatusCode, "Error fetching data from URL: "+url)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, dto.NewHttpError(http.StatusInternalServerError, "Error reading response body: "+err.Error())
	}

	return body, nil

}
