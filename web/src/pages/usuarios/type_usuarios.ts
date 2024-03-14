export interface NewUsuario {
  nombres: string;
  apellido1: string;
  apellido2?: string;
  documento?: string;
  celular?: string;
  correo?: string;
  sexo?: string;
  direccion?: string;
  username: string;
  password: string;
  roles: [];
  permisos_sueltos: [];
}
export interface UpdateUsuario {
  id: string;
  nombres: string;
  apellido1: string;
  apellido2: string;
  documento: string;
  celular: string;
  correo: string;
  sexo: string;
  direccion: string;
  username: string;
  password: string;
  roles: [];
  permisos_sueltos: [];
}

export interface QueryUsuarios {
  rol?: string;
}

export const InputNewUsuario: NewUsuario = {
  nombres: '',
  apellido1: '',
  apellido2: '',
  documento: '',
  celular: '',
  correo: '',
  sexo: '',
  direccion: '',
  username: '',
  password: '',
  roles: [],
  permisos_sueltos: []
}
