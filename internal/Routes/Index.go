package routes

import (
	repository "github/1AdrianM/go_inventario/internal/Repository"
	"github/1AdrianM/go_inventario/internal/helper"
	"html/template"
	"net/http"

	"path/filepath"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)

	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "index.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DashboardPage(w http.ResponseWriter, r *http.Request) {
	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)

	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "Protected", "dashboard.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userID := r.Context().Value("UserID").(uint)

	if userID == 0 {
		helper.Error("couldnt get user id from the context")
	}
	user, err := repository.GetUserById(userID)
	if err != nil {
		http.Error(w, "d", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, user)
}

/*
	func TodoListComponent(w http.ResponseWriter, r *http.Request) {
		baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
		if err != nil {
			http.Error(w, "404 Not found", http.StatusNotFound)

		} // Construir la ruta correcta para la plantilla
		viewPath := filepath.Join(baseDir, "internal", "Views", "Protected", "Components", "todo-list-component.html")
		tmpl := template.Must(template.ParseFiles(viewPath))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
*/

func NavBarComponent(w http.ResponseWriter, r *http.Request) {
	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)

	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "Protected", "Components", "nav-bar.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
func CalendarPage(w http.ResponseWriter, r *http.Request) {
	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)

	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "Protected", "Calendar", "calendar.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
