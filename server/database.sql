create table `usuarios`(
    `id` integer unsigned auto_increment not null primary key,
    `nombres` varchar(30) not null,
    `apellido1` varchar(30) not null,
    `apellido2` varchar(30),
    `documento` varchar(30),
    `celular` varchar(20),
    `correo` varchar(100),
    `sexo` char(1),
    `direccion` varchar(100),
    `estado` tinyint(1) not null default 1,
    `username` varchar(30) unique not null,
    `password` varchar(64) not null, -- hash
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    `fecha_update` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP
);

create table `roles`(
    `nombre` varchar(50) not null primary key,
    `descripcion` varchar(100),
    `jerarquia` tinyint(1) not null default 0,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00')
);

-- para permisos no hay crud, se agregan segun se crean las funcionaes 
-- en el propio codigo fuente 
create table `permisos`(
    -- en golang hay funciones, cada funcion representa un permiso, el nombre de esa funcion es el valor de metodo, (ej createPersona)
    `metodo` varchar(50) not null primary key,
    -- el nombre es el mismo que metodo, pero para que un humano lo lea, (ej crear persona)
    `nombre` varchar(50) not null,
    `descripcion` varchar(200),
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00')
);

create table `rol_permiso`(
    `rol` varchar(50) not null,
    `metodo` varchar(50) not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`rol`) references `roles`(`nombre`),
    foreign key(`metodo`) references `permisos`(`metodo`),
    primary key(`rol`,`metodo`)
);

create table `rol_usuario`(
    `rol` varchar(50) not null,
    `usuario_id` integer unsigned not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`rol`) references `roles`(`nombre`),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    primary key(`rol`,`usuario_id`)
);

create table `usuario_permiso`(
    `usuario_id` integer unsigned not null,
    `metodo` varchar(50) not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    foreign key(`metodo`) references `permisos`(`metodo`),
    primary key(`usuario_id`,`metodo`)
);

-- indice para optimizar la busqueda 
CREATE INDEX idx_username ON usuarios (username);

-- VALORES POR DEFECTO
INSERT INTO `usuarios` (`nombres`, `apellido1`, `username`, `password`)
VALUES
    ('Admin', '', 'admin', SHA2('admin', 256));

INSERT INTO `roles` (`nombre`, `descripcion`, `jerarquia`)
VALUES
    ('Administrador', 'Tiene acceso total al sistema.', 0);

INSERT INTO `permisos` (`metodo`, `nombre`, `descripcion`)
VALUES
    ('createUsuario', 'Crear Usuario', 'Permite crear un nuevo usuarios en el sistema.'),
    ('updateUsuario', 'Actualizar Usuario', 'Permite actualizar los datos de un usuarios en el sistema.'),
    ('createRol', 'Crear Rol', 'Permite crear un nuevo rol en el sistema.'),
    ('updateRol', 'Actualizar Rol', 'Permite actualizar los datos de un rol en el sistema.'),
    ('roles', 'Listar roles', 'Listar los roles en el sistema.'),
    ('permisos', 'Listar permisos', 'Listar los permisos en el sistema.'),
    ('usuarios', 'Listar usuarios', 'Listar los usuarios en el sistema.'),
    ('usuarioById', 'Listar usuario por id', 'Listar los datos de un usuario en el sistema.'),
    ('rolById', 'Listar rol por id', 'Listar los datos de un rol en el sistema.');

insert into `rol_permiso`(`rol`,`metodo`)
values 
    ('Administrador','createUsuario'),
    ('Administrador','updateUsuario'),
    ('Administrador','createRol'),
    ('Administrador','updateRol'),
    ('Administrador','roles'),
    ('Administrador','permisos'),
    ('Administrador','usuarios'),
    ('Administrador','usuarioById'),
    ('Administrador','rolById');

insert into `rol_usuario`(`rol`,`usuario_id`)
values 
    ('Administrador',1);







