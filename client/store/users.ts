import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import {
  list as apiUserList,
  create as apiUserCreate,
  update as apiUserUpdate,
  destroy as apiUserDestroy,
} from '../api/user';
import { User } from '../api/types/user';
import type { RootState } from './index';

export const namespace = 'users';

export class UsersState {
  users: Array<User> = [];
}

export const state = (): UsersState => new UsersState();

export const getters: GetterTree<UsersState, RootState> = {
  users(state: UsersState): Array<User> {
    return state.users;
  },
};

export const MutationType = {
  SET_USERS: 'setUsers',
}

export const mutations: MutationTree<UsersState> = {
  [MutationType.SET_USERS]: (state: UsersState, { users }: { users: Array<User> }) => {
    state.users = users;
  },
};

export const actionType = {
  LIST_USERS: 'listUsers',
  CREATE_USER: 'createUser',
  UPDATE_USER: 'updateUser',
  DESTROY_USER: 'destroyUser',
}

export const actions: ActionTree<UsersState, RootState> = {
  [actionType.LIST_USERS]: async ({ commit, state, rootGetters, dispatch }) => {
    const { users } = await apiUserList(this.$http);
    commit('setUsers', { users });
  },

  [actionType.CREATE_USER]: async ({ commit, state, rootGetters, dispatch }, { user }) => {
    await apiUserCreate(this.$http, { name: user.name, email: user.email });
    dispatch('listUsers');
  },

  [actionType.UPDATE_USER]: async ({ commit, state, rootGetters, dispatch }, { user }) => {
    await apiUserUpdate(this.$http, {
      id: user.id,
      name: user.name,
      email: user.email,
    });
    dispatch('listUsers');
  },

  [actionType.DESTROY_USER]: async ({ commit, state, rootGetters, dispatch }, { id }) {
    await apiUserDestroy(this.$http, { id });
    dispatch('listUsers');
  },
};
