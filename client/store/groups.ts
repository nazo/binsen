import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import {
  list as apiGroupList,
  create as apiGroupCreate,
  update as apiGroupUpdate,
  destroy as apiGroupDestroy,
} from '../api/group';

export class State {
  groups: Array<object> = [];
}

export const state = (): State => new State();

export const mutations: MutationTree<State> = {
  setGroups(state: State, { groups }) {
    state.groups = groups || null;
  },
};

export const getters: GetterTree<State, any> = {
  groups(state: State) {
    return state.groups;
  },
};

export const actions: ActionTree<State, any> = {
  async listGroups({ commit, state, rootGetters, dispatch }) {
    const { groups } = await apiGroupList(this.$axios);
    commit('setGroups', { groups });
  },
  async createGroup({ commit, state, rootGetters, dispatch }, { group }) {
    await apiGroupCreate(this.$axios, { name: group.name });
    dispatch('listGroups');
  },
  async updateGroup({ commit, state, rootGetters, dispatch }, { group }) {
    await apiGroupUpdate(this.$axios, { id: group.id, name: group.name });
    dispatch('listGroups');
  },
  async destroyGroup({ commit, state, rootGetters, dispatch }, { id }) {
    await apiGroupDestroy(this.$axios, { id });
    dispatch('listGroups');
  },
};
