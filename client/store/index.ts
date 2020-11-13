import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import { callback as apiGoogleCallback } from '~/api/auth/google';
import { list as apiWorkspaceList } from '~/api/workspace';
import { User } from '~/api/types/user';
import { Workspace } from '~/api/types/workspace';

export class RootState {
  user: User | null = null;
  currentWorkspaceId: number | null = null;
  workspaces: { [key: number]: Workspace } | null = null;
  locales: string[] = ['en', 'ja'];
  locale: string = 'en';
}

export const state = (): RootState => new RootState();

export const getters: GetterTree<RootState, RootState> = {
  isAuthenticated (state: RootState): boolean {
    return !!state.user;
  },

  loggedUser (state: RootState): User | null {
    return state.user;
  },

  currentWorkspace (state: RootState): Workspace | null {
    if (state.currentWorkspaceId === null) {
      return null;
    }
    if (state.workspaces === null) {
      return null;
    }
    if (!(state.currentWorkspaceId in state.workspaces)) {
      return null;
    }
    return state.workspaces[state.currentWorkspaceId];
  },

  workspaces (state: RootState): { [key: number]: Workspace } | null {
    return state.workspaces;
  }
};

export const MutationType = {
  SET_LANG: 'setLang',
  SET_USER: 'setUser',
  SET_CURRENT_WORKSPACE: 'setCurrentWorkspace',
  SET_WORKSPACES: 'setWorkspaces'
};

export const mutations: MutationTree<RootState> = {
  [MutationType.SET_LANG] (state: RootState, locale: string) {
    if (state.locales.includes(locale)) {
      state.locale = locale;
    }
  },

  [MutationType.SET_USER] (state: RootState, { user }: { user: User }) {
    state.user = user || null;
  },

  [MutationType.SET_CURRENT_WORKSPACE] (state: RootState, { workspaceId }: { workspaceId: number }) {
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

  [MutationType.SET_WORKSPACES] (state: RootState, { workspaces }: { workspaces: Workspace[] }) {
    if (workspaces === null) {
      return;
    }
    const workspacesById: { [key: number]: Workspace } = {};
    workspaces.forEach((workspace: Workspace) => {
      workspacesById[workspace.id] = workspace;
    });
    state.workspaces = workspacesById;
    if (state.currentWorkspaceId === null) {
      state.currentWorkspaceId = workspaces[0].id;
    }
  }
};

export const actionType = {
  GET_WORKSPACES: 'getWorkspaces',
  SET_WORKSPACE: 'setWorkspace',
  LOGIN_FROM_GOOGLE_CALLBACK: 'loginFromGoogleCallback'
};

export const actions: ActionTree<RootState, RootState> = {
  async [actionType.GET_WORKSPACES] ({ commit }: ActionContext<RootState, RootState>) {
    try {
      const { workspaces } = await apiWorkspaceList(this.$http);
      commit('setWorkspaces', { workspaces });
    } catch (e) {
      // TODO
    }
  },

  [actionType.SET_WORKSPACE] ({ commit }: ActionContext<RootState, RootState>, { workspaceId }: { workspaceId: number }) {
    commit('setCurrentWorkspace', { workspaceId });
  },

  async [actionType.LOGIN_FROM_GOOGLE_CALLBACK] ({ commit }: ActionContext<RootState, RootState>, { code, state }: { code: string; state: string }) {
    try {
      const { user } = await apiGoogleCallback(this.$http, { code, state });
      commit('setUser', { user });
    } catch (e) {
      // TODO
    }
  }
};
