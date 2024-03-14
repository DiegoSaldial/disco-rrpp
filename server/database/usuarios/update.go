package usuarios

import (
	"auth/graph/model"
	"database/sql"
	"strconv"
)

func Actualizar(db *sql.DB, input model.UpdateUsuario) (*model.Usuario, error) {
	if err := permisos_obligatorios(input.Roles, input.PermisosSueltos); err != nil {
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := `
	update usuarios set 
	nombres=?, 
	apellido1=?, 
	apellido2=?, 
	documento=?, 
	celular=?, 
	correo=?, 
	sexo=?, 
	direccion=?, 
	username=?
	where id= ? 
	`
	_, err = tx.Exec(sql,
		input.Nombres,
		input.Apellido1,
		input.Apellido2,
		input.Documento,
		input.Celular,
		input.Correo,
		input.Sexo,
		input.Direccion,
		input.Username,
		input.ID,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(input.Password) > 0 {
		_, err = tx.Exec("update usuarios set password=SHA2(?, 256) where id = ?", input.Password, input.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	_, err = tx.Exec("delete from rol_usuario where usuario_id = ?", input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("delete from usuario_permiso where usuario_id = ?", input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	id, _ := strconv.ParseInt(input.ID, 10, 64)

	// asignar roles
	if len(input.Roles) > 0 {
		err = asignarRoles(tx, input.Roles, id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// fi asignar roles

	// asignar permisos sueltos
	if len(input.PermisosSueltos) > 0 {
		err = asignarPermisos(tx, input.PermisosSueltos, id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// fin permisos sueltos

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return GetById(db, strconv.FormatInt(id, 10))
}
