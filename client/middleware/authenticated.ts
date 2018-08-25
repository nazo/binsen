import { get as getMe } from '../api/me';

export default async ({ $axios, store }) => {
  if (store.getters.isAuthenticated) {
    return;
  }
  const { user } = await getMe($axios);
  store.commit('setUser', { user });
};
