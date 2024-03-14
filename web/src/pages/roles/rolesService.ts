import gql from 'graphql-tag';
import { query, mutar } from 'stores/server'
import { NewRol } from './type_roles';

export default class RolesService {

  async roles(show_permisos:boolean) {
    const sql = gql`
      query roles($show_permisos:Boolean!){
        roles(show_permisos:$show_permisos){
          nombre
          descripcion
          jerarquia
          fecha_registro
          permisos{metodo nombre}
        }
      }
      `;
      return await query(sql,{show_permisos:show_permisos}).then(d=>d).catch(e=>e)
  }

  async createRol(input: NewRol) {
    const sql = gql`
      mutation createRol($input:NewRol!){
        createRol(input:$input){
          nombre
        }
      }
      `;
      return await mutar(sql,{input:input}).then(d=>d).catch(e=>e)
  }

  async updateRol(input: NewRol) {
    const sql = gql`
      mutation updateRol($input:NewRol!){
        updateRol(input:$input){
          nombre
        }
      }
      `;
      return await mutar(sql,{input:input}).then(d=>d).catch(e=>e)
  }


  async rolById(rol:string) {
    const sql = gql`
      query rolById($rol:String!){
        rolById(rol:$rol){
          nombre
          descripcion
          jerarquia
          fecha_registro
          permisos{metodo nombre descripcion}
        }
      }
      `;
      return await query(sql,{rol:rol}).then(d=>d).catch(e=>e)
  }

}
