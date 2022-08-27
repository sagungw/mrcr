package handler

import (
	"encoding/json"
	"net/http"
	"sagungw/mercari/core/indoarea"

	"github.com/pkg/errors"
	"github.com/sagungw/gotrunks/log"
)

type Handler struct {
	IndoareaService indoarea.IndoareaService
}

func (h *Handler) GetProvinces(w http.ResponseWriter, r *http.Request) {
	result, err := h.IndoareaService.GetProvinces(r.Context())
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(result.Provinces)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetCities(w http.ResponseWriter, r *http.Request) {
	result, err := h.IndoareaService.GetCities(r.Context())
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(result.Cities)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetDistricts(w http.ResponseWriter, r *http.Request) {
	result, err := h.IndoareaService.GetDistricts(r.Context())
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(result.Districts)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetSubDistricts(w http.ResponseWriter, r *http.Request) {
	result, err := h.IndoareaService.GetSubDistricts(r.Context())
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(result.SubDistricts)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.From("handler", "GetProvinces").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
