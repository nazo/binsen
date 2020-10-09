import { GetterTree, MutationTree, ActionTree, ActionContext, Commit, Dispatch } from 'vuex';
import {
  list as apiWorkspaceList,
  create as apiWorkspaceCreate,
} from '~/api/workspace';
import { Workspace } from '~/api/types/workspace';
import type { RootState } from './index';

export const namespace = 'workspaces';

export class WorkspacesState {
  workspaces: Workspace[] = [];
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
  [MutationType.SET_WORKSPACES]: function (state: WorkspacesState, { workspaces }: { workspaces: Workspace[] }) {
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
  [actionType.LIST_WORKSPACES]: async function ({ commit }: { commit: Commit }) {
    const { workspaces } = await apiWorkspaceList(this.$http);
    commit('setWorkspaces', { workspaces });
  },

  [actionType.CREATE_WORKSPACE]: async function (
    { dispatch }: { dispatch: Dispatch },
    { workspace }: { workspace: Workspace }
  ) {
    await apiWorkspaceCreate(this.$http, { name: workspace.name });
    dispatch('listWorkspaces');
  },

  [actionType.UPDATE_WORKSPACE]: async function (
    { dispatch }: { dispatch: Dispatch },
    { workspace }: { workspace: Workspace }
  ) {
    dispatch('listWorkspaces');
  },

  [actionType.DESTROY_WORKSPACE]: async function ({ dispatch }: { dispatch: Dispatch }, { id }: { id: number }) {
    dispatch('listWorkspaces');
  },
};
