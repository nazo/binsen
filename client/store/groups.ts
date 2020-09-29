import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators';
import {
  list as apiGroupList,
  create as apiGroupCreate,
  update as apiGroupUpdate,
  destroy as apiGroupDestroy,
} from '../api/group';
import type { RootState } from './index';

export const namespace = 'groups';

export class GroupsState {
  groups: Array<object> = [];
}

export const state = (): GroupsState => new GroupsState();

export const getters: GetterTree<GroupsState, RootState> = {
  groups(state: GroupsState) {
    return state.groups;
  },
};

export const MutationType = {
  SET_GROUPS: 'setGroups',
}

export const mutations: MutationTree<GroupsState> = {
  [MutationType.SET_GROUPS]: (state: GroupsState, { groups }) => {
    state.groups = groups || null;
  },
};

export const actionType = {
  LIST_GROUPS: 'listGroups',
  CREATE_GROUP: 'createGroup',
  UPDATE_GROUP: 'updateGroup',
  DESTROY_GROUP: 'destroyGroup',
}

export const actions: ActionTree<GroupsState, RootState> = {
  [actionType.LIST_GROUPS]: async ({ commit, state, rootGetters, dispatch }) => {
    const { groups } = await apiGroupList(this.$http);
    commit('setGroups', { groups });
  },

  [actionType.CREATE_GROUP]: async ({ commit, state, rootGetters, dispatch }, { group }) => {
    await apiGroupCreate(this.$http, { name: group.name });
    dispatch('listGroups');
  },

  [actionType.UPDATE_GROUP]: async ({ commit, state, rootGetters, dispatch }, { group }) => {
    await apiGroupUpdate(this.$http, { id: group.id, name: group.name });
    dispatch('listGroups');
  },

  [actionType.DESTROY_GROUP]: async ({ commit, state, rootGetters, dispatch }, { id }) => {
    await apiGroupDestroy(this.$http, { id });
    dispatch('listGroups');
  },
};
