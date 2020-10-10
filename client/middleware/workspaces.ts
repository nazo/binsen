import { NuxtHTTPInstance } from '@nuxt/http';
import { Store } from 'vuex';
import { list as apiWorkspaceList } from '~/api/workspace';
import { RootState } from '~/store';
import { defineNuxtMiddleware } from '@nuxtjs/composition-api';

export default defineNuxtMiddleware(async ({ $http, store }: { $http: NuxtHTTPInstance, store: Store<RootState> }) => {
  if (store.getters.workspaces !== null) {
    return;
  }
  const { workspaces } = await apiWorkspaceList($http);
  store.commit('setWorkspaces', { workspaces });
});
