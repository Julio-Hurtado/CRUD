package repositorio

import (
	"database/sql"

	model "github.com/Julio-Hurtado/CRUD/models"
)

type Especialidad struct{}

func (especialidadRepo *Especialidad) Listar(db *sql.DB) ([]model.Especialidad, error) {
	rows, err := db.Query("SELECT * FROM especialidad")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var especialidades []model.Especialidad
	for rows.Next() {
		especialidad := model.Especialidad{}
		rows.Scan(&especialidad.Id, &especialidad.Codigo, &especialidad.Nombre)
		especialidades = append(especialidades, especialidad)

	}
	return especialidades, nil
}
