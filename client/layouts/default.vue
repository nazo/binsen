<template>
  <v-app>
    <c-sidebar/>
    <v-app-bar fixed app clipped-left/>
    <v-main>
      <nuxt />
    </v-main>
    <c-footer/>
  </v-app>
</template>

<script lang="ts">
import CFooter from '~/components/footer.vue';
import CSidebar from '~/components/sidebar.vue';
import { reactive, computed, Ref, UnwrapRef, defineComponent, useFetch, useContext, watch } from '@nuxtjs/composition-api';

export default defineComponent({
  components: {
    CSidebar,
    CFooter,
  },

  setup(props, { root }) {
    const state = reactive({
      clipped: true,
      fixed: false,
      miniVariant: false,
    });

    return { state };
  }
});
</script>
