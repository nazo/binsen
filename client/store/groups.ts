import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators';
import type { RootState } from './index';
import { list as apiGroupList, create as apiGroupCreate, update as apiGroupUpdate, destroy as apiGroupDestroy } from '~/api/group';
import { Group } from '~/api/types/group';

export const namespace = 'groups';

export class GroupsState {
  groups: Group[] = [];
}

export const state = (): GroupsState => new GroupsState();

export const getters: GetterTree<GroupsState, RootState> = {
  groups (state: GroupsState) {
    return state.groups;
  }
};

export const MutationType = {
  SET_GROUPS: 'setGroups'
};

export const mutations: MutationTree<GroupsState> = {
  [MutationType.SET_GROUPS] (state: GroupsState, { groups }: { groups: Group[] }) {
    state.groups = groups || null;
  }
};

export const actionType = {
  LIST_GROUPS: 'listGroups',
  CREATE_GROUP: 'createGroup',
  UPDATE_GROUP: 'updateGroup',
  DESTROY_GROUP: 'destroyGroup'
};

export const actions: ActionTree<GroupsState, RootState> = {
  async [actionType.LIST_GROUPS] ({ commit }: ActionContext<GroupsState, RootState>) {
    try {
      const { groups } = await apiGroupList(this.$http);
      commit('setGroups', { groups });
    } catch (e) {
      // TODO
    }
  },

  async [actionType.CREATE_GROUP] ({ dispatch }: ActionContext<GroupsState, RootState>, { group }: { group: Group }) {
    try {
      await apiGroupCreate(this.$http, { name: group.name });
      dispatch('listGroups');
    } catch (e) {
      // TODO
    }
  },

  async [actionType.UPDATE_GROUP] ({ dispatch }: ActionContext<GroupsState, RootState>, { group }: { group: Group }) {
    try {
      await apiGroupUpdate(this.$http, { id: group.id, name: group.name });
      dispatch('listGroups');
    } catch (e) {
      // TODO
    }
  },

  async [actionType.DESTROY_GROUP] ({ dispatch }: ActionContext<GroupsState, RootState>, { id }: { id: number }) {
    try {
      await apiGroupDestroy(this.$http, { id });
      dispatch('listGroups');
    } catch (e) {
      // TODO
    }
  }
};
