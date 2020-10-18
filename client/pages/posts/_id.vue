<template>
  <v-container fluid>
    <v-layout column>
      <v-flex v-if="currentPost" xs12>
        <v-layout align-space-between justify-space-between column fill-height>
          <v-flex xs12>
            <h3 class="display-3">
              {{ currentPost.title }}
            </h3>
            <div v-html="markedBody" />
          </v-flex>
          <v-flex xs12>
            <nuxt-link :to="editUrl">
              <v-btn color="teal" text value="recent">
                <span>Edit</span>
                <v-icon>{{ editIcon }}</v-icon>
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
import { reactive, computed, ref, defineComponent, useFetch, useContext } from '@nuxtjs/composition-api';
import { mdiPencil } from '@mdi/js';
import { User } from '~/api/types/user';
import { Post } from '~/api/types/post';
import { actionType as PostsAction, namespace as PostsNamespace } from '~/store/posts';

export default defineComponent({
  middleware: ['authenticated', 'workspaces'],
  setup (props, { root }) {
    const { store, redirect, route, params } = useContext();

    const currentPost = ref<Post | null>(null);
    const markedBody = computed(() => marked(currentPost.value !== null ? currentPost.value.body : ''));
    const editUrl = computed(() => `/postsEdit?id=${currentPost.value?.id}`);
    const editIcon = ref(mdiPencil);

    useFetch(async () => {
      try {
        await store.dispatch(`${PostsNamespace}/${PostsAction.GET_POST}`, {
          id: params.value.id
        });
        currentPost.value = store.getters[`${PostsNamespace}/currentPost`];
      } catch (e) {
        redirect('/');
      }
    });

    function validate () {
      return /^\d+$/.test(params.value.id);
    }

    return {
      currentPost,
      markedBody,
      editUrl,
      editIcon
    };
  }
});
</script>
