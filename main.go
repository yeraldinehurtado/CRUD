package main

import (
	"fmt"
	"gomysql/controllers"
	"gomysql/db"
	"gomysql/models"
	"log"
	"net/http"
)

func main() {
	// conexion
	db.Connect()
	//fmt.Println(db.ExistsTable("empleados"))
	db.CreateTable(models.UserSchema, "empleados")
	//db.TruncateTable("empleados")
	//db.Ping() // para comprobar la conexion a la bd
	//db.Close()

	TemplatesFile := http.FileServer(http.Dir("views"))

	// mux es una ruta asociada a un handler
	// mux obtiene la ruta como tambien el handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Inicio)
	mux.HandleFunc("/crear", controllers.Crear)
	mux.HandleFunc("/insertar", controllers.Insertar)
	mux.HandleFunc("/borrar", controllers.Borrar)
	mux.HandleFunc("/editar", controllers.Editar)
	mux.HandleFunc("/actualizar", controllers.Actualizar)

	mux.Handle("/views/", http.StripPrefix("/views/", TemplatesFile)) //! Mux of Templates file

	// Crear un servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux)) // ruta donde se va a ejecutar nuestro servidor

	log.Fatal(server.ListenAndServe())

}
