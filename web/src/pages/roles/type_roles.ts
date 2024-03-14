export interface NewRol {
  nombre: string;
  descripcion?: string;
  jerarquia: number;
  permisos: []
}

export const InputNewRol: NewRol = {
  nombre: '',
  descripcion: '',
  jerarquia: 0,
  permisos: [],
}
