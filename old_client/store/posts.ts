import { GetterTree, MutationTree, ActionTree, ActionContext } from 'vuex';
import {
  list as apiPostList,
  get as apiPostGet,
  create as apiPostCreate,
  update as apiPostUpdate,
} from '../api/post';
import { Post } from '../api/types/post';

export class State {
  currentPost: Post | null = null;
  posts: Array<Post> = [];
}

export const state = (): State => new State();

export const mutations: MutationTree<State> = {
  setCurrentPost(state: State, { post }: { post: Post }) {
    state.currentPost = post || null;
  },
  setPosts(state: State, { posts }: { posts: Array<Post> }) {
    state.posts = posts || null;
  },
};

export const getters: GetterTree<State, any> = {
  posts(state: State): Array<Post> {
    return state.posts;
  },
  currentPost(state: State): Post | null {
    return state.currentPost;
  },
};

export const actions: ActionTree<State, any> = {
  async getPost({ commit }, { id }) {
    const { post } = await apiPostGet(this.$axios, id);
    commit('setCurrentPost', { post });
  },
  async listPosts({ commit, state, rootGetters, dispatch }, { page }) {
    const workspace = rootGetters.currentWorkspace;
    if (!workspace) {
      return;
    }
    const { posts } = await apiPostList(this.$axios, {
      workspace_id: workspace.id,
      page,
    });
    commit('setPosts', { posts });
  },
  createPost({ commit }, { workspaceId, title, body }) {
    return apiPostCreate(this.$axios, {
      workspace_id: workspaceId,
      title,
      body,
    });
  },
  updatePost({ commit }, { id, title, body }) {
    return apiPostUpdate(this.$axios, {
      id,
      title,
      body,
    });
  },
};
