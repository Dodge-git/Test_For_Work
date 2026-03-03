package transport

import (
	"encoding/json"
	"net/http"

	"github.com/Dodge-git/Test_For_Work/internal/service"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}

func handleServiceError(w http.ResponseWriter, err error) {
	switch err {

	case service.ErrDepartmentNotFound:
		writeError(w, http.StatusNotFound, err.Error())

	case service.ErrParentNotFound:
		writeError(w, http.StatusNotFound, err.Error())

	case service.ErrInvalidName:
		writeError(w, http.StatusBadRequest, err.Error())

	case service.ErrDuplicateName:
		writeError(w, http.StatusConflict, err.Error())

	case service.ErrInvalidParent:
		writeError(w, http.StatusBadRequest, err.Error())

	case service.ErrCycleDetected:
		writeError(w, http.StatusConflict, err.Error())

	case service.ErrReassignRequired:
		writeError(w, http.StatusBadRequest, err.Error())

	case service.ErrInvalidDeleteMode:
		writeError(w, http.StatusBadRequest, err.Error())

	case service.ErrDepthExceeded:
		writeError(w, http.StatusBadRequest, err.Error())
		
	case service.ErrInvalidFullName:
		writeError(w, http.StatusBadRequest, err.Error())

	case service.ErrInvalidPosition:
		writeError(w, http.StatusBadRequest, err.Error())

	default:
		writeError(w, http.StatusInternalServerError, "internal error")
	}
}
