import { NuxtHTTPInstance } from '@nuxt/http';
import { Store } from 'vuex';
import { list as apiWorkspaceList } from '../api/workspace';

export default async ({ $http, store }: { $http: NuxtHTTPInstance, store: Store<any> }) => {
  if (store.getters.workspaces !== null) {
    return;
  }
  const { workspaces } = await apiWorkspaceList($http);
  store.commit('setWorkspaces', { workspaces });
};
