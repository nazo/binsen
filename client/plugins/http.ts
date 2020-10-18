import { defineNuxtPlugin } from '@nuxtjs/composition-api';
import { Context } from '@nuxt/types/app';
import ky, { NormalizedOptions } from 'ky';

export default defineNuxtPlugin(({ $http, redirect }: Context) => {
  ($http as any)._defaults.credentials = 'include';

  $http.onError((error: ky.HTTPError) => {
    const code = error.response?.status;
    if (code === 401) {
      redirect('/signin');
    }
  });
});
