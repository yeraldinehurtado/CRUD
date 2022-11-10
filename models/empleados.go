package models

import "gomysql/db"

type Empleados struct {
	Id       int64
	User     string
	Name     string
	Password string
	Email    string
}

// creo la lista de empleados
type Empleado []Empleados

const UserSchema string = `CREATE TABLE empleados (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(300) NOT NULL,
	name VARCHAR(300) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)` // unsigned es que acepta solo numeros positivos

// constructor

func Constructor(username, name, password, email string) *Empleados {
	emp := &Empleados{
		User:     username,
		Name:     name,
		Password: password,
		Email:    email,
	}
	return emp
}

// crear empleado e insertar en la bd
func CreateEmpleado(username, name, password, email string) *Empleados {
	emp := Constructor(username, name, password, email)
	emp.insert()
	return emp
}

// Insertar registro
func (e *Empleados) insert() {
	sql := "INSERT empleados SET username=?, name=? password=?, email=?"
	result, _ := db.Exec(sql, e.User, e.Name, e.Password, e.Email)
	e.Id, _ = result.LastInsertId()
}

// listar todos los registros
func ListEmp() Empleado { // devolver lista de empleads
	sql := "SELECT id, username, name, password, email FROM empleados"
	emp := Empleado{}
	rows, _ := db.Query(sql) // usar query para obtener todos los registros
	for rows.Next() {
		empl := Empleados{}
		rows.Scan(&empl.Id, &empl.User, &empl.Name, &empl.Password, &empl.Email) //recuperar id, username, name, password y email
		emp = append(emp, empl)
	} // recorrer el rows para obtener cada registro

	return emp
}

// obtener un registro

func GetEmp(id int) *Empleados {
	emp := Constructor("", "", "", "")

	sql := "SELECT id, username, name, password, email FROM empleados WHERE id=?"
	rows, _ := db.Query(sql, id) // buscar un registro por id
	for rows.Next() {
		rows.Scan(&emp.Id, &emp.User, &emp.Name, &emp.Password, &emp.Email) //recuperar id, username, name, password y email
	} // recorrer el rows para obtener cada registro
	return emp
}

// actualizar registro
func (e *Empleados) update(){
	sql := "UPDATE empleados SET username=?, name=? password=?, email=? WHERE id=?"
	db.Exec(sql, e.User, e.Name, e.Password, e.Email, e.Id)
}

//Guardar o editar registro
func (e *Empleados) Save(){
	if e.Id == 0 {
		e.insert()
	}else {
		e.update()
	}
}

//eliminar un registro
func (e *Empleados) Delete(){
	sql := "DELETE FROM empleados WHERE id=?"
	db.Exec(sql, e.Id)
}