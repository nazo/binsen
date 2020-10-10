import { defineNuxtPlugin } from '@nuxtjs/composition-api';
import { Context } from '@nuxt/types/app';

export default defineNuxtPlugin(({ $http, redirect }: Context) => {
  $http.onError(error => {
    const code = error.response && error.response!.status;
    if(code === 401) {
      redirect('/signin');
    }
  })
});
