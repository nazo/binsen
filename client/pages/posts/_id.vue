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
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';
import { State, Getter, Action, Mutation, namespace } from 'vuex-class';
import { User } from '../../api/types/user';
import { Post } from '../../api/types/post';

const PostsModule = namespace('posts');

@Component({
  middleware: ['authenticated', 'workspaces'],
})
export default class extends Vue {
  @PostsModule.Getter('currentPost')
  currentPost!: Post;

  @Getter('loggedUser')
  loggedUser!: User | null;

  get editUrl() {
    return `/postsEdit?id=${this.currentPost.id}`;
  }

  get markedBody() {
    return marked(this.currentPost.body);
  }

  async fetch({ store, params, redirect }: { store: Store<any>, params: Route['params'], redirect(path: string, query?: Route['query']): void }) {
    try {
      await store.dispatch('posts/getPost', { id: params.id });
    } catch (e) {
      redirect('/');
    }
  }

  validate({ params }: { params: Route['params'] }) {
    return /^\d+$/.test(params.id);
  }
}
</script>
