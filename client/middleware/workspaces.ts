import { list as apiWorkspaceList } from '../api/workspace';

export default async ({ $axios, store }) => {
  if (store.getters.workspaces !== null) {
    return;
  }
  const { workspaces } = await apiWorkspaceList($axios);
  store.commit('setWorkspaces', { workspaces });
};
