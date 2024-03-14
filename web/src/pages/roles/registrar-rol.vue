<template>
  <q-dialog v-model="alert" persistent square>
    <q-card flat bordered v-on:click="checkClickSession()">
      <q-card-section>
        <div class="text-h6"> {{ is_edit?'Actualizar':'Registrar' }} </div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-form @submit="onSubmit" >
          <div class="row q-col-gutter-xs">
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.nombre" label="Nombre" lazy-rules dense :rules="[ val => validaciones.val_nombre(val) ]" /> </div>
            <div class="col col-sm-6"> <q-input filled v-model.trim="input.descripcion" label="Descripcion" lazy-rules dense counter :rules="[ val => validaciones.val_descripcion(val) ]"/> </div>
            <div class="col col-sm-6"> <q-input filled v-model.number="input.jerarquia" label="Jerarquia" type="number" lazy-rules dense counter :rules="[ val => validaciones.val_jerarquia(val)]" /> </div>


            <div class="col col-sm-12">
              <q-table
                flat
                color="orange"
                :loading="loading_perms"
                title="permisos"
                hide-pagination :rows-per-page-options="[0]"
                dense
                :rows="permisos"
                :columns="columnas_perm"
                row-key="metodo"
                :selected-rows-label="getSelectedString"
                selection="multiple"
                v-model:selected="permisos_sel"
              />
            </div>
          </div>

          <div class="q-mt-md" :align="'right'">
            <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
            <q-btn :disable="loading" label="cerrar" color="red" icon="close" square flat @click="cerrar()"/>
            <q-btn :disable="loading" :label="is_edit?'Actualizar':'Registrar'" icon="done" type="submit" color="green" square/>
          </div>
        </q-form>
      </q-card-section>

    </q-card>
  </q-dialog>
</template>


<script>
import { ref } from 'vue'
import RolesService from 'pages/roles/rolesService'
import PermisoService from 'pages/permisos/permisoService';
import { Notify } from 'quasar';
import { InputNewRol } from './type_roles';
import Validaciones from './validador'
import click from '../../shared/session'

const columnas_perm = [
  { name: 'nombre', label: 'Nombre', field: 'nombre', align:'left' },
  { name: 'metodo', label: 'Codigo', field: 'metodo', align:'left' },
]

export default {
  setup (_,vue) {
    const alert = ref(false);
    const is_edit = ref(false);
    const loading = ref(false);
    const loading_perms = ref(false);
    const input = ref(InputNewRol);
    const rolesService = new RolesService();
    const permisoService = new PermisoService();
    const validaciones = new Validaciones();
    const permisos = ref([])
    const permisos_sel = ref([])

    const open = (id=null)=> {
      is_edit.value = false;
      alert.value = true;
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
      cargarPermisos()
      if(id) getRolById(id);
    }

    const cargarPermisos = async () => {
      loading_perms.value = true;
      let res = await permisoService.permisos().then(d=>d).catch(e=>e)
      permisos.value = res.permisos;
      loading_perms.value = false;
    }

    const getRolById = async (id) => {
      is_edit.value = true;
      loading.value = true;
      let res = await rolesService.rolById(id).then(d=>d).catch(e=>e)
      const xpermisos = res.rolById.permisos.map(item => item);
      let rol = res.rolById;
      delete rol.estado;
      delete rol.fecha_registro;
      delete rol.fecha_update;
      input.value = rol;
      permisos_sel.value = xpermisos;
      loading.value = false;
    }

    const onSubmit = async () => {
      if(!input.value.nombre) return;
      const xpermisos = permisos_sel.value.map(item => item.metodo);
      input.value.permisos = xpermisos;
      if(xpermisos.length == 0){
        Notify.create({message:'Seleccione al menos un permiso',color:'red'})
        return;
      }

      if(is_edit.value) actualizar();
      else registrar();
    }

    const registrar = async () => {
      loading.value = true;
      let res = await rolesService.createRol(input.value).then(d=>d).catch(e=>e)
      loading.value = false;
      if(res.createRol){
        cerrar();
        vue.emit('success')
      }
    }

    const actualizar = async () => {
      loading.value = true;
      let res = await rolesService.updateRol(input.value).then(d=>d).catch(e=>e)
      loading.value = false;
      if(res.updateRol){
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
      loading_perms,
      columnas_perm,
      is_edit,
      input,
      permisos,
      permisos_sel,
      open,
      onSubmit,
      cerrar,
      validaciones,
      checkClickSession:click.setup().checkClickSession,
      getSelectedString () {
        return `${permisos_sel.value.length} de ${permisos.value.length}`
      }
    }
  }
}
</script>
