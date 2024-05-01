package main

import (
	"log"
	"net/http"

	"github.com/nfaria01/go-restapi/db"
	"github.com/nfaria01/go-restapi/models"
	"github.com/nfaria01/go-restapi/routes"
	"github.com/gorilla/mux"	
	"github.com/joho/godotenv"
)

func main() {
	
	
		err := godotenv.Load(".env")
		if err != nil{
		 log.Fatalf("Error loading .env file: %s", err)
		}

	// database connection
	db.DBConnection()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	// Index route
	r.HandleFunc("/", routes.HomeHandler)

	s := r.PathPrefix("/api").Subrouter()

	// tasks routes
	s.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	s.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	s.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")

	// users routes
	s.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	s.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	s.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	s.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":4000", r)

}