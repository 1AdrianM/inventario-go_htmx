package routes

import (
	"fmt"
	controller "github/1AdrianM/go_inventario/internal/Controller"
	model "github/1AdrianM/go_inventario/internal/Model"
	repository "github/1AdrianM/go_inventario/internal/Repository"
	"github/1AdrianM/go_inventario/internal/helper"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/gorilla/mux"
)

func MedicionTablePage(w http.ResponseWriter, r *http.Request) {
	helper.Info("GET Mediciones Function called on GET")

	baseDir, err := filepath.Abs("../..") // Subir 3 directorios desde cmd/app
	if err != nil {
		http.Error(w, "404 Not found", http.StatusNotFound)

	} // Construir la ruta correcta para la plantilla
	viewPath := filepath.Join(baseDir, "internal", "Views", "Protected", "Tables", "tables.html")
	tmpl := template.Must(template.ParseFiles(viewPath))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	mediciones, err := controller.GetAllMediciones()

	if err != nil {
		http.Error(w, "could not find mediciones", http.StatusInternalServerError)
		return
	}
	var userId uint
	userIDFromContext := r.Context().Value("UserID").(uint)

	user, err := repository.GetUserById(userIDFromContext)
	if err != nil {
		http.Error(w, "Error", http.StatusNotFound)
		return
	}
	type UserView struct {
		Name     string
		Lastname string
	}

	userView := UserView{
		Name:     user.Name,
		Lastname: user.Lastname,
	}

	type MedicionesView struct {
		Fecha_mantenimiento string
		Motor_ubicacion     string
		Motor_tag           string
		Motor_KW            int
		Motor_RPM           int
		Bearing_d           float64
		Bearing_t           float64
		Bearing_a           float64

		ID uint // Explicit foreign key definition
	}
	var medicionesView []MedicionesView

	for _, m := range mediciones {
		fechaStr := ""
		userId = m.UserID
		if userIDFromContext == userId {

			if !m.Fecha_mantenimiento.IsZero() {
				fechaStr = m.Fecha_mantenimiento.Format("2006-01-02")
			}
			medicionesView = append(medicionesView, MedicionesView{
				Fecha_mantenimiento: fechaStr,
				Motor_ubicacion:     m.Motor_ubicacion,
				Motor_tag:           m.Motor_tag,
				Motor_KW:            m.Motor_KW,
				Motor_RPM:           m.Motor_RPM,
				Bearing_d:           m.Bearing_d,
				Bearing_t:           m.Bearing_t,
				Bearing_a:           m.Bearing_a,

				ID: m.ID,
			})
			fmt.Print(medicionesView)

		}
	}
	// Convert the fechaMantenimiento to a string
	fmt.Print(medicionesView)

	tmpl.Execute(w, struct {
		Mediciones []MedicionesView
		UserView
	}{Mediciones: medicionesView, UserView: userView})

}

