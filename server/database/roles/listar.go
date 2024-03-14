package roles

import (
	"auth/database/permisos"
	"auth/graph/model"
	"database/sql"
	"errors"
)

func GetRolesByUsuario(db *sql.DB, userid string, show_permisos bool) ([]*model.ResponseRolMe, error) {
	sql := `
	select r.nombre, r.descripcion, r.jerarquia, r.fecha_registro,ru.usuario_id, ru.fecha_registro as fecha_asignado 
	from roles r
	left join rol_usuario ru on ru.rol = r.nombre
	where ru.usuario_id = ?;
	`
	rows, err := db.Query(sql, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*model.ResponseRolMe{}

	for rows.Next() {
		r := model.ResponseRolMe{}
		er := parse(rows, &r)
		if er != nil {
			return nil, er
		}
		if show_permisos {
			r.Permisos, er = permisos.GetPermisosByRol(db, r.Nombre)
			if er != nil {
				return nil, er
			}
		}
		roles = append(roles, &r)
	}

	return roles, nil
}

func GetRolById(db *sql.DB, rol string) (*model.Rol, error) {
	sq := "select nombre,descripcion,jerarquia,fecha_registro from roles where nombre = ?"
	row := db.QueryRow(sq, rol)
	r := model.Rol{}
	err := parseRow(row, &r)
	if err == sql.ErrNoRows {
		return nil, errors.New("rol no existente")
	}
	return &r, nil
}

func GetRoles(db *sql.DB, show_permisos bool) ([]*model.ResponseRolCreate, error) {
	sql := `select nombre,descripcion,jerarquia,fecha_registro from roles order by jerarquia asc`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roles := []*model.ResponseRolCreate{}

	for rows.Next() {
		r := model.ResponseRolCreate{}
		er := parseRes(rows, &r)
		if er != nil {
			return nil, er
		}
		if show_permisos {
			r.Permisos, er = permisos.GetPermisosByRol(db, r.Nombre)
			if er != nil {
				return nil, er
			}
		}

		roles = append(roles, &r)
	}

	return roles, nil
}

func GetRolById2(db *sql.DB, rol string, show_permisos bool) (*model.ResponseRolCreate, error) {
	sq := `
	select r.nombre, r.descripcion, r.jerarquia, r.fecha_registro
	from roles r 
	where r.nombre  = ?;
	`
	row := db.QueryRow(sq, rol)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.New("no se encontro el registo del rol")
	}

	ro, err := GetRolById(db, rol)
	if err != nil {
		return nil, err
	}

	r := model.ResponseRolCreate{}
	r.Nombre = ro.Nombre
	r.Descripcion = ro.Descripcion
	r.Jerarquia = ro.Jerarquia
	r.FechaRegistro = ro.FechaRegistro
	r.Permisos, err = permisos.GetPermisosByRol(db, r.Nombre)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
