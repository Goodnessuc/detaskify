package http

import (
	task0 "detaskify/internal/tasks"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// CreateTask handles the creation of a new task
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task task0.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Task.CreateTask(r.Context(), &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		return
	}
}

// GetTask retrieves a task by its ID
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	// Assuming you have a way to get the ID from the URL
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.Task.GetTask(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		return
	}
}

// DeleteTask deletes a task by its ID

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.Task.DeleteTask(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
	if err != nil {
		return
	}
}

// GetUserTasks retrieves all task0 for a specific user
func (h *Handler) GetUserTasks(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	tasks, err := h.Task.GetUserTasks(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}

// ListTasksByStatus lists task0 with a specific status
func (h *Handler) ListTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	tasks, err := h.Task.ListTasksByStatus(r.Context(), status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}

// SearchTasks handles the searching of task0 based on various criteria
func (h *Handler) SearchTasks(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	description := r.URL.Query().Get("description")
	priority := r.URL.Query().Get("priority")

	tasks, err := h.Task.SearchTasks(r.Context(), title, description, priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}

// AddAssigneeToTask adds an assignee to a task
func (h *Handler) AddAssigneeToTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("taskID")
	taskID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	assignee := r.URL.Query().Get("assignee")

	err = h.Task.AddAssigneeToTask(r.Context(), uint(taskID), assignee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Assignee added successfully"})
	if err != nil {
		return
	}
}

// RemoveAssigneeFromTask removes an assignee from a task
func (h *Handler) RemoveAssigneeFromTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("taskID")
	taskID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	assignee := r.URL.Query().Get("assignee")

	err = h.Task.RemoveAssigneeFromTask(r.Context(), uint(taskID), assignee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Assignee removed successfully"})
	if err != nil {
		return
	}
}

// AddTagToTask adds a tag to a task
func (h *Handler) AddTagToTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("taskID")
	taskID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	tag := r.URL.Query().Get("tag")

	err = h.Task.AddTagToTask(r.Context(), uint(taskID), tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Tag added successfully"})
	if err != nil {
		return
	}
}

// RemoveTagFromTask removes a tag from a task
func (h *Handler) RemoveTagFromTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("taskID")
	taskID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	tag := r.URL.Query().Get("tag")

	err = h.Task.RemoveTagFromTask(r.Context(), uint(taskID), tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Tag removed successfully"})
	if err != nil {
		return
	}
}

// ListTasksByPriority lists task0 based on their priority
func (h *Handler) ListTasksByPriority(w http.ResponseWriter, r *http.Request) {
	priority := r.URL.Query().Get("priority")

	tasks, err := h.Task.ListTasksByPriority(r.Context(), priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}

func (h *Handler) ListTasksForReminder(w http.ResponseWriter, r *http.Request) {
	// Parse start and end dates from URL
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	tasks, err := h.Task.ListTasksForReminder(r.Context(), start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateTask updates an existing task
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Decode the request body into a task object
	var task task0.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err = h.Task.UpdateTask(r.Context(), uint(id), &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ListOverdueTasks(w http.ResponseWriter, r *http.Request) {
	// Use the server's current date
	currentDate := time.Now()

	tasks, err := h.Task.ListOverdueTasks(r.Context(), currentDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ListTasksByDeadline(w http.ResponseWriter, r *http.Request) {
	// Parse start and end dates from URL
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	tasks, err := h.Task.ListTasksByDeadline(r.Context(), start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
