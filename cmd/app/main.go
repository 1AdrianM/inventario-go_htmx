package main

import (
	db "github/1AdrianM/go_inventario/internal/Db"
	middleware "github/1AdrianM/go_inventario/internal/Middleware"
	model "github/1AdrianM/go_inventario/internal/Model"
	routes "github/1AdrianM/go_inventario/internal/Routes"
	"github/1AdrianM/go_inventario/internal/config"
	logger "github/1AdrianM/go_inventario/internal/helper"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
	logger.Info("Application started")
	//Conexcion a bases de datos
	db.DbConnection()
	db.Db.AutoMigrate(model.User{}, model.Mediciones{}, model.Task{})
	r := mux.NewRouter()
	//handler que sirve los archivos estaticos
	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		logger.Error("404 Not found")

	} // Construir la ruta correcta para la plantilla
	assetsPath := filepath.Join(baseDir, "internal", "Assets")
	fs := http.FileServer(http.Dir(assetsPath))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	//rutas protegidas
	Protected := r.PathPrefix("/dashboard").Subrouter()

	Protected.HandleFunc("", routes.DashboardPage).Methods("GET")
	Protected.HandleFunc("/todoList", routes.TodoListPage).Methods("GET")
	Protected.HandleFunc("/calendar", routes.CalendarPage).Methods("GET")

	//Protected.HandleFunc("/todoListComponent", routes.TodoListComponent).Methods("GET")
	Protected.HandleFunc("/navbar", routes.NavBarComponent).Methods("GET")

	Protected.HandleFunc("/tables", routes.MedicionTablePage).Methods("GET")
	Protected.HandleFunc("/LoadUserData/{id}", routes.LoadUserData).Methods("GET")

	Protected.HandleFunc("/create/mediciones", routes.HandleCreateMedicion).Methods("POST")
	Protected.HandleFunc("/create/excel", routes.ExportExcelHandler).Methods("GET")
	Protected.HandleFunc("/create/task", routes.HandlerTaskCreate).Methods("POST")
	Protected.HandleFunc("/edit/task/{id}", routes.EditModalComponent).Methods("GET")

	Protected.HandleFunc("/update/mediciones/{id}", routes.UpdateMedicionesHandler).Methods("PUT")
	Protected.HandleFunc("/delete/mediciones/{id}", routes.DeleteMedicionesHandler).Methods("DELETE")

	Protected.HandleFunc("/logout", routes.LogoutHandler).Methods("POST")

	Protected.Use(middleware.AuthMiddleware([]byte(config.JwtKey)))

	//rutas publicas
	r.HandleFunc("/", routes.LoginPage)
	r.HandleFunc("/signup", routes.SignupPage).Methods("GET")
	r.HandleFunc("/signup", routes.SignupPostHandler).Methods("POST")
	r.HandleFunc("/login", routes.LoginPage).Methods("GET")
	r.HandleFunc("/login", routes.LoginPostHandler).Methods("POST")
	//desplique de app en puerto 3000
	log.Fatal(http.ListenAndServe(":3000", r))
	logger.Info("running on PORT 3000")
}
