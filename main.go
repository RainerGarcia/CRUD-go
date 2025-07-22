package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Persona struct {
	ID     int
	Nombre string
	Email  string
}

func conexion() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Password := ""
	NombreBD := "crud-go"

	conexion, err := sql.Open(Driver, Usuario+":"+Password+"@tcp(127.0.0.1:3306)/"+NombreBD)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err.Error())
		return nil
	}

	return conexion
}

var templates = template.Must(template.ParseGlob("templates/*"))

func index(w http.ResponseWriter, r *http.Request) {

	conexion := conexion()

	registros, err := conexion.Query("SELECT * FROM personas")
	if err != nil {
		fmt.Println("Error al consultar la base de datos:", err.Error())
		return
	}

	persona := Persona{}

	arregloPersona := []Persona{}

	for registros.Next() {
		var id int
		var nombre, email string

		err := registros.Scan(&id, &nombre, &email)
		if err != nil {
			fmt.Println("Error al leer los registros:", err.Error())
		}

		persona.ID = id
		persona.Nombre = nombre
		persona.Email = email

		arregloPersona = append(arregloPersona, persona)
	}

	templates.ExecuteTemplate(w, "index", arregloPersona)
}

func create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

func insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nombre := r.FormValue("name")
		email := r.FormValue("email")
		conexion := conexion()

		insertRegistro, err := conexion.Prepare("INSERT INTO personas(nombre, correo) VALUES(?, ?)")
		if err != nil {
			fmt.Println("Error al preparar la consulta:", err.Error())
			return
		}

		insertRegistro.Exec(nombre, email)
		http.Redirect(w, r, "/", 301)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	idEmpleado, _ := strconv.Atoi(id)

	conexion := conexion()

	deleteRegistro, err := conexion.Prepare("DELETE FROM personas WHERE id = ?")
	if err != nil {
		fmt.Println("Error al preparar la consulta de eliminación:", err.Error())
		return
	}

	deleteRegistro.Exec(idEmpleado)
	http.Redirect(w, r, "/", 301)
}

func update(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	idEmpleado, _ := strconv.Atoi(id)

	conexion := conexion()
	registro, err := conexion.Query("SELECT * FROM personas WHERE id = ?", idEmpleado)
	if err != nil {
		fmt.Println("Error al consultar el registro:", err.Error())
		return
	}

	persona := Persona{}

	for registro.Next() {
		var id int
		var nombre, email string

		err := registro.Scan(&id, &nombre, &email)
		if err != nil {
			fmt.Println("Error al leer el registro:", err.Error())
			return
		}

		persona.ID = id
		persona.Nombre = nombre
		persona.Email = email
	}

	templates.ExecuteTemplate(w, "update", persona)

}

func edit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("name")
		email := r.FormValue("email")

		idEmpleado, _ := strconv.Atoi(id)

		conexion := conexion()

		updateRegistro, err := conexion.Prepare("UPDATE personas SET nombre = ?, correo = ? WHERE id = ?")
		if err != nil {
			fmt.Println("Error al preparar la consulta de actualización:", err.Error())
			return
		}

		updateRegistro.Exec(nombre, email, idEmpleado)
		http.Redirect(w, r, "/", 301)
	}

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/update", update)
	http.HandleFunc("/edit", edit)

	fmt.Println("Server is running on http://localhost:8090")
	err := http.ListenAndServe("localhost:8090", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
