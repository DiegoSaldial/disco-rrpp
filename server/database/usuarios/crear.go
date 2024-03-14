package usuarios

import (
	"auth/graph/model"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func Crear(db *sql.DB, input model.NewUsuario) (*model.Usuario, error) {
	if err := permisos_obligatorios(input.Roles, input.PermisosSueltos); err != nil {
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := `
	INSERT INTO usuarios(nombres, apellido1, apellido2, documento, celular, correo, sexo, direccion, username, password)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, SHA2(?, 256));
	`
	res, err := tx.Exec(sql,
		input.Nombres,
		input.Apellido1,
		input.Apellido2,
		input.Documento,
		input.Celular,
		input.Correo,
		input.Sexo,
		input.Direccion,
		input.Username,
		input.Password,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, _ := res.LastInsertId()

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

func asignarRoles(tx *sql.Tx, roles []string, userid int64) error {
	user_rols := "replace into `rol_usuario`(`rol`,`usuario_id`) values %s"
	places := make([]string, len(roles))
	args := make([]interface{}, len(roles)*2)

	for i, r := range roles {
		places[i] = "(?,?)"
		args[i*2] = r
		args[i*2+1] = userid
	}

	user_rols = fmt.Sprintf(user_rols, strings.Join(places, ", "))
	_, err := tx.Exec(user_rols, args...)
	return err
}

func asignarPermisos(tx *sql.Tx, permisosSueltos []string, userid int64) error {
	user_perms := "replace into `usuario_permiso`(`usuario_id`,`metodo`) values %s"
	places2 := make([]string, len(permisosSueltos))
	args2 := make([]interface{}, len(permisosSueltos)*2)

	for i, p := range permisosSueltos {
		places2[i] = "(?,?)"
		args2[i*2] = userid
		args2[i*2+1] = p
	}

	user_perms = fmt.Sprintf(user_perms, strings.Join(places2, ", "))
	_, err := tx.Exec(user_perms, args2...)
	return err
}
