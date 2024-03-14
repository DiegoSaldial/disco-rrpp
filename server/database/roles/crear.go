package roles

import (
	"auth/database/permisos"
	"auth/graph/model"
	"database/sql"
	"fmt"
	"strings"
)

func Crear(db *sql.DB, input model.NewRol) (*model.ResponseRolCreate, error) {
	sql := `insert into roles(nombre,descripcion,jerarquia) values (?,?,?)`
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(sql, input.Nombre, input.Descripcion, input.Jerarquia)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	sql = "insert into rol_permiso(rol, metodo) values %s"
	places := make([]string, len(input.Permisos))
	args := make([]interface{}, len(input.Permisos)*2)

	for i, p := range input.Permisos {
		places[i] = "(?,?)"
		args[i*2] = input.Nombre
		args[i*2+1] = p
	}

	sql = fmt.Sprintf(sql, strings.Join(places, ", "))
	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	r, err := GetRolById(db, input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res := model.ResponseRolCreate{}
	res.Nombre = r.Nombre
	res.Descripcion = r.Descripcion
	res.Jerarquia = r.Jerarquia
	res.FechaRegistro = r.FechaRegistro
	res.Permisos, err = permisos.GetPermisosByRol(db, input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &res, nil
}