func HandleCreateMedicion(w http.ResponseWriter, r *http.Request) {
	// Parsear el formulario
	helper.Info("CREATE Mediciones Function called on POST")

	if err := r.ParseForm(); err != nil {
		log.Printf("Error al parsear el formulario: %v", err)
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

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
	// Parsear los valores del formulario
	fechaMantenimiento := r.FormValue("fechaMantenimiento")
	motorUbicacion := r.FormValue("motorUbicacion")
	motorTag := r.FormValue("motorTag")
	motorKWStr := r.FormValue("motorKW")
	motorRPMStr := r.FormValue("motorRPM")
	bearingDStr := r.FormValue("bearingD")
	bearingTStr := r.FormValue("bearingT")
	bearingAStr := r.FormValue("bearingA")

	// Validar y parsear fecha
	fechaParsed, err := time.Parse("2006-01-02", fechaMantenimiento)
	if err != nil {
		log.Printf("Error al parsear la fecha: %v \n", err)
		http.Error(w, "Formato de fecha inválido", http.StatusBadRequest)
		return
	}

	// Validar y parsear valores numéricos
	motorKW, err := strconv.Atoi(motorKWStr)
	if err != nil {
		log.Printf("Error al convertir motorKW: %v \n", err)
		http.Error(w, "Formato inválido para motorKW", http.StatusBadRequest)
		return
	}

	motorRPM, err := strconv.Atoi(motorRPMStr)
	if err != nil {
		log.Printf("Error al convertir motorRPM: %v", err)
		http.Error(w, "Formato inválido para motorRPM", http.StatusBadRequest)
		return
	}

	bearingD, err := strconv.ParseFloat(bearingDStr, 64)
	if err != nil {
		log.Printf("Error al convertir bearingD: %v", err)
		http.Error(w, "Formato inválido para bearingD", http.StatusBadRequest)
		return
	}

	bearingT, err := strconv.ParseFloat(bearingTStr, 64)
	if err != nil {
		log.Printf("Error al convertir bearingT: %v", err)
		http.Error(w, "Formato inválido para bearingT", http.StatusBadRequest)
		return
	}
	bearingA, err := strconv.ParseFloat(bearingAStr, 64)
	if err != nil {
		log.Printf("Error al convertir bearingT: %v", err)
		http.Error(w, "Formato inválido para bearingT", http.StatusBadRequest)
		return
	}
	// Crear la medición
	medicion := model.Mediciones{
		Fecha_mantenimiento: fechaParsed,
		Motor_ubicacion:     motorUbicacion,
		Motor_tag:           motorTag,
		Motor_KW:            motorKW,
		Motor_RPM:           motorRPM,
		Bearing_d:           bearingD,
		Bearing_t:           bearingT,
		Bearing_a:           bearingA,

		UserID: user.ID,
	}

	// Guardar la medición en el repositorio
	if err := repository.CreateMediciones(&medicion); err != nil {
		log.Printf("Error al guardar la medición: %v \n", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.Header().Set("HX-Refresh", "#dataTable")
	w.WriteHeader(http.StatusCreated)
	log.Println("Medición creada exitosamente")
}

func ExportExcelHandler(w http.ResponseWriter, r *http.Request) {
	// Crear un archivo Excel
	f := excelize.NewFile()
	s, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Printf("Error al crear el archivo Excel: %v", err)
	}
	f.SetCellValue("Sheet1", "A1", "FechaMantenimiento")
	f.SetCellValue("Sheet1", "B1", "Ubicación")
	f.SetCellValue("Sheet1", "C1", "Tag")
	f.SetCellValue("Sheet1", "D1", "KW")
	f.SetCellValue("Sheet1", "E1", "RPM")
	f.SetCellValue("Sheet1", "F1", "BearingD")
	f.SetCellValue("Sheet1", "G1", "BearingT")
	f.SetCellValue("Sheet1", "H1", "BearingA")
	medicionData, err := repository.GetMediciones()
	if err != nil {
		log.Printf("Error al obtener las mediciones: %v", err)
	}
	for i, m := range medicionData {
		row := i + 2
		Fecha := m.Fecha_mantenimiento.Format("2006-01-02")
		KW := strconv.Itoa(m.Motor_KW)
		RPM := strconv.Itoa(m.Motor_RPM)
		D := strconv.FormatFloat(m.Bearing_d, 'f', 2, 64)
		T := strconv.FormatFloat(m.Bearing_t, 'f', 2, 64)
		A := strconv.FormatFloat(m.Bearing_a, 'f', 2, 64)
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), Fecha)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), m.Motor_ubicacion)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), m.Motor_tag)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), KW)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), RPM)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), D)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), T)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), A)
		fmt.Printf("bearing d:%.2f", m.Bearing_d)
		fmt.Printf("bearing TAG:%s", m.Motor_tag)
		fmt.Printf("bearing d stringified:%s", D)

	}
	f.SetActiveSheet(s)
	if err := f.SaveAs("mediciones.xlsx"); err != nil {
		log.Printf("Error al guardar el archivo Excel: %v", err)
	}
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=\"mediciones.xlsx\"")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	fmt.Println("Excel file generated successfully.")
	w.WriteHeader(http.StatusOK)

}

func DeleteMedicionesHandler(w http.ResponseWriter, r *http.Request) {
	helper.Info("DELETE Mediciones Function called on DELETE")

	pathVar := mux.Vars(r)
	var IdVar string = pathVar["id"]

	if IdVar == "" {
		helper.Error("id is empty")

	}
	Id, err := strconv.ParseUint(IdVar, 10, 32)
	if err != nil {
		log.Printf("Error al convertir IdVar: %v", err)
	}
	if err := controller.DeleteMedicion(uint(Id)); err != nil {
		http.Error(w, "Not able to delete Medicion", http.StatusInternalServerError)
	}
	w.Header().Set("HX-Refresh", "#dataTable")
	w.WriteHeader(http.StatusNoContent)
}

