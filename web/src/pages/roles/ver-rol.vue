<template>
  <q-dialog v-model="alert" persistent square>
    <q-card flat bordered v-on:click="checkClickSession()">
      <q-card-section>
        <div class="text-h6"> {{ input.nombre }} </div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <div class="row">
          <!-- <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Nombre:</b> {{ input.nombre }} </p> </div> -->
          <div class="col col-sm-6"> <p class="q-mb-xs"> <b>Jerarquia:</b> {{ input.jerarquia }} </p> </div>
          <div class="col col-sm-12"> <p class="q-mb-xs"> <b>Descripcion:</b> {{ input.descripcion }} </p> </div>
          <div class="col col-sm-12"> <p class="q-mt-lg"> <b>Permisos:</b> </p> </div>

          <template v-for="(r,i) in permisos_sel" :key="i">
            <div class="col col-sm-4"> <p class="q-mb-xs"> {{ r.nombre }} </p> </div>
          </template>
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
import RolesService from 'pages/roles/rolesService'
import PermisoService from 'pages/permisos/permisoService';
import { Notify } from 'quasar';
import { InputNewRol } from './type_roles';
import click from '../../shared/session'

export default {
  setup (_,vue) {
    const alert = ref(false);
    const is_edit = ref(false);
    const loading = ref(false);
    const loading_perms = ref(false);
    const input = ref(InputNewRol);
    const rolesService = new RolesService();
    const permisoService = new PermisoService();
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
      is_edit,
      input,
      permisos,
      permisos_sel,
      open,
      onSubmit,
      cerrar,
      checkClickSession:click.setup().checkClickSession,
    }
  }
}
</script>
