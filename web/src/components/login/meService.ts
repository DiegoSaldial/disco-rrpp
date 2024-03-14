import gql from 'graphql-tag';
import { query } from 'stores/server'

export default class MeService {
  async me() {

    const sql = gql`
      query{
        me(input:{show_roles:true, show_permisos:true}){
          usuario{id nombres apellido1 apellido2 documento celular correo sexo direccion username fecha_registro fecha_update estado}
          roles{nombre jerarquia fecha_asignado permisos{metodo nombre fecha_asignado}}
          permisos_sueltos{metodo nombre fecha_asignado}
        }
      }
      `;

      return await query(sql,{}).then(d=>d).catch(e=>e)

  }
}
