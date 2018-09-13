<template>
  <v-layout/>
</template>

<script lang="ts">
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';
import { State, Getter, Action, Mutation, namespace } from 'vuex-class';

@Component({
  layout: 'simple',
})
export default class extends Vue {
  @Action('loginFromGoogleCallback')
  loginFromGoogleCallback: any;

  async mounted() {
    try {
      await this.loginFromGoogleCallback({
        code: this.$route.query.code,
        state: this.$route.query.state,
      });
      if (this.$store.getters.isAuthenticated) {
        this.$router.replace('/');
      }
    } catch (e) {
      this.$router.replace('/signin');
    }
  }
}
</script>
