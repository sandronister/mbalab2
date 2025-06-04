package usecase

import (
	"net/http"

	"githuc.com/sandronister/mbalab2/cep_service/internal/dto"
)

func (u *UCep) Valid(cep string) dto.HttpError {
	var errHttp dto.HttpError

	if cep == "" {
		errHttp.StatusCode = http.StatusUnprocessableEntity
		errHttp.Message = "CEP cannot be empty"
		return errHttp
	}

	if len(cep) != 8 {
		errHttp.StatusCode = http.StatusUnprocessableEntity
		errHttp.Message = "CEP must be 8 characters long"
		return errHttp
	}

	return errHttp
}
