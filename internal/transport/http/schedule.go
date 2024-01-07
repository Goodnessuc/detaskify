package http

import (
	"detaskify/internal/scheduler"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// CreateSchedule handles the creation of a new schedule
func (h *Handler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule scheduler.Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Scheduler.CreateSchedule(r.Context(), &schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(schedule)
	if err != nil {
		return
	}
}

// GetSchedule retrieves a schedule by its ID
func (h *Handler) GetSchedule(w http.ResponseWriter, r *http.Request) {
	scheduleIDStr := r.URL.Query().Get("id")
	scheduleID, err := uuid.Parse(scheduleIDStr)
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

	schedule, err := h.Scheduler.GetSchedule(r.Context(), scheduleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(schedule)
	if err != nil {
		return
	}
}

// UpdateSchedule updates an existing schedule
func (h *Handler) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	scheduleIDStr := r.URL.Query().Get("id")
	scheduleID, err := uuid.Parse(scheduleIDStr)
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

	var updatedSchedule scheduler.Schedule
	err = json.NewDecoder(r.Body).Decode(&updatedSchedule)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err = h.Scheduler.UpdateSchedule(r.Context(), scheduleID, &updatedSchedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedSchedule)
	if err != nil {
		return
	}
}

// DeleteSchedule deletes a schedule by its ID
func (h *Handler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	scheduleIDStr := r.URL.Query().Get("id")
	scheduleID, err := uuid.Parse(scheduleIDStr)
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

	err = h.Scheduler.DeleteSchedule(r.Context(), scheduleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Schedule deleted successfully"})
	if err != nil {
		return
	}
}

// GetUserSchedules retrieves all schedules for a specific user
func (h *Handler) GetUserSchedules(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	schedules, err := h.Scheduler.GetUserSchedules(r.Context(), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schedules)
	if err != nil {
		return
	}
}

// ListSchedules lists all schedules
func (h *Handler) ListSchedules(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.Scheduler.ListSchedules(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schedules)
	if err != nil {
		return
	}
}
