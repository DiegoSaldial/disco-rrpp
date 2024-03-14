package permisos

import (
	"auth/graph/model"
	"database/sql"
)

func parse(rows *sql.Rows, t *model.ResponsePermisoMe) error {
	return rows.Scan(
		&t.Metodo,
		&t.Nombre,
		&t.Descripcion,
		&t.FechaRegistro,
		&t.FechaAsignado,
	)
}

func parseRows(rows *sql.Rows, t *model.Permiso) error {
	return rows.Scan(
		&t.Metodo,
		&t.Nombre,
		&t.Descripcion,
		&t.FechaRegistro,
	)
}
