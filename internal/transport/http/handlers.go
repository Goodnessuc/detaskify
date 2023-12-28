package http

import (
	"context"
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
	Router    *mux.Router
	Server    *http.Server
	Users     users.UserRepository
	Validator *utils.Validator
}

// NewHandler - returns a pointer to a Handler
func NewHandler(users users.UserRepository) *Handler {
	log.Println("setting up our handler")
	h := &Handler{
		Users:     users,
		Validator: utils.NewValidator(),
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

}
