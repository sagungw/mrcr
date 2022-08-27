package handler

import (
	"encoding/json"
	"net/http"
	"sagungw/mercari/core/service"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sagungw/gotrunks/log"
)

type Handler struct {
	UserService         service.UserService
	LoginHistoryService service.LoginHistoryService
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.From("handler", "Register").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := r.PostFormValue("user_email")
	password := r.PostFormValue("user_password")
	err = h.UserService.Register(r.Context(), email, password)
	if err != nil {
		log.From("handler", "Register").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.From("handler", "Register").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.From("handler", "Login").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := r.PostFormValue("user_email")
	password := r.PostFormValue("user_password")
	token, err := h.UserService.Login(r.Context(), email, password)
	if err != nil {
		log.From("handler", "Login").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(&LoginResponse{
		UserToken: token,
	})
	if err != nil {
		log.From("handler", "Login").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.From("handler", "Login").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoginHistory(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		log.From("handler", "LoginHistory").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loginHistory, err := h.LoginHistoryService.GetLoginHistory(r.Context(), userID, page)
	if err != nil {
		log.From("handler", "LoginHistory").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := []*LoginHistoryResponse{}
	for _, e := range loginHistory {
		response = append(response, &LoginHistoryResponse{Timestamp: time.Unix(int64(e.Time.T), 0).Format("2006-01-02T15:04:05+0700")})
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.From("handler", "LoginHistory").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.From("handler", "LoginHistory").Error(errors.Unwrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
