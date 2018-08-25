import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import {
  list as apiUserList,
  create as apiUserCreate,
  update as apiUserUpdate,
  destroy as apiUserDestroy,
} from '../api/user';

export class State {
  users: Array<object> = null;
}

export const state = (): State => new State();

export const mutations: MutationTree<State> = {
  setUsers(state: State, { users }) {
    state.users = users || null;
  },
};

export const getters: GetterTree<State, any> = {
  users(state: State) {
    return state.users;
  },
};

export const actions: ActionTree<State, any> = {
  async listUsers({ commit, state, rootGetters, dispatch }) {
    const { users } = await apiUserList(this.$axios);
    commit('setUsers', { users });
  },
  async createUser({ commit, state, rootGetters, dispatch }, { user }) {
    await apiUserCreate(this.$axios, { name: user.name, email: user.email });
    dispatch('listUsers');
  },
  async updateUser({ commit, state, rootGetters, dispatch }, { user }) {
    await apiUserUpdate(this.$axios, { id: user.id, name: user.name, email: user.email });
    dispatch('listUsers');
  },
  async destroyUser({ commit, state, rootGetters, dispatch }, { id }) {
    await apiUserDestroy(this.$axios, { id });
    dispatch('listUsers');
  },
};
