package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Dodge-git/Test_For_Work/internal/config"
	"github.com/Dodge-git/Test_For_Work/internal/database"
	"github.com/Dodge-git/Test_For_Work/internal/repository"
	"github.com/Dodge-git/Test_For_Work/internal/service"
	"github.com/Dodge-git/Test_For_Work/internal/transport"
)

func main() {

	cfg := config.Load()

	db := database.NewPostgres(cfg)

	depRepo := repository.NewDepartmentRepository(db)
	empRepo := repository.NewEmployeeRepository(db)
	txManager := repository.NewTransactionManager(db)

	depService := service.NewDepartmentService(depRepo, empRepo, txManager)
	empService := service.NewEmployeeService(depRepo, empRepo)

	depHandler := transport.NewDepartmentHandler(depService)
	empHandler := transport.NewEmployeeHandler(empService)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/departments/", func(w http.ResponseWriter, r *http.Request) {

		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		// /departments/id/employees
		if len(parts) == 3 && parts[2] == "employees" {
			if r.Method == http.MethodPost {
				empHandler.CreateEmployee(w, r)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// /departments/id
		if len(parts) == 2 {
			switch r.Method {
			case http.MethodGet:
				depHandler.GetDepartment(w, r)
			case http.MethodPatch:
				depHandler.UpdateDepartment(w, r)
			case http.MethodDelete:
				depHandler.DeleteDepartment(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})

	mux.HandleFunc("/departments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			depHandler.CreateDepartment(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		log.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}