<template>
  <q-dialog v-model="alert" square>
    <q-card flat bordered v-on:click="checkClickSession()">
      <q-card-section>
        <div class="text-h6 text-purple"> Ver usuario </div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <div class="row">
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Nombres:</b> {{ input.nombres }} </p> </div>
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Apellido 1:</b> {{ input.apellido1 }} </p> </div>
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Apellido 2:</b> {{ input.apellido2 }} </p> </div>
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Documento:</b> {{ input.documento }} </p> </div>
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Celular:</b> {{ input.celular }} </p> </div>
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Sexo:</b> {{ input.sexo }} </p> </div>
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Correo:</b> {{ input.correo }} </p> </div>
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Direccion:</b> {{ input.direccion }} </p> </div>
          <div class="col col-sm-12"> <p class="q-mb-lg"> <b>Username:</b> {{ input.username }} </p> </div>
          <template v-for="(r,i) in roles_sel" :key="i">
            <div class="col col-sm-4"> <p class="q-mb-xs"> <b>Rol:</b> {{ r.nombre }} </p> </div>
            <div class="col col-sm-8"> <p class="q-mb-xs"> <b>Permisos:</b> {{ r.xpermisos.join(', ') }} </p> </div>
          </template>

          <div class="col col-sm-12"> <p class="q-mt-lg"> <b>Permisos sueltos:</b> {{ permisos_sel }} </p> </div>
        </div>

        <div class="q-mt-md" :align="'right'">
          <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
          <q-btn label="cerrar" color="red" icon="close" square flat @click="cerrar()"/>
        </div>
      </q-card-section>

    </q-card>
  </q-dialog>
</template>


<script>
import { ref } from 'vue'
import { InputNewUsuario } from './type_usuarios';
import UsuariosService from './usuariosService';
import click from '../../shared/session'

export default {
  setup () {
    const alert = ref(false);
    const loading = ref(false);
    const input = ref(InputNewUsuario);
    const usuarioService = new UsuariosService();
    const roles_sel = ref([])
    const permisos_sel = ref([])

    const open = (id=null)=> {
      loading.value = false;
      alert.value = true;
      roles_sel.value = []
      permisos_sel.value = []
      if(id) getUserById(id);
    }

    const getUserById = async (id) => {
      loading.value = true;
      let res = await usuarioService.usuarioById(id).then(d=>d).catch(e=>e)
      const xroles = res.usuarioById.roles.map(item => item);
      const xpermisos = res.usuarioById.permisos_sueltos.map(item => item.nombre);
      xroles.forEach(r => {
        r.xpermisos = r.permisos.map(item => item.nombre)
      });

      let us = res.usuarioById.usuario;
      input.value = us;
      roles_sel.value = xroles;
      permisos_sel.value = xpermisos.join(', ');
      loading.value = false;
    }

    const cerrar = () => {
      alert.value = false;
    }

    return {
      alert,
      loading,
      input,
      roles_sel,
      permisos_sel,
      open,
      cerrar,
      checkClickSession:click.setup().checkClickSession,
    }
  }
}
</script>
