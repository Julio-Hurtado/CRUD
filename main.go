package main

import (
	db "github.com/Julio-Hurtado/CRUD/db/configuraciones"
	repositorio "github.com/Julio-Hurtado/CRUD/db/repositorios"
)

func main() {
	db.AbrirConexion()
	especialidadRepo := repositorio.Especialidad{}
	especialidadRepo.Listar(db.DB)
	db.CerrarConexion()
}
