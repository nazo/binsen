<template>
  <v-layout/>
</template>

<script lang="ts">
import { signout as apiSignout } from '../api/signout';
import { Component, Vue } from 'nuxt-property-decorator';

@Component({
  layout: 'simple'
})
export default class extends Vue {
  async mounted() {
    await apiSignout(this.$axios);
    this.$router.replace('/signin');
  }
}
</script>
