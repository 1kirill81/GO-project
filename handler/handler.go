package handler

import (
	"encoding/json"
	"fmt"
	"go-project/service"
	"io"
	"net/http"
	"strings"
)

type TestHandler struct {
	service service.TestService
}

func NewTestHandler(service service.TestService) *TestHandler {
	return &TestHandler{service: service}
}

func (h *TestHandler) Test(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	msg := h.service.GetMessage()
	fmt.Fprint(w, msg)
}

func (h *TestHandler) DBTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	body := strings.TrimSpace(string(bodyBytes))
	if body == "" {
		http.Error(w, "request body is empty", http.StatusBadRequest)
		return
	}

	record, err := h.service.SaveDBTest(r.Context(), body)
	if err != nil {
		http.Error(w, "failed to save body to database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(record)
}
