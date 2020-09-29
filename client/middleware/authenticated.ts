import { NuxtHTTPInstance } from '@nuxt/http';
import { Store } from 'vuex';
import { get as getMe } from '../api/me';
import { RootState } from '../store';

export default async ({ $http, store }: { $http: NuxtHTTPInstance, store: Store<any> }) => {
  if (store.getters.isAuthenticated) {
    return;
  }
  const { user } = await getMe($http);
  store.commit('setUser', { user });
};
