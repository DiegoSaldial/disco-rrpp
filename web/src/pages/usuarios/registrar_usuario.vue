<template>
  <q-dialog v-model="alert" persistent square >
    <q-card flat bordered v-on:click="checkClickSession()" >
      <q-card-section>
        <div class="text-h6"> {{ input.id?'Actualizar':'Registrar' }} </div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-form @submit="onSubmit" >
          <div class="row q-col-gutter-xs">
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.nombres" label="Nombres" dense :rules="[val => validaciones.val_nombre(val)]" /> </div>
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.apellido1" label="Apellido 1" lazy-rules dense :rules="[val => validaciones.val_apellido1(val)]" /> </div>
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.apellido2" label="Apellido 2" lazy-rules dense counter :rules="[val => validaciones.val_apellido2(val)]" /> </div>
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.documento" label="Documento" lazy-rules dense counter :rules="[val => validaciones.val_documento(val)]" /> </div>
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.celular" label="Celular" lazy-rules dense counter :rules="[val => validaciones.val_celular(val)]" /> </div>
            <div class="col col-sm-6">
              <q-radio v-model="input.sexo" val="M" label="Masculino" class="q-pa-none" />
              <q-radio v-model="input.sexo" val="F" label="Femenino" class="q-pa-none" />
            </div>
            <div class="col col-sm-12"> <q-input filled v-model.trim="input.correo" label="Correo" lazy-rules dense counter :rules="[val => validaciones.val_correo(val)]" /> </div>
            <div class="col col-sm-12"> <q-input filled v-model.trim="input.direccion" label="Direccion" lazy-rules dense counter :rules="[val => validaciones.val_direccion(val)]" /> </div>
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.username" label="username" lazy-rules dense counter :rules="[val => validaciones.val_username(val)]" /> </div>
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.password" label="password" :placeholder="input.id?'vacio sin cambios':''" lazy-rules dense counter :rules="[val => validaciones.val_password(val,input)]" /> </div>

            <div class="col col-sm-12">
              <q-list padding bordered>
                <q-expansion-item popup header-class="text-purple" default-opened expand-separator icon="group_add" label="Roles" caption="Un rol contiene un grupo de permisos">
                  <q-table flat color="orange" :loading="loading_roles" title="" hide-pagination :rows-per-page-options="[0]" dense :rows="roles" :columns="columnas_rols" row-key="nombre"  selection="multiple" v-model:selected="roles_sel" />
                </q-expansion-item>

                <q-expansion-item popup header-class="text-orange" expand-separator icon="key" label="Permisos sueltos" caption="Independientes del rol">
                  <q-table flat color="orange" :loading="loading_perms" title="" hide-pagination :rows-per-page-options="[0]" dense :rows="permisos" :columns="columnas_perm" row-key="metodo"  selection="multiple" v-model:selected="permisos_sel" />
                </q-expansion-item>
              </q-list>
            </div>


            <!-- <div class="col col-sm-12">
              <q-select filled v-model="roles_sel" multiple :options="roles" use-chips stack-label label="Roles" option-value="nombre" option-label="nombre" :loading="loading_roles"/>
            </div>
            <div class="col col-sm-12 q-pt-md">
              <q-select filled v-model="permisos_sel" multiple :options="permisos" use-chips stack-label label="Permisos sueltos" option-value="metodo" option-label="metodo" :loading="loading_perms"/>
            </div> -->
          </div>

          <div class="q-mt-md" :align="'right'">
            <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
            <q-btn :disable="loading" label="cerrar" color="red" icon="close" square flat @click="cerrar()"/>
            <q-btn :disable="loading" :label="input.id?'Actualizar':'Registrar'" icon="done" type="submit" color="green" square/>
          </div>
        </q-form>
      </q-card-section>

    </q-card>
  </q-dialog>
</template>


<script>
import { ref } from 'vue'
import { InputNewUsuario } from './type_usuarios';
import UsuariosService from './usuariosService';
import RolesService from 'pages/roles/rolesService'
import PermisoService from 'pages/permisos/permisoService';
import { Notify } from 'quasar';
import Validaciones from './validador';
import click from '../../shared/session'

