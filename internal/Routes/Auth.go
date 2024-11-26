package routes

import (
	"fmt"
	controller "github/1AdrianM/go_inventario/internal/Controller"
	model "github/1AdrianM/go_inventario/internal/Model"
	service "github/1AdrianM/go_inventario/internal/Service"
	"github/1AdrianM/go_inventario/internal/helper"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func SignupPage(w http.ResponseWriter, r *http.Request) {
	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)

	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "Auth", "signup.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
func SignupPostHandler(w http.ResponseWriter, r *http.Request) {
	helper.Info("Signup handler function called")

	r.ParseForm()
	name := r.FormValue("name")
	lastname := r.FormValue("lastname")

	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	if password != confirmPassword {
		// Return error as partial HTML for HTMX
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<div style="color: red;">Passwords do not match!</div>`)
		return
	}
	user := model.User{Name: name, Lastname: lastname, Email: email, Password: password}

	err := controller.SignUp(&user)
	if err != nil {
		// Return error as partial HTML for HTMX
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<div style="color: red;">%v</div>`, err.Error())

	}
	w.Header().Set("HX-Redirect", "/login")
	fmt.Println("Redirection sent")
}
func LoginPage(w http.ResponseWriter, r *http.Request) {
	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)
	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "Auth", "login.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	helper.Info("Login handler function called")

	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<div style="color: red;">invalid inputs</div>`)
		return
	}
	user, err := controller.LogIn(email, password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := service.GenerateJWT(user.Email, user.ID)
	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<div style="color: red;">Error generating token</div>`)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true, // Solo accesible desde el servidor
	})

	// Redirigir a la página de inicio
	w.Header().Set("HX-Redirect", "/dashboard")
	fmt.Println("Redirection sent")
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Invalidar la cookie estableciendo una fecha de expiración pasada
	helper.Info("Logout handler function called")

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour), // Fecha de expiración pasada
		HttpOnly: true,
		Secure:   false, // Asegúrate de usar esto en producción con HTTPS
		SameSite: http.SameSiteStrictMode,
	})

	// Redirigir al cliente o devolver una respuesta exitosa
	w.Header().Set("HX-Redirect", "/login")
	fmt.Println("Redirection sent")
}
