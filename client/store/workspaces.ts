import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import {
  list as apiWorkspaceList,
  create as apiWorkspaceCreate,
} from '../api/workspace';
import type { RootState } from './index';

export const namespace = 'workspaces';

export class WorkspacesState {
  workspaces: Array<object> = [];
}

export const state = (): WorkspacesState => new WorkspacesState();

export const getters: GetterTree<WorkspacesState, RootState> = {
  workspaces(state: WorkspacesState) {
    return state.workspaces;
  },
};

export const MutationType = {
  SET_WORKSPACES: 'setUsers',
}

export const mutations: MutationTree<WorkspacesState> = {
  [MutationType.SET_WORKSPACES]: (state: WorkspacesState, { workspaces }) => {
    state.workspaces = workspaces || null;
  },
};

export const actionType = {
  LIST_WORKSPACES: 'listWorkspaces',
  CREATE_WORKSPACE: 'createWorkspace',
  UPDATE_WORKSPACE: 'updateWorkspace',
  DESTROY_WORKSPACE: 'destroyWorkspace',
}

export const actions: ActionTree<WorkspacesState, RootState> = {
  [actionType.LIST_WORKSPACES]: async ({ commit, state, rootGetters, dispatch }) => {
    const { workspaces } = await apiWorkspaceList(this.$http);
    commit('setWorkspaces', { workspaces });
  },

  [actionType.CREATE_WORKSPACE]: async (
    { commit, state, rootGetters, dispatch },
    { workspace }
  ) => {
    await apiWorkspaceCreate(this.$http, { name: workspace.name });
    dispatch('listWorkspaces');
  },

  [actionType.UPDATE_WORKSPACE]: async (
    { commit, state, rootGetters, dispatch },
    { workspace }
  ) => {
    dispatch('listWorkspaces');
  },

  [actionType.DESTROY_WORKSPACE]: async ({ commit, state, rootGetters, dispatch }, { id }) => {
    dispatch('listWorkspaces');
  },
};
