import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import { Mutation } from 'vuex-module-decorators';
import {
  list as apiPostList,
  get as apiPostGet,
  create as apiPostCreate,
  update as apiPostUpdate,
} from '~/api/post';
import { Post } from '~/api/types/post';
import type { RootState } from './index';

export const namespace = 'posts';

export interface PostsState {
  currentPost: Post | null,
  posts: Post[],
}

export const state = (): PostsState => ({
  currentPost: null,
  posts: []
});

export const getters: GetterTree<PostsState, RootState> = {
  posts(state: PostsState): Post[] {
    return state.posts;
  },
  currentPost(state: PostsState): Post | null {
    return state.currentPost;
  },
};

export const MutationType = {
  SET_CURRENT_POST: 'setCurrentPost',
  SET_POSTS: 'setPosts',
}

export const mutations: MutationTree<PostsState> = {
  [MutationType.SET_CURRENT_POST]: function (state: PostsState, { post }: { post: Post }) {
    state.currentPost = post || null;
  },

  [MutationType.SET_POSTS]: function (state: PostsState, { posts }: { posts: Post[] }) {
    state.posts = posts || null;
  },
};

export const actionType = {
  GET_POST: 'getPost',
  LIST_POSTS: 'listPosts',
  CREATE_POST: 'createPost',
  UPDATE_POST: 'updatePost',
}

export const actions: ActionTree<PostsState, RootState> = {
  [actionType.GET_POST]: async function ({ commit }: ActionContext<PostsState, RootState>, { id }: { id: number }) {
    const { post } = await apiPostGet(this.$http, { id });
    commit('setCurrentPost', { post });
  },

  [actionType.LIST_POSTS]: async function ({ commit, rootGetters }: ActionContext<PostsState, RootState>, { page }: { page: number }) {
    const workspace = rootGetters.currentWorkspace;
    if (!workspace) {
      return;
    }
    const { posts } = await apiPostList(this.$http, {
      workspace_id: workspace.id,
      page,
    });
    commit('setPosts', { posts });
  },

  [actionType.CREATE_POST]: function ({}: ActionContext<PostsState, RootState>, { workspaceId, title, body }: { workspaceId: number, title: string, body: string }) {
    return apiPostCreate(this.$http, {
      workspace_id: workspaceId,
      title,
      body,
    });
  },

  [actionType.UPDATE_POST]: function ({}: ActionContext<PostsState, RootState>, { id, title, body }: { id: number, title: string, body: string }) {
    return apiPostUpdate(this.$http, {
      id,
      title,
      body,
    });
  },
};
