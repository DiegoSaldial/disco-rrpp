/* eslint-disable @typescript-eslint/no-explicit-any */

export default class Validaciones {

  val_nombre(val:string){
    if(!val) return 'dato obligatorio';
    if(val.length > 50) return 'maximo 50 caracteres';
  };

  val_descripcion(val:string){
    if(!val) return true;
    if(val.length > 100) return 'maximo 100 caracteres';
  };

  val_jerarquia(val:number){
    if(val===0) return true;
    if(val===null || !val) return 'dato obigatorio';
  };

}


