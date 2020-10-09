<template>
  <v-layout/>
</template>

<script lang="ts">
import { signout as apiSignout } from '../api/signout';
import { defineComponent, useFetch } from '@nuxtjs/composition-api';
import { NuxtHTTPInstance } from '@nuxt/http';

export default defineComponent({
  layout: 'simple',
  async setup(_props, { root }) {
    useFetch(async ({ $http }) => {
      await apiSignout($http);
      root.$router.replace('/signin');
    });
    return {};
  }
});
</script>
