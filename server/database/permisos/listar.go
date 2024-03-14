package permisos

import (
	"auth/graph/model"
	"database/sql"
	"errors"
	"os"
)

func VerificarPermiso(db *sql.DB, userid, metodo string) error {
	sq := `
	SELECT 
		CASE 
			WHEN up.usuario_id IS NOT NULL THEN 'Directo'
			WHEN rp.rol IS NOT NULL THEN 'A través de roles' 
		END AS metodo_de_asignacion
	FROM usuarios u
	LEFT JOIN usuario_permiso up ON u.id = up.usuario_id AND up.metodo = ?
	LEFT JOIN rol_usuario ru ON u.id = ru.usuario_id
	LEFT JOIN rol_permiso rp ON ru.rol = rp.rol AND rp.metodo = ?
	WHERE u.id = ? AND (up.usuario_id IS NOT NULL OR rp.rol IS NOT NULL)
	`
	texto := ""
	err := db.QueryRow(sq, metodo, metodo, userid).Scan(&texto)

	if err == sql.ErrNoRows {
		show_name := os.Getenv("AUTH_SHOW_NAME_PERMISO")
		if show_name == "1" {
			return errors.New("no tiene permiso para: " + metodo)
		} else {
			return errors.New("no tiene permiso")
		}
	}
	return err
}

func GetPermisosByRol(db *sql.DB, rol string) ([]*model.ResponsePermisoMe, error) {
	sql := `
	select p.metodo, p.nombre, p.descripcion, p.fecha_registro, rp.fecha_registro as fecha_asignado 
	from permisos p
	left join rol_permiso rp on rp.metodo = p.metodo
	where rp.rol = ?;
	`

	rows, err := db.Query(sql, rol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	perms := []*model.ResponsePermisoMe{}
	for rows.Next() {
		p := model.ResponsePermisoMe{}
		er := parse(rows, &p)
		if er != nil {
			return nil, er
		}
		perms = append(perms, &p)
	}

	return perms, nil
}

func GetPermisosSueltosByUser(db *sql.DB, userid string) ([]*model.ResponsePermisoMe, error) {
	sql := `
	select p.metodo, p.nombre, p.descripcion, p.fecha_registro, up.fecha_registro as fecha_asignado 
	from permisos p
	inner join usuario_permiso up on up.metodo  = p.metodo 
	where up.usuario_id = ?
	`

	rows, err := db.Query(sql, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	perms := []*model.ResponsePermisoMe{}
	for rows.Next() {
		p := model.ResponsePermisoMe{}
		er := parse(rows, &p)
		if er != nil {
			return nil, er
		}
		perms = append(perms, &p)
	}

	return perms, nil
}

func GetPermisos(db *sql.DB) ([]*model.Permiso, error) {
	sql := `select metodo,nombre, descripcion,fecha_registro from permisos`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permisos := []*model.Permiso{}
	for rows.Next() {
		p := model.Permiso{}
		er := parseRows(rows, &p)
		if er != nil {
			return nil, er
		}
		permisos = append(permisos, &p)
	}

	return permisos, nil
}
