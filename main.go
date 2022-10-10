package main

import (
	"net/http"

	"github.com/eduar766/go-crud/db"
	"github.com/eduar766/go-crud/models"
	"github.com/eduar766/go-crud/routes"
	"github.com/gorilla/mux"
)


func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Tasks{}, models.User{})
	db.DB.AutoMigrate( models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")


	http.ListenAndServe(":3000", r)
}