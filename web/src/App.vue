<template>
  <router-view v-on:click="checkClickSession()" />
</template>

<script>
import session from './shared/session'

export default {
  name: 'MyLayout',
  components:{},

  setup () {

    return {
      checkClickSession: session.setup().checkClickSession,
    }

  }
}
</script>
