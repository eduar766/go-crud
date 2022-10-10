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
	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}