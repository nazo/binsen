<template>
  <v-container fluid>
    <v-layout column>
      <v-flex xs12>
        <v-layout
          align-space-between
          justify-space-between
          column
          fill-height>
          <v-flex xs12>
            <h3 class="display-3">{{ currentPost.title }}</h3>
            <div v-html="markedBody" />
          </v-flex>
          <v-flex xs12>
            <nuxt-link :to="editUrl">
              <v-btn
                color="teal"
                flat
                value="recent"
              >
                <span>Edit</span>
                <v-icon>edit</v-icon>
              </v-btn>
            </nuxt-link>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import marked from 'marked';
import { Store } from 'vuex';
import { Route } from 'vue-router';
import { reactive, computed, Ref, UnwrapRef, defineComponent, useFetch, useContext } from '@nuxtjs/composition-api';
import { User } from '~/api/types/user';
import { Post } from '~/api/types/post';
import { actionType as PostsAction, namespace as PostsNamespace } from '~/store/posts';

export default defineComponent({
  middleware: ['authenticated', 'workspaces'],
  setup(props, { root }) {
    const { store, redirect, route, params } = useContext();

    type State = {
      loggedUser: User | null,
      markedBody: UnwrapRef<String>,
      editUrl: UnwrapRef<String>,
      currentPost: UnwrapRef<Post>,
    };
    const state: State = reactive({
      loggedUser: null,
      markedBody: computed(() => marked(state.currentPost.body)),
      editUrl: computed(() => `/postsEdit?id=${state.currentPost.id}`),
      currentPost: computed(() => store.getters[`${PostsNamespace}/currentPost`]),
    });

    useFetch(async () => {
      try {
        await store.dispatch(`${PostsNamespace}/${PostsAction.GET_POST}`, { id: params.value.id });
      } catch (e) {
        redirect('/');
      }
    });

    function validate() {
      return /^\d+$/.test(params.value.id);
    }

    return {
      state,
    }
  }
});
</script>
