<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm8 md4>
      <v-card class="elevation-12">
        <v-app-bar dark color="primary">
          <v-toolbar-title>{{ $t('pages.signin.title') }}</v-toolbar-title>
        </v-app-bar>
        <v-card-actions>
          <v-btn color="primary" @click="signinByGoogle">
            {{ $t('pages.signin.google_signin') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api';

export default defineComponent({
  layout: 'simple',
  setup (_props, { root }) {
    async function signinByGoogle () {
      root.$router.push('/auth/google/request');
    }
    return {
      signinByGoogle
    };
  }
});
</script>
