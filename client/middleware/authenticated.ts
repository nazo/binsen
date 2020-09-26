import { NuxtAxiosInstance } from '@nuxtjs/axios';
import { Store } from 'vuex';
import { get as getMe } from '../api/me';
import { State } from '../store';

export default async ({ $axios, store }: { $axios: NuxtAxiosInstance, store: Store<any> }) => {
  if (store.getters.isAuthenticated) {
    return;
  }
  const { user } = await getMe($axios);
  store.commit('setUser', { user });
};
