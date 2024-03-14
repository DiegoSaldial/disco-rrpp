<template>
  <div class="q-pa-sm">
    <h6 class="q-my-sm text-center"> Usuarios del sistema </h6>
    <q-table
      color="primary"
      :rows="rows"
      :columns="columns"
      row-key="id"
      hide-pagination :rows-per-page-options="[0]"
      :visible-columns="visibleColumns"
      :filter="filter"
      :loading="loading">

      <template v-slot:top-left>
        <q-toggle v-model="more_datos" @update:model-value="columnas()" color="orange" label="mostrar otros datos" class="q-my-none" />
      </template>

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model.trim="filter" placeholder="buscar ..." class="q-mx-none">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
        <q-select outlined dense v-model="rol_select" :options="roles" label="Filtro de rol" option-value="nombre" option-label="nombre" :loading="loading_roles" :disable="loading" clearable/>
        <q-btn label="Registrar" color="green" icon="person_add" class="q-ml-xs" square @click="registrar()" />
      </template>

      <template v-slot:loading>
        <q-inner-loading showing color="primary" />
      </template>

      <template v-slot:body-cell-fecha_registro="props">
        <q-td :props="props">
          {{ parseFecha(props.row.fecha_registro, true) }}
        </q-td>
      </template>

      <template v-slot:body-cell-fecha_update="props">
        <q-td :props="props">
          {{ parseFecha(props.row.fecha_update, true) }}
        </q-td>
      </template>

      <template v-slot:body-cell-estado="props">
        <q-td :props="props">
          {{ props.row.estado?'Activo':'Inactivo' }}
        </q-td>
      </template>

      <template v-slot:body-cell-opt="props">
        <q-td :props="props">
          <q-btn color="green-10" square flat icon="more_vert" size="small">
            <q-menu anchor="top right" self="top left">
              <q-list style="min-width: 110px">
                <q-item clickable v-ripple @click="visualizar(props.row)" >
                  <q-item-section avatar>
                    <q-icon color="accent" name="visibility" right class="q-px-none" />
                  </q-item-section>
                  <q-item-section> Ver </q-item-section>
                </q-item>
                <q-item clickable v-ripple @click="actualizar(props.row)" >
                  <q-item-section avatar>
                    <q-icon color="accent" name="edit" right class="q-px-none" />
                  </q-item-section>
                  <q-item-section> Editar </q-item-section>
                </q-item>
              </q-list>
            </q-menu>
          </q-btn>
        </q-td>
      </template>

    </q-table>

    <Registrar ref="refRegistrar" v-on:success="listar" />
    <Ver ref="refVer"/>
  </div>
</template>

<script>
import { onMounted, ref,watch } from 'vue';
import UsuariosService from './usuariosService'
import RolesService from '../roles/rolesService';
import Registrar from './registrar_usuario.vue'
import Ver from './ver-usuario.vue'
import { parseFecha } from 'stores/utils'
import { useRoute, useRouter } from 'vue-router';

const columns = [
  { name: 'id', align: 'center', label: 'ID', field: 'id', sortable: true },
  { name: 'nombres', label: 'Nombres', field: 'nombres', sortable: true },
  { name: 'apellido1', label: 'apellido1', field: 'apellido1' },
  { name: 'apellido2', label: 'apellido2', field: 'apellido2' },
  { name: 'documento', label: 'documento', field: 'documento' },
  { name: 'celular', label: 'celular', field: 'celular' },
  { name: 'correo', label: 'correo', field: 'correo' },
  { name: 'sexo', label: 'sexo', field: 'sexo' },
  { name: 'direccion', label: 'direccion', field: 'direccion' },
  { name: 'username', label: 'username', field: 'username' },
  { name: 'fecha_registro', label: 'registro', field: 'fecha_registro' },
  { name: 'fecha_update', label: 'modificado', field: 'fecha_update' },
  { name: 'estado', label: 'estado', field: 'estado' },
  { name: 'opt', label: '', field: 'opt' },
]

export default {
  components:{ Registrar,Ver },
  setup () {
    const usuariosService = new UsuariosService();
    const rolesService = new RolesService();
    const rows = ref([]);
    const loading = ref(false)
    const filter = ref('')
    const refRegistrar = ref()
    const refVer = ref()
    const more_datos = ref(false)
    const visibleColumns = ref([])
    const roles = ref([])
    const loading_roles = ref(false)
    const rol_select = ref(null)
    const router = useRouter()
    const route = useRoute()

    const listar = async () => {
      loading.value = true;
      rows.value = [];
      const query = {
        rol: route.query.rol?route.query.rol:null,
      }
      let res = await usuariosService.usuarios(query).then(u=>u).catch(e=>e)
      if(res.usuarios) rows.value = res.usuarios;
      loading.value = false;
    }

    const listarRoles = async () => {
      loading_roles.value = true;
      let res = await rolesService.roles(false)
      roles.value = res.roles;
      loading_roles.value = false;
      let r = route.query.rol?route.query.rol:null;
      rol_select.value = roles.value.find(x=>x.nombre==r);
    }

    const registrar = () => refRegistrar.value.open();
    const actualizar = (item) => refRegistrar.value.open(item.id);
    const visualizar = (item) => refVer.value.open(item.id);

    const columnas = () => {
      if(more_datos.value) visibleColumns.value = ['id','nombres','apellido1','apellido2','documento','celular','correo','sexo','direccion','username','fecha_registro','fecha_update','estado','opt'];
      else visibleColumns.value = ['id','nombres','apellido1','apellido2','opt'];
    }

    const setParams = async ()=> {
      if(rol_select.value) await router.push({path:'',query:{rol:rol_select.value.nombre}});
      else await router.push({path:'',query:null});
      listar();
    }

    watch(()=>rol_select.value, ()=>setParams())

    onMounted(()=>{
      columnas();
      listar();
      listarRoles();
    })

    return {
      columns,
      rows,
      loading,
      filter,
      refRegistrar,
      refVer,
      visibleColumns,
      more_datos,
      columnas,
      roles,
      loading_roles,
      rol_select,
      listar,
      registrar,
      actualizar,
      visualizar,
      parseFecha
    }
  }
}
</script>
