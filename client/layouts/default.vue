<template>
  <v-app>
    <c-sidebar />
    <v-app-bar fixed app clipped-left />
    <v-main>
      <nuxt />
    </v-main>
    <c-footer />
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api';
import CFooter from '~/components/footer.vue';
import CSidebar from '~/components/sidebar.vue';

export default defineComponent({
  components: {
    CSidebar,
    CFooter
  }
});
</script>
