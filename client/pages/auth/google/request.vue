<template>
  <v-layout/>
</template>

<script lang="ts">
import { defineComponent, useFetch } from '@nuxtjs/composition-api';
import { auth as apiGoogleAuth } from '../../../api/auth/google';

export default defineComponent({
  layout: 'simple',
  async setup(_props, { root }) {
    useFetch(async ({ $http }) => {
      try {
        const { redirectUri } = await apiGoogleAuth($http);
        window.location.href = redirectUri;
      } catch (e) {
        root.$router.replace('/signin');
      }
    });
    return {};
  }
});
</script>
