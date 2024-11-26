package routes

import (
	"fmt"
	controller "github/1AdrianM/go_inventario/internal/Controller"
	model "github/1AdrianM/go_inventario/internal/Model"
	repository "github/1AdrianM/go_inventario/internal/Repository"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

func TodoListPage(w http.ResponseWriter, r *http.Request) {
	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)

	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "Protected", "TodoList", "todo-list.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	task, err := controller.GetAllTask()
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Error:%s", err), http.StatusInternalServerError)
	}
	type TaskView struct {
		Title     string
		Text      string
		Completed bool
		ID        uint
	}
	var taskView []TaskView
	for _, m := range task {
		taskView = append(taskView, TaskView{
			Title:     m.Title,
			Text:      m.Text,
			Completed: m.Completed,
			ID:        m.ID,
		})
	}
	fmt.Printf("Leyendo datos de Tarea : %v\n", taskView)
	// Renderizar la plantilla con los datos
	tmpl.Execute(w, struct {
		Task []TaskView
	}{
		Task: taskView})
}
func HandlerTaskCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Calling create Task on POST")
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form:%s", err), http.StatusBadRequest)
		return
	}
	Title := r.FormValue("Title")
	Text := r.FormValue("Text")
	fmt.Printf("Titulo: %s\n, Texto: %s\n", Title, Text)
	// Obtener el UserID del contexto
	userId, ok := r.Context().Value("UserID").(uint)
	if !ok {
		log.Println("Error: UserID no presente en el contexto")
		http.Error(w, "Usuario no autenticado", http.StatusUnauthorized)
		return
	}
	user, err := repository.GetUserById(userId)
	if err != nil {
		log.Println("Error: UsuarioID no encontrado")
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	task := model.Task{
		Title:  Title,
		Text:   Text,
		UserID: user.ID,
	}

	if err := repository.CreateTask(&task); err != nil {
		http.Error(w, fmt.Sprintf("Error creating task:%s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Print("Tarea Creada Exitosamente")
}
func EditModalComponent(w http.ResponseWriter, r *http.Request) {
	taskIdStr := mux.Vars(r)["id"]
	ID, strErr := strconv.ParseUint(taskIdStr, 10, 32)
	if strErr != nil {
		http.Error(w, fmt.Sprintf("Http Error 400 : %s", strErr), http.StatusBadRequest)
	}
	t, taskErr := repository.GetTaskById(uint(ID))
	if taskErr != nil {
		http.Error(w, "Error getting task", http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf(`
	
               <form hx-post="edit/task/%d"  hx-swap="beforeend" id="todo-list-form" class="needs-validation container" novalidate>
                
                <div class="form-row">
                  <div class="col">

                    <label for="" class="form-label">Titulo</label>
                     <input type="text" value="%s" class="form-control" id="Title" name="Title" placeholder="Ingrese su nueva tarea aqui" required>                 
                     <label for="" class="form-label">Descripcion</label>

                     <textarea  class="form-control" id="Text" name="Text" placeholder="Descripcion" required>
                     %s
					 </textarea> 
                  </div>
                </div>
      
                <button type="submit" class="icon-button">
                    <i class="fa-solid fa-square-plus"></i>                   
                </button>
              </form>
            
	
	`, t.ID, t.Title, t.Text)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}

func HandlerTaskUpdate(w http.ResponseWriter, r *http.Request) {

}

func HandlerTaskDelete(w http.ResponseWriter, r *http.Request) {

}