const columnas_rols = [
  { name: 'nombre', label: '', field: 'nombre', align:'left' },
  { name: 'descripcion', label: '', field: 'descripcion', align:'left' },
  { name: 'jerarquia', label: '', field: 'jerarquia', align:'left' },
]
const columnas_perm = [
  { name: 'nombre', label: '', field: 'nombre', align:'left' },
  { name: 'jerarquia', label: '', field: 'descripcion', align:'left' },
]

export default {
  setup (_,vue) {
    const alert = ref(false);
    const loading = ref(false);
    const loading_roles = ref(false);
    const loading_perms = ref(false);
    const input = ref(InputNewUsuario);
    const usuarioService = new UsuariosService();
    const rolesService = new RolesService();
    const permisoService = new PermisoService();
    const validaciones = new Validaciones();
    const roles = ref([])
    const roles_sel = ref([])
    const permisos = ref([])
    const permisos_sel = ref([])

    const open = (id=null)=> {
      alert.value = true;
      roles_sel.value = []
      permisos_sel.value = []
      delete input.value.id;
      for (let key in input.value) {
        if (typeof input.value[key] === 'string') {
          input.value[key] = '';
        } else if (Array.isArray(input.value[key])) {
          input.value[key] = [];
        } else if (typeof input.value[key] === 'number') {
          input.value[key] = 0;
        }
      }
      cargarRoles()
      cargarPermisos()
      if(id) getUserById(id);
    }

    const cargarRoles = async () => {
      loading_roles.value = true;
      let res = await rolesService.roles(false).then(d=>d).catch(e=>e)
      roles.value = res.roles;
      loading_roles.value = false;
    }

    const cargarPermisos = async () => {
      loading_perms.value = true;
      let res = await permisoService.permisos().then(d=>d).catch(e=>e)
      permisos.value = res.permisos;
      loading_perms.value = false;
    }

    const getUserById = async (id) => {
      loading.value = true;
      let res = await usuarioService.usuarioById(id).then(d=>d).catch(e=>e)
      const xroles = res.usuarioById.roles.map(item => item);
      const xpermisos = res.usuarioById.permisos_sueltos.map(item => item);
      let us = res.usuarioById.usuario;
      Object.entries(us).forEach(([key, value]) => {
        if (value === null) us[key] = '';
      });
      delete us.estado;
      delete us.fecha_registro;
      delete us.fecha_update;
      us.password = '';
      input.value = us;
      roles_sel.value = xroles;
      permisos_sel.value = xpermisos;
      loading.value = false;
    }

    const onSubmit = async () => {
      const xroles = roles_sel.value.map(item => item.nombre);
      const xpermisos = permisos_sel.value.map(item => item.metodo);
      input.value.roles = xroles;
      input.value.permisos_sueltos = xpermisos;
      if(xroles.length == 0 && xpermisos.length == 0){
        Notify.create({message:'Seleccione al menos un rol o un permiso',color:'red'})
        return;
      }

      if(input.value.id) actualizar();
      else registrar();
    }

    const registrar = async () => {
      loading.value = true;
      let res = await usuarioService.createUsuario(input.value).then(d=>d).catch(e=>e)
      loading.value = false;
      if(res.createUsuario){
        cerrar();
        vue.emit('success')
      }
    }

    const actualizar = async () => {
      loading.value = true;
      let res = await usuarioService.updateUsuario(input.value).then(d=>d).catch(e=>e)
      loading.value = false;
      if(res.updateUsuario){
        cerrar();
        vue.emit('success')
      }
    }

    const cerrar = () => {
      alert.value = false;

    }

    return {
      alert,
      loading,
      loading_roles,
      loading_perms,
      input,
      roles,
      roles_sel,
      permisos,
      permisos_sel,
      checkClickSession:click.setup().checkClickSession,
      open,
      onSubmit,
      cerrar,
      columnas_rols,
      columnas_perm,
      validaciones,
    }
  }
}
</script>
