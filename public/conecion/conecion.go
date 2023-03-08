package bd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() *sql.DB {
	conexion, err := sql.Open("mysql", "root:"+""+"@tcp(localhost:3306)/asistenciaalumnos")
	if err != nil {
		panic(err.Error())
	}
	return conexion
}
