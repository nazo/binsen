<template>
  <v-layout />
</template>

<script lang="ts">
import { defineComponent, onMounted } from '@nuxtjs/composition-api';
import { NuxtHTTPInstance } from '@nuxt/http';
import { signout as apiSignout } from '../api/signout';

export default defineComponent({
  layout: 'simple',
  setup (_props, { root }) {
    onMounted(async () => {
      await apiSignout(root.$http);
      root.$router.replace('/signin');
    });
    return {};
  }
});
</script>
