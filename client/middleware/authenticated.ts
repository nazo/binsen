import { NuxtHTTPInstance } from '@nuxt/http';
import { get as getMe } from '~/api/me';
import { RootState } from '~/store';
import { defineNuxtMiddleware } from '@nuxtjs/composition-api';
import { Store } from 'vuex';

export default defineNuxtMiddleware(async ({ $http, store }: { $http: NuxtHTTPInstance, store: Store<RootState> }) => {
  if (store.getters.isAuthenticated) {
    return;
  }
  const { user } = await getMe($http);
  store.commit('setUser', { user });
});
