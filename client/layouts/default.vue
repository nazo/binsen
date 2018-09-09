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
import CFooter from '../components/footer.vue';
import CSidebar from '../components/sidebar.vue';
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
