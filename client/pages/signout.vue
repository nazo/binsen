<template>
  <v-layout/>
</template>

<script lang="ts">
import { signout as apiSignout } from '../api/signout';
import { defineComponent, onMounted } from '@nuxtjs/composition-api';
import { NuxtHTTPInstance } from '@nuxt/http';

export default defineComponent({
  layout: 'simple',
  async setup(_props, { root }) {
    onMounted(async () => {
      await apiSignout(root.$http);
      root.$router.replace('/signin');
    });
    return {};
  }
});
</script>
