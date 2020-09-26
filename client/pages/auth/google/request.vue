<template>
  <v-layout/>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';
import { auth as apiGoogleAuth } from '../../../api/auth/google';

@Component({
  layout: 'simple',
})
export default class extends Vue {
  async mounted() {
    try {
      const { redirectUri } = await apiGoogleAuth(this.$axios);
      window.location.href = redirectUri;
    } catch (e) {
      this.$router.replace('/signin');
    }
  }
}
</script>
