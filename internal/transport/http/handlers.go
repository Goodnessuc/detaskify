package http

import (
	"context"
	"detaskify/internal/scheduler"
	"detaskify/internal/tasks"
	"detaskify/internal/users"
	"detaskify/internal/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Handler struct {
	Router             *mux.Router
	Server             *http.Server
	Users              users.UserRepository
	TaskReminders      tasks.ReminderRepository
	Teams              users.TeamRepository
	Task               tasks.TaskRepository
	Validator          *utils.Validator
	OAuthService       *OAuthService ``
	TaskComments       tasks.TaskCommentRepository
	SchedulerLogs      scheduler.ExecutionLogRepository
	SchedulerReminders scheduler.ReminderRepository
	Scheduler          scheduler.ScheduleRepository
}

// NewHandler - returns a pointer to a Handler
func NewHandler(users users.UserRepository, teams users.TeamRepository, task tasks.TaskRepository, taskReminder tasks.ReminderRepository, taskComments tasks.TaskCommentRepository, logs scheduler.ExecutionLogRepository, scheduleReminders scheduler.ReminderRepository, scheduler scheduler.ScheduleRepository) *Handler {
	log.Println("setting up our handler")
	h := &Handler{
		Users:              users,
		Task:               task,
		TaskReminders:      taskReminder,
		Teams:              teams,
		TaskComments:       taskComments,
		SchedulerLogs:      logs,
		SchedulerReminders: scheduleReminders,
		Scheduler:          scheduler,
		Validator:          utils.NewValidator(),
		OAuthService:       NewOAuthService(),
	}

	h.Router = mux.NewRouter()

	h.mapRoutes()

	h.Router.Use(JSONMiddleware)

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8080", // Good practice to set timeouts to avoid Slow-loris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}

	return h
}

// Serve - gracefully serves our new setup handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c
	// CreateAccount a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := h.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down gracefully")
	return nil
}

func (h *Handler) AliveCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode("Hey: I"); err != nil {
		panic(err)
	}
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", h.AliveCheck).Methods("GET")

	// router declaration for version 1 of the API
	apiV1Router := h.Router.PathPrefix("/api/v1").Subrouter()

	// sub-routers of the V1 sub-router
	authRouter := apiV1Router.PathPrefix("/auth").Subrouter()
	commentRouter := apiV1Router.PathPrefix("/comments").Subrouter()
	oauthRouter := authRouter.PathPrefix("/oauth").Subrouter()
	githubRouter := oauthRouter.PathPrefix("/github").Subrouter()
	gitlabRouter := oauthRouter.PathPrefix("/gitlab").Subrouter()
	tasksRouter := apiV1Router.PathPrefix("/tasks").Subrouter()
	remindersRouter := tasksRouter.PathPrefix("/reminders").Subrouter()
	usersRouter := apiV1Router.PathPrefix("/users").Subrouter()
	teamsRouter := apiV1Router.PathPrefix("/teams").Subrouter()

	// task comments route declarations
	commentRouter.HandleFunc("/create", h.CreateComment).Methods("POST")
	commentRouter.HandleFunc("/", h.GetComment).Methods("GET")
	commentRouter.HandleFunc("/update", h.UpdateComment).Methods("PATCH")
	commentRouter.HandleFunc("/delete", h.DeleteComment).Methods("DELETE")
	commentRouter.HandleFunc("/listByTask", h.ListCommentsByTaskID).Methods("GET")

	// github oauth route declarations
	githubRouter.HandleFunc("/login", h.HandleGitHubLogin).Methods("GET")
	githubRouter.HandleFunc("/callback", h.HandleGitHubCallback).Methods("GET")

	// gitlab router declarations
	gitlabRouter.HandleFunc("/login", h.HandleGitLabLogin).Methods("GET")
	gitlabRouter.HandleFunc("/callback", h.HandleGitLabCallback).Methods("GET")

	// task router declarations
	tasksRouter.HandleFunc("/create", h.CreateTask).Methods("POST")
	tasksRouter.HandleFunc("/get", h.GetTask).Methods("GET")
	tasksRouter.HandleFunc("/delete", h.DeleteTask).Methods("DELETE")
	tasksRouter.HandleFunc("/getUserTasks", h.GetUserTasks).Methods("GET")
	tasksRouter.HandleFunc("/listByStatus", h.ListTasksByStatus).Methods("GET")
	tasksRouter.HandleFunc("/search", h.SearchTasks).Methods("GET")
	tasksRouter.HandleFunc("/addAssignee", h.AddAssigneeToTask).Methods("POST")
	tasksRouter.HandleFunc("/removeAssignee", h.RemoveAssigneeFromTask).Methods("DELETE")
	tasksRouter.HandleFunc("/addTag", h.AddTagToTask).Methods("POST")
	tasksRouter.HandleFunc("/removeTag", h.RemoveTagFromTask).Methods("DELETE")
	tasksRouter.HandleFunc("/listByPriority", h.ListTasksByPriority).Methods("GET")
	tasksRouter.HandleFunc("/listForReminder", h.ListTasksForReminder).Methods("GET")
	tasksRouter.HandleFunc("/update", h.UpdateTask).Methods("PATCH")
	tasksRouter.HandleFunc("/listOverdue", h.ListOverdueTasks).Methods("GET")
	tasksRouter.HandleFunc("/listByDeadline", h.ListTasksByDeadline).Methods("GET")

	// Routes for task reminders
	remindersRouter.HandleFunc("/create", h.CreateReminder).Methods("POST")
	remindersRouter.HandleFunc("/get", h.GetReminder).Methods("GET")
	remindersRouter.HandleFunc("/update", h.UpdateReminder).Methods("PATCH")
	remindersRouter.HandleFunc("/delete", h.DeleteReminder).Methods("DELETE")
	remindersRouter.HandleFunc("/listByTaskID", h.ListRemindersByTaskID).Methods("GET")

	// routes for user-specific functionality
	usersRouter.HandleFunc("/create", h.CreateUser).Methods("POST")
	usersRouter.HandleFunc("/getByUsername", h.getUserByUsername).Methods("GET")
	usersRouter.HandleFunc("/getByEmail", h.GetUserByEmail).Methods("GET")
	usersRouter.HandleFunc("/update", h.UpdateUser).Methods("PATCH")
	usersRouter.HandleFunc("/delete", h.DeleteUser).Methods("DELETE")
	usersRouter.HandleFunc("/updatePassword", h.UpdateUserPassword).Methods("PATCH")
	usersRouter.HandleFunc("/validateSignIn", h.ValidateSignInData).Methods("POST")

	// routes for team-specific functionality
	teamsRouter.HandleFunc("/create", h.CreateTeam).Methods("POST")
	teamsRouter.HandleFunc("/getByID", h.GetTeamByID).Methods("GET")
	teamsRouter.HandleFunc("/update", h.UpdateTeam).Methods("PATCH")
	teamsRouter.HandleFunc("/delete", h.DeleteTeam).Methods("DELETE")
	teamsRouter.HandleFunc("/addUser", h.AddUserToTeam).Methods("POST")
	teamsRouter.HandleFunc("/removeUser", h.RemoveUserFromTeam).Methods("DELETE")
}
