package http

import (
	"detaskify/internal/tasks"
	"encoding/json"
	"net/http"
	"strconv"
)

// CreateReminder handles the creation of a new reminder
func (h *Handler) CreateReminder(w http.ResponseWriter, r *http.Request) {
	var reminder tasks.TaskReminders
	err := json.NewDecoder(r.Body).Decode(&reminder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.TaskReminders.CreateReminder(r.Context(), &reminder)
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

// GetReminder retrieves a reminder by its ID
func (h *Handler) GetReminder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	reminder, err := h.TaskReminders.GetReminder(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(reminder)
	if err != nil {
		return
	}
}

// UpdateReminder updates an existing reminder
func (h *Handler) UpdateReminder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	var reminder tasks.TaskReminders
	err = json.NewDecoder(r.Body).Decode(&reminder)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err = h.TaskReminders.UpdateReminder(r.Context(), uint(id), &reminder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(reminder)
	if err != nil {
		return
	}
}

// DeleteReminder deletes a reminder by its ID
func (h *Handler) DeleteReminder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	err = h.TaskReminders.DeleteReminder(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "TaskReminders deleted successfully"})
	if err != nil {
		return
	}
}

// ListRemindersByTaskID retrieves all reminders for a specific task
func (h *Handler) ListRemindersByTaskID(w http.ResponseWriter, r *http.Request) {
	taskIDStr := r.URL.Query().Get("taskID")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	reminders, err := h.TaskReminders.ListRemindersByTaskID(r.Context(), uint(taskID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(reminders)
	if err != nil {
		return
	}
}
