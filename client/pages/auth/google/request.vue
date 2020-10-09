<template>
  <v-layout/>
</template>

<script lang="ts">
import { defineComponent, onMounted } from '@nuxtjs/composition-api';
import { auth as apiGoogleAuth } from '../../../api/auth/google';

export default defineComponent({
  layout: 'simple',
  async setup(_props, { root }) {
    onMounted(async () => {
      try {
        const { redirectUri } = await apiGoogleAuth(root.$http);
        window.location.href = redirectUri;
      } catch (e) {
        root.$router.replace('/signin');
      }
    });

    return {};
  }
});
</script>
