<template>
  <v-layout/>
</template>

<script lang="ts">
import { signout as apiSignout } from '../api/signout';
import { Component, Vue } from 'nuxt-property-decorator';
import { NuxtHTTPInstance } from '@nuxt/http';

@Component({
  layout: 'simple',
})
export default class extends Vue {
  async mounted({ $http }: { $http: NuxtHTTPInstance }) {
    await apiSignout($http);
    this.$router.replace('/signin');
  }
}
</script>
