package main

import (
	"log"
	"net/http"
	"text/template"

	"sistema/public/conecion"
	"sistema/public/grupo"
)

var plantillas = template.Must(template.ParseGlob("public/plantillas/*"))

func main() {
	//funcion main de la aplicacion

	http.HandleFunc("/index", Inicio)
	http.HandleFunc("/panelMaestro", PanelMaestro)
	http.HandleFunc("/listaAlumnos",ListaAlumnos)
	log.Print("Servidor Corriendo..")
	http.ListenAndServe(":5600", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "index", nil)
}

func PanelMaestro(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conecion.ConexionDB()
	registros, err := conexionEstablecida.Query("select g.ID,g.clave,ma.nombre ,m.nombre, p.nombre,c.nombre from grupo g inner join materia m on g.fk_materia=m.ID inner join maestro ma on g.fk_maestro = ma.ID inner join periodo p on g.fk_periodo = p.ID inner join  carrera c on g.fk_carrera = c.ID")

	if err != nil {
		log.Print("error")
		panic(err.Error())
	}

	//se inicia una estructura
	grupoT := grupo.Grupo{}
	//se incializa un areglo de tipo estrucutra 
	ArregloGrupo := []grupo.Grupo{}
	for registros.Next() {
		var id int
		var clave, maestro, materia, periodo, carrera string

		err = registros.Scan(&id, &clave, &maestro, &materia, &periodo, &carrera)

		if err != nil {
			panic(err.Error())
		}

		grupoT.Id = id
		grupoT.Clave = clave
		grupoT.Maestro = maestro
		grupoT.Periodo = periodo
		grupoT.Materia = materia
		grupoT.Carrera = carrera

		ArregloGrupo = append(ArregloGrupo, grupoT)
	}
	plantillas.ExecuteTemplate(w, "panelMaestro", ArregloGrupo)
}


func ListaAlumnos(w http.ResponseWriter, r *http.Request)  {
	plantillas.ExecuteTemplate(w, "listaAlumnos", nil)
}