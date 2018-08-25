import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import { callback as apiGoogleCallback } from '../api/auth/google';
import { list as apiWorkspaceList } from '../api/workspace';

export class State {
  user: object | null = null;
  currentWorkspaceId: number | null = null;
  workspaces: object | null = null;
  locales: Array<string> = ['en', 'ja'];
  locale: string = 'en';
}

export const state = (): State => new State();

export const mutations: MutationTree<State> = {
  setLang(state: State, locale: string) {
    if (state.locales.indexOf(locale) !== -1) {
      state.locale = locale;
    }
  },
  setUser(state: State, { user }) {
    state.user = user || null;
  },
  setCurrentWorkspace(state: State, { workspaceId }) {
    if (state.workspaces === null) {
      state.currentWorkspaceId = null;
      return;
    }
    if (!(workspaceId in state.workspaces)) {
      state.currentWorkspaceId = null;
      return;
    }
    state.currentWorkspaceId = workspaceId;
  },
  setWorkspaces(state: State, { workspaces }) {
    if (workspaces === null) {
      return;
    }
    const workspacesById = {};
    workspaces.forEach((workspace) => {
      workspacesById[workspace.id] = workspace;
    });
    state.workspaces = workspacesById;
    if (state.currentWorkspaceId === null) {
      state.currentWorkspaceId = workspaces[0].id;
    }
  },
};

export const getters: GetterTree<State, any> = {
  isAuthenticated(state: State) {
    return !!state.user;
  },
  loggedUser(state: State) {
    return state.user;
  },
  currentWorkspace(state: State) {
    if (state.currentWorkspaceId === null) {
      return null;
    }
    if (!(state.currentWorkspaceId in state.workspaces)) {
      return null;
    }
    return state.workspaces[state.currentWorkspaceId];
  },
  workspaces(state: State) {
    return state.workspaces;
  },
};

export const actions: ActionTree<State, any> = {
  async getWorkspaces({ commit }) {
    const { workspaces } = await apiWorkspaceList(this.$axios);
    commit('setWorkspaces', { workspaces });
  },
  setWorkspace({ commit }, { workspaceId }) {
    commit('setCurrentWorkspace', { workspaceId });
  },
  async loginFromGoogleCallback({ commit }, { code, state }) {
    const { user } = await apiGoogleCallback(this.$axios, code, state);
    commit('setUser', { user });
  },
};
