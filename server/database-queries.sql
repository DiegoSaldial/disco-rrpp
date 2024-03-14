
-- listar todos los roles de un usuario, se puede agregar sus permisos con la consulta de abajo
select r.nombre, r.descripcion, r.jerarquia, r.fecha_registro,ru.usuario_id, ru.fecha_registro as fecha_asignado 
from roles r
left join rol_usuario ru on ru.rol = r.nombre
where ru.usuario_id = 43;

-- listar todos los permisos de un rol se puede agregar al listado de roles de arriba
select p.metodo, p.nombre, p.descripcion, p.fecha_registro, rp.fecha_registro as fecha_asignado 
from permisos p
left join rol_permiso rp on rp.metodo = p.metodo
where rp.rol = "Administrador";

-- listar todos los permisos sueltos de un usuario, se puede agregar como un array separado de los permisos de arriba
select p.metodo, p.nombre, p.descripcion, p.fecha_registro, up.fecha_registro as fecha_asignado 
from permisos p
inner join usuario_permiso up on up.metodo  = p.metodo 
where up.usuario_id = 34;



-- verificar si un usuario tiene un permiso ya sea suelto o en un rol 
SELECT 
    CASE 
        WHEN up.usuario_id IS NOT NULL THEN 'Directo'
        WHEN rp.rol IS NOT NULL THEN 'A través de roles' 
    END AS metodo_de_asignacion
FROM usuarios u
LEFT JOIN usuario_permiso up ON u.id = up.usuario_id AND up.metodo = 'crearUsuario'
LEFT JOIN rol_usuario ru ON u.id = ru.usuario_id
LEFT JOIN rol_permiso rp ON ru.rol = rp.rol AND rp.metodo = 'crearUsuario'
WHERE u.id = 34 AND (up.usuario_id IS NOT NULL OR rp.rol IS NOT NULL);




-- crear un usuario 
INSERT INTO `usuarios`(nombres, apellido1, apellido2, documento, celular, correo, sexo, direccion, username, password)
VALUES ('Nombre', 'Apellido1', 'Apellido2', 'Documento', 'Celular', 'Correo', 'M', 'Dirección', 'hola5', SHA2('S1nclave', 256));
-- 

-- actualizar usaurio
update `usuarios` set 
nombres="Diego", 
apellido1="Saldias", 
apellido2="Villa", 
documento="10721310", 
celular="78227092", 
correo="eclip@gmail.com", 
sexo="M", 
direccion="B/Morros"
where id = 73;
-- 

-- asignar nuevos credenciales
update `usuarios` set 
username = "dsaldias",
password = SHA2('S1nclave', 256)
where id = 73;
-- 

-- login 
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
where u.username = "dsaldias" 
and u.password = SHA2('S1nclave', 256);
-- 