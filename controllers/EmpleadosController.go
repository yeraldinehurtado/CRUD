package controllers

import (
	"gomysql/models"
	"net/http"
	"strconv"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("views/*"))

func Inicio(w http.ResponseWriter, r *http.Request) { // request es lo que estan enviando
	lista := models.ListEmp()
	renderTemplate(w, "inicio", lista)
	//plantillas.ExecuteTemplate(w, "inicio", lista)

}

func Crear(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "crear", nil)
	//plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // si hay datos post, vamos a reseccionar esos datos

		usuario := r.FormValue("usuario")
		nombre := r.FormValue("nombre") // en id dentro del formulario html le puse "nombre"
		correo := r.FormValue("correo")
		password := r.FormValue("password")

		models.CreateEmpleado(usuario, nombre, correo, password)
		http.Redirect(w, r, "/", http.StatusFound)

	}
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)

	borrar := models.GetEmp(idInt)
	borrar.Delete()

	http.Redirect(w, r, "/", http.StatusFound)

}

func handlerError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	//errorTemplate.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	err := plantillas.ExecuteTemplate(w, name, data)

	if err != nil {
		handlerError(w, http.StatusInternalServerError)
	}
}
