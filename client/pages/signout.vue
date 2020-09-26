<template>
  <v-layout/>
</template>

<script lang="ts">
import { signout as apiSignout } from '../api/signout';
import { Component, Vue } from 'nuxt-property-decorator';
import { NuxtAxiosInstance } from '@nuxtjs/axios';

@Component({
  layout: 'simple',
})
export default class extends Vue {
  async mounted({ $axios }: { $axios: NuxtAxiosInstance }) {
    await apiSignout($axios);
    this.$router.replace('/signin');
  }
}
</script>
