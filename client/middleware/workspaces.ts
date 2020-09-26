import { NuxtAxiosInstance } from '@nuxtjs/axios';
import { Store } from 'vuex';
import { list as apiWorkspaceList } from '../api/workspace';

export default async ({ $axios, store }: { $axios: NuxtAxiosInstance, store: Store<any> }) => {
  if (store.getters.workspaces !== null) {
    return;
  }
  const { workspaces } = await apiWorkspaceList($axios);
  store.commit('setWorkspaces', { workspaces });
};
