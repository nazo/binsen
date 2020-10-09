<template>
  <v-layout/>
</template>

<script lang="ts">
import { defineComponent, useFetch } from '@nuxtjs/composition-api';

export default defineComponent({
  layout: 'simple',
  async setup(_props, { root }) {
    try {
      await root.$store.dispatch('loginFromGoogleCallback', {
        code: root.$route.query.code,
        state: root.$route.query.state,
      });
      if (root.$store.getters.isAuthenticated) {
        root.$router.replace('/');
      }
    } catch (e) {
      root.$router.replace('/signin');
    }

    return {};
  }
});
</script>
