package usuarios

import (
	"auth/database/permisos"
	"auth/database/roles"
	"auth/graph/model"
	"database/sql"
	"errors"
	"fmt"
)

func GetById(db *sql.DB, id string) (*model.Usuario, error) {
	sq := `
		select 
		u.id,
		u.nombres, 
		u.apellido1, 
		u.apellido2, 
		u.documento,
		u.celular,
		u.correo,
		u.sexo,
		u.direccion,
		u.estado,
		u.username,
		u.fecha_registro,
		u.fecha_update
		from usuarios u 
		where u.id = ?`

	row := db.QueryRow(sq, id)

	us := model.Usuario{}
	err := parseRow(row, &us)

	if err == sql.ErrNoRows {
		return nil, errors.New("usuario no encontrado por id")
	}
	if err != nil {
		return nil, err
	}

	return &us, nil
}

func GetByUserPass(db *sql.DB, user, pass string) (*model.Usuario, error) {
	sq := `
		select 
		u.id,
		u.nombres, 
		u.apellido1, 
		u.apellido2, 
		u.documento,
		u.celular,
		u.correo,
		u.sexo,
		u.direccion,
		u.estado,
		u.username,
		u.fecha_registro,
		u.fecha_update
		from usuarios u 
		where u.username = ? 
		and u.password = SHA2( ?, 256)`

	row := db.QueryRow(sq, user, pass)

	us := model.Usuario{}
	err := parseRow(row, &us)

	if err == sql.ErrNoRows {
		return nil, errors.New("usuario o clave incorrectos")
	}
	if err != nil {
		return nil, err
	}

	return &us, nil
}

func GetMe(db *sql.DB, input model.InputMe, userid string) (*model.ResponseMe, error) {
	us, err := GetById(db, userid)
	if err != nil {
		return nil, err
	}

	user := model.ResponseMe{}
	user.Usuario = us

	if input.ShowRoles {
		user.Roles, err = roles.GetRolesByUsuario(db, us.ID, input.ShowPermisos)
		if err != nil {
			return nil, errors.Join(err, errors.New("error al cargar roles al usuario"))
		}
	}

	if input.ShowPermisos {
		user.PermisosSueltos, err = permisos.GetPermisosSueltosByUser(db, us.ID)
		if err != nil {
			return nil, errors.Join(err, errors.New("error al cargar permisos sueltos del usuario"))
		}
	}

	return &user, nil
}

func GetUsuarios(db *sql.DB, query model.QueryUsuarios) ([]*model.Usuario, error) {
	filter_by_rol := ""
	if query.Rol != nil {
		filter_by_rol = "where id in (select usuario_id from rol_usuario where rol='%s')"
		filter_by_rol = fmt.Sprintf(filter_by_rol, *query.Rol)
	}

	sql := `select id, nombres,apellido1,apellido2,documento,celular,correo,sexo,direccion,estado,username,fecha_registro,fecha_update from usuarios %s`
	sql = fmt.Sprintf(sql, filter_by_rol)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []*model.Usuario{}
	for rows.Next() {
		u := model.Usuario{}
		er := parseRows(rows, &u)
		if er != nil {
			return nil, er
		}
		usuarios = append(usuarios, &u)
	}

	return usuarios, nil
}

func GetById2(db *sql.DB, input model.GetUser) (*model.ResponseMe, error) {
	us, err := GetById(db, input.ID)
	if err != nil {
		return nil, err
	}

	user := model.ResponseMe{}
	user.Usuario = us

	if input.ShowRoles {
		user.Roles, err = roles.GetRolesByUsuario(db, us.ID, input.ShowPermisos)
		if err != nil {
			return nil, errors.Join(err, errors.New("error al cargar roles al usuario"))
		}
	}

	if input.ShowPermisos {
		user.PermisosSueltos, err = permisos.GetPermisosSueltosByUser(db, us.ID)
		if err != nil {
			return nil, errors.Join(err, errors.New("error al cargar permisos sueltos del usuario"))
		}
	}

	return &user, nil
}
