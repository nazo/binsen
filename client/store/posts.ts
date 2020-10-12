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
  CLEAR_POST: 'clearPost',
  LIST_POSTS: 'listPosts',
  CREATE_POST: 'createPost',
  UPDATE_POST: 'updatePost',
}

export const actions: ActionTree<PostsState, RootState> = {
  [actionType.GET_POST]: async function ({ commit }: ActionContext<PostsState, RootState>, { id }: { id: number }) {
    try {
      const { post } = await apiPostGet(this.$http, { id });
      commit(MutationType.SET_CURRENT_POST, { post });
    } catch(e) {
      // TODO
    }
  },

  [actionType.CLEAR_POST]: async function ({ commit }: ActionContext<PostsState, RootState>) {
    commit(MutationType.SET_CURRENT_POST, { post: null });
  },

  [actionType.LIST_POSTS]: async function ({ commit, rootGetters }: ActionContext<PostsState, RootState>, { page }: { page: number }) {
    try {
      const workspace = rootGetters.currentWorkspace;
      if (!workspace) {
        return;
      }
      const { posts } = await apiPostList(this.$http, {
        workspace_id: workspace.id,
        page,
      });
      commit(MutationType.SET_POSTS, { posts });
    } catch(e) {
      // TODO
    }
  },

  [actionType.CREATE_POST]: function ({}: ActionContext<PostsState, RootState>, { workspaceId, title, body }: { workspaceId: number, title: string, body: string }) {
    try {
      return apiPostCreate(this.$http, {
        workspace_id: workspaceId,
        title,
        body,
      });
    } catch(e) {
      // TODO
    }
  },

  [actionType.UPDATE_POST]: function ({}: ActionContext<PostsState, RootState>, { id, title, body }: { id: number, title: string, body: string }) {
    try {
      return apiPostUpdate(this.$http, {
        id,
        title,
        body,
      });
    } catch(e) {
      // TODO
    }
  },
};
