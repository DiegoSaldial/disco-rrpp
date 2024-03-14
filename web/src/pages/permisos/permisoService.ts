import gql from 'graphql-tag';
import { query } from 'stores/server'

export default class PermisoService {
  async permisos() {

    const sql = gql`
      query{
        permisos{metodo nombre descripcion fecha_registro}
      }
      `;

      return await query(sql,{}).then(d=>d).catch(e=>e)
  }

}
