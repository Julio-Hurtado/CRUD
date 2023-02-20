package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func AbrirConexion() {
	conexion, err := sql.Open("postgres", ConnStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("conexion realizada con exito")
	DB = conexion
}

func CerrarConexion() {
	fmt.Println("conexion cerrada con exito")
	DB.Close()
}
