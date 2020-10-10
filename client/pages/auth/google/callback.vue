<template>
  <v-layout/>
</template>

<script lang="ts">
import { defineComponent, onMounted, useContext } from '@nuxtjs/composition-api';

export default defineComponent({
  layout: 'simple',
  
  setup(_props, { root }) {
    const { store, query } = useContext();

    onMounted(async () => {
      try {
        await store.dispatch('loginFromGoogleCallback', {
          code: query.value.code,
          state: query.value.state,
        });
        if (store.getters.isAuthenticated) {
          root.$router.replace('/');
        }
      } catch (e) {
        root.$router.replace('/signin');
      }
    });

    return {};
  }
});
</script>
