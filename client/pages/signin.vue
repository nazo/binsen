<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm8 md4>
      <v-card class="elevation-12">
        <v-app-bar dark color="primary">
          <v-toolbar-title>{{ $t('pages.signin.title') }}</v-toolbar-title>
        </v-app-bar>
        <v-card-actions>
          <v-btn color="primary" @click="signinByGoogle">{{ $t('pages.signin.google_signin') }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';

@Component({
  layout: 'simple',
})
export default class extends Vue {
  async signinByGoogle() {
    this.$router.push('/auth/google/request');
  }
}
</script>