func LoadUserData(w http.ResponseWriter, r *http.Request) {
	// Assuming you have a User struct
	ID := mux.Vars(r)["id"]
	fmt.Printf(" User Id: %s", ID)
	mID, errorP := strconv.ParseUint(ID, 10, 32)
	if errorP != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := controller.GetMedicionById(uint(mID))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	responseHTML := fmt.Sprintf(`
    <div style="max-width: 600px; margin: 0 auto; padding: 10px;">
	          <form hx-put="update/mediciones/%d" hx-swap="afterSwap" id="Updateuser_form" class="needs-validation container" novalidate>

        <label for="fechaMantenimiento" class="" style="font-weight: bold; color: #333;">Fecha de Mantenimiento</label>
        <input type="date" class="form-control" id="fechaMantenimiento" value="%s" name="fechaMantenimiento" required style="width: 100%%; padding: 8px; margin: 5px 0; border: 1px solid #ccc;">
        
        <label for="motorUbicacion" class="" style="font-weight: bold; color: #333;">Motor Ubicación</label>
        <input type="text" class="form-control" id="motorUbicacion" name="motorUbicacion" value="%s" placeholder="Ubicación del motor" required style="width: 100%%; padding: 8px; margin: 5px 0; border: 1px solid #ccc;">
        
        <label for="motorTag" class="" style="font-weight: bold; color: #333;">Motor Tag</label>
        <input type="text" class="form-control" id="motorTag" name="motorTag" value="%s" placeholder="Etiqueta del motor" required style="width: 100%%; padding: 8px; margin: 5px 0; border: 1px solid #ccc;">
        
        <label for="motorKW" class="" style="font-weight: bold; color: #333;">Motor KW</label>
        <input type="number" class="form-control" id="motorKW" name="motorKW" value="%d" placeholder="Potencia del motor en KW" required style="width: 100%%; padding: 8px; margin: 5px 0; border: 1px solid #ccc;">
        
        <label for="motorRPM" class="" style="font-weight: bold; color: #333;">Motor RPM</label>
        <input type="number" class="form-control" id="motorRPM" name="motorRPM" value="%d" placeholder="Velocidad del motor en RPM" required style="width: 100%%; padding: 8px; margin: 5px 0; border: 1px solid #ccc;">
        
        <label for="bearingD" class="" style="font-weight: bold; color: #333;">Bearing D</label>
        <input type="number" step="0.01" class="form-control" id="bearingD" name="bearingD" value="%.2f" placeholder="Diámetro del rodamiento" required style="width: 100%%; padding: 8px; margin: 5px 0; border: 1px solid #ccc;">
        
        <label for="bearingT" class="" style="font-weight: bold; color: #333;">Bearing T</label>
        <input type="number" step="0.01" class="form-control" id="bearingT" value="%.2f" name="bearingT" placeholder="Espesor del rodamiento" required style="width: 100%%; padding: 8px; margin: 5px 0; border: 1px solid #ccc;">
                    <button type="submit" class="btn btn-primary mt-4">Actualizar</button>

		</form>
     </div>
`, mID, user.Fecha_mantenimiento, user.Motor_ubicacion, user.Motor_tag, user.Motor_KW, user.Motor_RPM, user.Bearing_d, user.Bearing_t)
	// Set Content-Type and write the response back to HTMX
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(responseHTML))
}

func UpdateMedicionesHandler(w http.ResponseWriter, r *http.Request) {
	helper.Info("UPDATE Mediciones Function called ")
	ID := mux.Vars(r)["id"]
	fmt.Printf(" User Id: %s", ID)
	mID, errorP := strconv.ParseUint(ID, 10, 32)
	if errorP != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("Error al parsear el formulario: %v", err)
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

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
	// Parsear los valores del formulario
	fechaMantenimiento := r.FormValue("fechaMantenimiento")
	motorUbicacion := r.FormValue("motorUbicacion")
	motorTag := r.FormValue("motorTag")
	motorKWStr := r.FormValue("motorKW")
	motorRPMStr := r.FormValue("motorRPM")
	bearingDStr := r.FormValue("bearingD")
	bearingTStr := r.FormValue("bearingT")

	// Validar y parsear fecha
	fechaParsed, err := time.Parse("2006-01-02", fechaMantenimiento)
	if err != nil {
		log.Printf("Error al parsear la fecha: %v \n", err)
		http.Error(w, "Formato de fecha inválido", http.StatusBadRequest)
		return
	}

	// Validar y parsear valores numéricos
	motorKW, err := strconv.Atoi(motorKWStr)
	if err != nil {
		log.Printf("Error al convertir motorKW: %v \n", err)
		http.Error(w, "Formato inválido para motorKW", http.StatusBadRequest)
		return
	}

	motorRPM, err := strconv.Atoi(motorRPMStr)
	if err != nil {
		log.Printf("Error al convertir motorRPM: %v", err)
		http.Error(w, "Formato inválido para motorRPM", http.StatusBadRequest)
		return
	}

	bearingD, err := strconv.ParseFloat(bearingDStr, 64)
	if err != nil {
		log.Printf("Error al convertir bearingD: %v", err)
		http.Error(w, "Formato inválido para bearingD", http.StatusBadRequest)
		return
	}

	bearingT, err := strconv.ParseFloat(bearingTStr, 64)
	if err != nil {
		log.Printf("Error al convertir bearingT: %v", err)
		http.Error(w, "Formato inválido para bearingT", http.StatusBadRequest)
		return
	}

	// Crear la medición
	medicion := model.Mediciones{
		Fecha_mantenimiento: fechaParsed,
		Motor_ubicacion:     motorUbicacion,
		Motor_tag:           motorTag,
		Motor_KW:            motorKW,
		Motor_RPM:           motorRPM,
		Bearing_d:           bearingD,
		Bearing_t:           bearingT,
		UserID:              user.ID,
	}

	// Guardar la medición en el repositorio
	if err := repository.UpdateMediciones(uint(mID), &medicion); err != nil {
		log.Printf("Error al guardar la medición: %v \n", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	log.Println("Medición creada exitosamente")
}
