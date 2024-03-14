import gql from 'graphql-tag';
import {mutar} from 'stores/server'

export default class LoginService {
  async login(username:string, password:string) {

    const sql = gql`
      mutation login($input:NewLogin!) {login(input:$input){token refreshToken}}
    `;

    const v = { input:{username:username,password:password}}
    return await mutar(sql,v).then(d=>d).catch(e=>e)

  }
}
