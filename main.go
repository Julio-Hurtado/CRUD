package main

import (
	"fmt"

	db "github.com/Julio-Hurtado/CRUD/db/configuraciones"
	repositorio "github.com/Julio-Hurtado/CRUD/db/repositorios"
)

func main() {
	db.AbrirConexion()
	especialidadRepo := repositorio.Especialidad{}
	especialidades, err := especialidadRepo.Listar(db.DB)
	if err != nil {
		fmt.Println(err)
		panic("No se pudo realizar la peticion")
	}
	fmt.Println(especialidades)
	db.CerrarConexion()
}
