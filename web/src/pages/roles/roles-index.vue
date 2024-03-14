<template>
  <div class="q-pa-sm">
    <q-table
      title="Roles del sistema"
      color="primary"
      :rows="rows"
      :columns="columns"
      row-key="nombre"
      hide-pagination :rows-per-page-options="[0]"
      :filter="filter"
      :loading="loading">

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model.trim="filter" placeholder="buscar ..." class="q-mx-none">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
        <q-btn label="Registrar" color="green" icon="post_add" class="q-ml-xs" square @click="registrar()" />
      </template>

      <template v-slot:loading>
        <q-inner-loading showing color="primary" />
      </template>

      <template v-slot:body-cell-fecha_registro="props">
        <q-td :props="props">
          {{ parseFecha(props.row.fecha_registro, true) }}
        </q-td>
      </template>

      <template v-slot:body-cell-permisos="props">
        <q-td :props="props">
          {{ props.row.permisos.length }}
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
    <Ver ref="refVer" />
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';
import RolesService from 'pages/roles/rolesService'
import Registrar from 'pages/roles/registrar-rol.vue'
import Ver from 'pages/roles/ver-rol.vue'
import { parseFecha } from 'stores/utils'

const columns = [
  { name: 'nombre', label: 'Nombre', field: 'nombre', sortable: true },
  { name: 'descripcion', label: 'descripcion', field: 'descripcion' },
  { name: 'jerarquia', label: 'jerarquia', field: 'jerarquia' },
  { name: 'fecha_registro', label: 'registro', field: 'fecha_registro' },
  { name: 'permisos', label: 'permisos', field: 'permisos' },
  { name: 'opt', label: '', field: 'opt' },
]

export default {
  components:{ Registrar,Ver },
  setup () {
    const rolesService = new RolesService();
    const rows = ref([]);
    const loading = ref(false)
    const filter = ref('')
    const refRegistrar = ref()
    const refVer = ref()

    const listar = async () => {
      loading.value = true;
      rows.value = [];
      let res = await rolesService.roles(true).then(u=>u).catch(e=>e)
      if(res.roles) rows.value = res.roles;
      loading.value = false;
    }

    const registrar = () => refRegistrar.value.open();
    const actualizar = (item) => refRegistrar.value.open(item.nombre);
    const visualizar = (item) => refVer.value.open(item.nombre);

    onMounted(()=>{
      listar()
    })

    return {
      columns,
      rows,
      loading,
      filter,
      refRegistrar,
      refVer,
      listar,
      registrar,
      actualizar,
      visualizar,
      parseFecha,
    }
  }
}
</script>
