package handler

import (
	"github.com/labstack/echo"
	"github.com/sandronister/mbalab2/cep_service/internal/dto"
	"github.com/sandronister/mbalab2/cep_service/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type CepHandler struct {
	usecase usecase.ICep
}

func NewCepHandler(usecase usecase.ICep) *CepHandler {
	return &CepHandler{
		usecase: usecase,
	}
}

func (h *CepHandler) GetWeather(c echo.Context) error {
	carrier := propagation.HeaderCarrier(c.Request().Header)
	ctx := c.Request().Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	tracer := otel.Tracer("cep_service")
	_, span := tracer.Start(ctx, "GetWeather")
	defer span.End()

	var cepDTO dto.Cep

	if err := c.Bind(cepDTO); err != nil {
		return c.JSON(400, dto.NewHttpError(400, "Invalid request body"))
	}

	result, err := h.usecase.Do(c.Request().Context(), cepDTO.Cep)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}

	return c.JSON(200, result)
}
