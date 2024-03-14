<template>
  <div class="q-pa-md column items-center" >
    <h5 class="q-mb-lg"> Acceder al sistema </h5>
    <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md" >
      <q-input filled v-model.trim="username" label="Nombre de usuario" lazy-rules dense :color="$q.dark.isActive?'orange':'primary'" :rules="[ val => val && val.length > 0 || 'dato obligatorio']" />

      <q-input filled type="text" v-model.trim="clave" label="Clave de acceso" lazy-rules dense :color="$q.dark.isActive?'orange':'primary'" :rules="[ val => val && val.length > 0 || 'dato obligatorio']" />

      <div class="column items-center">
        <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
        <q-btn :disable="loading" icon="person" stretch label="Ingresar" type="submit" color="orange" />
      </div>
    </q-form>

  </div>
</template>

<script>
import { ref } from 'vue'
import LoginService from './loginService'
import {useLoginStore} from 'stores/login-store'
import MeService from './meService'

export default {
  setup () {

    const username = ref('admin')
    const clave = ref('admin')
    const loading = ref(false)
    const service = new LoginService();
    const meService = new MeService();
    const useLogin = useLoginStore();

    async function getMe(){
      loading.value = true;
      let meres = await meService.me().then(e=>e).catch(e=>e);
      if(meres.me){
        useLogin.setUser(meres.me);
      }
      loading.value = false;
    }

    return {
      username,
      clave,
      loading,

      async onSubmit () {
        loading.value = true;
        let res = await service.login(username.value,clave.value).then(x=>x).catch(e=>e)
        loading.value = false;

        if( res.login ){
          const l = res.login;
          useLogin.setToken(l.token,l.refreshToken)
          getMe();
        }

      },

      onReset () {
        username.value = null
        clave.value = null
      }
    }
  }
}
</script>
