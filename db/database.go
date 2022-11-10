package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// generar url
//tenemos que tener: username, password@tcp(localhost:3306)/database

const url = "root:1234@tcp(localhost:3306)/sistema"

// guarda la conexion
var db *sql.DB // conection la guardamos en una variable global

// funcion para hacer la conexion
func Connect() {
	conection, err := sql.Open("mysql", url) // nombre del dirver y la ruta de la base de datos
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión exitosa")

	db = conection
}

// funcion para cerrar la conexión
func Close() {
	db.Close()
}

// verificar la conexión si sigue conectada o no
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// verificar si una tabla existe
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return rows.Next()
}

// crear una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}

}

//Borrar y reiniciar los registros de una tabla
//trucate table
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// polimorfismo de exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...) // argumentos indefinidos
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// polimorfismo de query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...) // argumentos indefinidos
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

// https://github.com/go-sql-driver/mysql/
// aqui esta todo, drivers...
