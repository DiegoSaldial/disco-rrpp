/* eslint-disable @typescript-eslint/no-explicit-any */
import gql from 'graphql-tag';
import { query, mutar } from 'stores/server'
import { NewUsuario,QueryUsuarios,UpdateUsuario } from './type_usuarios';

export default class UsuariosService {
  async usuarios(xquery: QueryUsuarios) {
    const sql = gql`
      query usuarios($query:QueryUsuarios!){
        usuarios(query:$query){
          id
          nombres
          apellido1
          apellido2
          documento
          celular
          correo
          sexo
          direccion
          username
          fecha_registro
          fecha_update
          estado
        }
      }
      `;

      return await query(sql,{query:xquery}).then((d:any)=>d).catch((e:any)=>e)
  }

  async createUsuario(input:NewUsuario) {
    const sql = gql`
      mutation createUsuario($input:NewUsuario!){
        createUsuario(input:$input){
          id
          nombres
        }
      }
      `;
      return await mutar(sql,{input:input}).then(d=>d).catch(e=>e)
  }

  async updateUsuario(input:UpdateUsuario) {
    const sql = gql`
      mutation updateUsuario($input:UpdateUsuario!){
        updateUsuario(input:$input){
          id
          nombres
        }
      }
      `;
      return await mutar(sql,{input:input}).then(d=>d).catch(e=>e)
  }

  async usuarioById(id:string) {
    const sql = gql`
      query{
        usuarioById(input:{
          id:${id},
          show_roles: true,
          show_permisos: true
        }){
          usuario{
            id
            nombres
            apellido1
            apellido2
            documento
            celular
            correo
            sexo
            direccion
            username
            fecha_registro
            fecha_update
            estado
          }
          roles{nombre permisos{metodo, nombre}}
          permisos_sueltos{metodo, nombre}
        }
      }
      `;
      return await query(sql,{}).then(d=>d).catch(e=>e)
  }
}
