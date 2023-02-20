package repositorio

import (
	"database/sql"
	"fmt"
)

type Especialidad struct{}

func (especialidadRepo *Especialidad) Listar(db *sql.DB) {
	fmt.Println(db)
	rows, err := db.Query("SELECT * FROM medico")

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println(rows.Next())
}
