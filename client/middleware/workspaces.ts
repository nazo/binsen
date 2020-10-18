import { NuxtHTTPInstance } from '@nuxt/http';
import { Store } from 'vuex';
import { defineNuxtMiddleware } from '@nuxtjs/composition-api';
import { list as apiWorkspaceList } from '~/api/workspace';
import { RootState } from '~/store';

export default defineNuxtMiddleware(async ({ $http, store }: { $http: NuxtHTTPInstance; store: Store<RootState> }) => {
  if (store.getters.workspaces !== null) {
    return;
  }
  const { workspaces } = await apiWorkspaceList($http);
  store.commit('setWorkspaces', { workspaces });
});
