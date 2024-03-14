package usuarios

import (
	"auth/graph/model"
	"database/sql"
	"errors"
)

func parseRow(row *sql.Row, t *model.Usuario) error {
	return row.Scan(
		&t.ID,
		&t.Nombres,
		&t.Apellido1,
		&t.Apellido2,
		&t.Documento,
		&t.Celular,
		&t.Correo,
		&t.Sexo,
		&t.Direccion,
		&t.Estado,
		&t.Username,
		&t.FechaRegistro,
		&t.FechaUpdate,
	)
}

func parseRows(rows *sql.Rows, t *model.Usuario) error {
	return rows.Scan(
		&t.ID,
		&t.Nombres,
		&t.Apellido1,
		&t.Apellido2,
		&t.Documento,
		&t.Celular,
		&t.Correo,
		&t.Sexo,
		&t.Direccion,
		&t.Estado,
		&t.Username,
		&t.FechaRegistro,
		&t.FechaUpdate,
	)
}

func permisos_obligatorios(roles, permisosueltos []string) error {
	if len(roles) == 0 && len(permisosueltos) == 0 {
		return errors.New("selecciona al menos un rol o un permiso")
	}
	return nil
}
