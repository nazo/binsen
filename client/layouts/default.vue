<template>
  <v-app>
    <c-sidebar/>
    <v-toolbar
      fixed
      app
      clipped-left/>
    <v-content>
      <nuxt />
    </v-content>
    <c-footer/>
  </v-app>
</template>

<script lang="ts">
import CFooter from '../components/footer';
import CSidebar from '../components/sidebar';
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';

@Component({
  components: {
    CSidebar,
    CFooter,
  }
})
export default class extends Vue {
  clipped = true
  fixed = false
  miniVariant = false
}
</script>
