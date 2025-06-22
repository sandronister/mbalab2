package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sandronister/mba-lab2/cep-service/internal/dto"
	"github.com/sandronister/mba-lab2/cep-service/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type LocaleHandler struct {
	localeFinder usecase.Finder
}

func NewLocaleHandler(localeFinder usecase.Finder) *LocaleHandler {
	return &LocaleHandler{
		localeFinder: localeFinder,
	}
}

func (h *LocaleHandler) Handle(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	tracer := otel.Tracer("cep-service")
	_, span := tracer.Start(ctx, "locale-handler")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")

	var input dto.LocaleInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(msg)
		return
	}

	if len(input.Cep) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "invalid zipcode",
		})
		return
	}

	output, err := h.localeFinder.Execute(ctx, input.Cep)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if output.Localidade == "" {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusNotFound,
			Message:    "can not find zipcode",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("%+v", output)
	_ = json.NewEncoder(w).Encode(output)
}
