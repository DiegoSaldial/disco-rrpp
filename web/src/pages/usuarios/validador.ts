/* eslint-disable @typescript-eslint/no-explicit-any */

export default class Validaciones {

  val_nombre(val:string){
    if(!val) return 'dato obligatorio';
    if(val.length > 30) return 'maximo 30 caracteres';
  };

  val_apellido1(val:string){
    if(!val) return 'dato obligatorio';
    if(val.length > 30) return 'maximo 30 caracteres';
  };

  val_apellido2(val:string){
    if(!val) return true;
    if(val.length > 30) return 'maximo 30 caracteres';
  };

  val_documento(val:string){
    if(!val) return true;
    if(val.length > 30) return 'maximo 30 caracteres';
  };

  val_celular(val:string){
    if(!val) return true;
    if(val.length > 20) return 'maximo 20 caracteres';
  };

  val_correo(val:string){
    if(!val) return true;
    if(val.length > 100) return 'maximo 100 caracteres';
  };

  val_direccion(val:string){
    if(!val) return true;
    if(val.length > 100) return 'maximo 100 caracteres';
  };

  val_username(val:string){
    if(!val) return 'dato obligatorio';
    if(val.length > 30) return 'maximo 30 caracteres';
  };

  val_password(val:string,input: any){
    if(!val && !input.id) return 'datos obligatorio';
    if(!val && input.id) return true;
    if(val.length > 64) return 'maximo 64 caracteres';
  };

}


