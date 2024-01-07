package http

import (
	"detaskify/internal/scheduler"
	"encoding/json"
	"net/http"
	"strconv"
)

// CreateScheduleReminder handles the creation of a new schedule reminder
func (h *Handler) CreateScheduleReminder(w http.ResponseWriter, r *http.Request) {
	var reminder scheduler.ScheduleReminder
	err := json.NewDecoder(r.Body).Decode(&reminder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.SchedulerReminders.CreateScheduleReminder(r.Context(), &reminder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(reminder)
	if err != nil {
		return
	}
}

// GetScheduleReminder retrieves a schedule reminder by its ID
func (h *Handler) GetScheduleReminder(w http.ResponseWriter, r *http.Request) {
	reminderIDStr := r.URL.Query().Get("id")
	reminderID, err := strconv.ParseUint(reminderIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	reminder, err := h.SchedulerReminders.GetScheduleReminder(r.Context(), uint(reminderID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(reminder)
	if err != nil {
		return
	}
}

// UpdateScheduleReminder updates an existing schedule reminder
func (h *Handler) UpdateScheduleReminder(w http.ResponseWriter, r *http.Request) {
	reminderIDStr := r.URL.Query().Get("id")
	reminderID, err := strconv.ParseUint(reminderIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	var updatedReminder scheduler.ScheduleReminder
	err = json.NewDecoder(r.Body).Decode(&updatedReminder)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err = h.SchedulerReminders.UpdateScheduleReminder(r.Context(), uint(reminderID), &updatedReminder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedReminder)
	if err != nil {
		return
	}
}

// DeleteScheduleReminder deletes a schedule reminder by its ID
func (h *Handler) DeleteScheduleReminder(w http.ResponseWriter, r *http.Request) {
	reminderIDStr := r.URL.Query().Get("id")
	reminderID, err := strconv.ParseUint(reminderIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	err = h.SchedulerReminders.DeleteScheduleReminder(r.Context(), uint(reminderID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Reminder deleted successfully"})
	if err != nil {
		return
	}
}

// ListRemindersInAscendingOrder lists all schedule reminders in ascending order of reminder time
func (h *Handler) ListRemindersInAscendingOrder(w http.ResponseWriter, r *http.Request) {
	reminders, err := h.SchedulerReminders.ListRemindersInAscendingOrder(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(reminders)
	if err != nil {
		return
	}
}
