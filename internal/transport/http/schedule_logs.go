package http

import (
	"detaskify/internal/scheduler"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// LogExecution handles logging a new execution
func (h *Handler) LogExecution(w http.ResponseWriter, r *http.Request) {
	var log scheduler.ScheduleExecutionLog
	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Logs.LogExecution(r.Context(), &log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(log)
	if err != nil {
		return
	}
}

// GetExecutionLog retrieves an execution log by its ID
func (h *Handler) GetExecutionLog(w http.ResponseWriter, r *http.Request) {
	logIDStr := r.URL.Query().Get("id")
	logID, err := strconv.ParseUint(logIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid log ID", http.StatusBadRequest)
		return
	}

	log, err := h.Logs.GetExecutionLog(r.Context(), uint(logID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(log)
	if err != nil {
		return
	}
}

// UpdateExecutionLog updates an existing execution log
func (h *Handler) UpdateExecutionLog(w http.ResponseWriter, r *http.Request) {
	logIDStr := r.URL.Query().Get("id")
	logID, err := strconv.ParseUint(logIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid log ID", http.StatusBadRequest)
		return
	}

	var updatedLog scheduler.ScheduleExecutionLog
	err = json.NewDecoder(r.Body).Decode(&updatedLog)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err = h.Logs.UpdateExecutionLog(r.Context(), uint(logID), &updatedLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedLog)
	if err != nil {
		return
	}
}

// DeleteExecutionLog deletes an execution log by its ID
func (h *Handler) DeleteExecutionLog(w http.ResponseWriter, r *http.Request) {
	logIDStr := r.URL.Query().Get("id")
	logID, err := strconv.ParseUint(logIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid log ID", http.StatusBadRequest)
		return
	}

	err = h.Logs.DeleteExecutionLog(r.Context(), uint(logID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Log deleted successfully"})
	if err != nil {
		return
	}
}

// ListExecutionLogsByScheduleID retrieves all execution logs for a specific schedule ID
func (h *Handler) ListExecutionLogsByScheduleID(w http.ResponseWriter, r *http.Request) {
	scheduleIDStr := r.URL.Query().Get("scheduleID")
	scheduleID, err := uuid.Parse(scheduleIDStr)
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

	logs, err := h.Logs.ListExecutionLogsByScheduleID(r.Context(), scheduleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(logs)
	if err != nil {
		return
	}
}
