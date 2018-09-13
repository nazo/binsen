import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import {
  list as apiWorkspaceList,
  create as apiWorkspaceCreate,
} from '../api/workspace';

export class State {
  workspaces: Array<object> = [];
}

export const state = (): State => new State();

export const mutations: MutationTree<State> = {
  setWorkspaces(state: State, { workspaces }) {
    state.workspaces = workspaces || null;
  },
};

export const getters: GetterTree<State, any> = {
  workspaces(state: State) {
    return state.workspaces;
  },
};

export const actions: ActionTree<State, any> = {
  async listWorkspaces({ commit, state, rootGetters, dispatch }) {
    const { workspaces } = await apiWorkspaceList(this.$axios);
    commit('setWorkspaces', { workspaces });
  },
  async createWorkspace(
    { commit, state, rootGetters, dispatch },
    { workspace }
  ) {
    await apiWorkspaceCreate(this.$axios, { name: workspace.name });
    dispatch('listWorkspaces');
  },
  async updateWorkspace(
    { commit, state, rootGetters, dispatch },
    { workspace }
  ) {
    dispatch('listWorkspaces');
  },
  async destroyWorkspace({ commit, state, rootGetters, dispatch }, { id }) {
    dispatch('listWorkspaces');
  },
};
