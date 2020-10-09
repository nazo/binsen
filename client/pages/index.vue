<template>
  <v-container fluid>
    <v-layout
      column
      justify-center
      align-center>
      <v-flex
        xs12
        sm8
        md6>
        <v-list
          two-line
          v-if="posts">
          <template v-for="post in posts">
            <v-list-item
              :key="post.id"
              avatar
              @click="showPost(post)"
            >
              <v-list-item-avatar>
                <img src="">
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title v-text="post.title"/>
                <v-list-item-sub-title v-text="post.body"/>
              </v-list-item-content>
            </v-list-item>
          </template>
        </v-list>
        <v-pagination
          v-model="page"
          :length="6"
        />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { reactive, computed, defineComponent, ref, useAsync, useContext, useFetch, useMeta } from '@nuxtjs/composition-api';
import { Post } from '~/api/types/post';
import { actionType as RootAction } from '~/store';
import { actionType as PostsAction, namespace as PostsNamespace } from '~/store/posts';

export default defineComponent({
  setup(props, { root }) {
    type State = {
      page: Number;
    };
    const state: State = reactive({
      page: 1,
    });

    useFetch(async () => {
      await root.$store.dispatch(RootAction.GET_WORKSPACES);
      await root.$store.dispatch(`${PostsNamespace}/${PostsAction.LIST_POSTS}`, { page: 1 });
    });

    function showPost(post: Post): void {
      root.$router.push(`/posts/${post.id}`);
    }

    function posts(): Post[] {
      return root.$store.getters[`${PostsNamespace}/posts`];
    }

    return {
      state,
      showPost,
    };
  }
});
</script>
