package delay

import (
	"net/http"
	"time"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	delay := r.URL.Query().Get("delay")

	timeout, err := time.ParseDuration(delay)
	if err != nil {
		http.Error(w, "invalid delay param", http.StatusBadRequest)
		return
	}

	if timeout > 10*time.Second {
		http.Error(w, "timeout should be less than 10 seconds", http.StatusBadRequest)
		return
	}

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-timer.C:
		w.WriteHeader(http.StatusOK)
		return
	case <-r.Context().Done():
		return
	}
}
