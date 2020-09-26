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
import { Store } from 'vuex';
import { Route } from 'vue-router';
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';
import { State, Getter, Action, Mutation, namespace } from 'vuex-class';
import { Post } from '../api/types/post';

const PostsModule = namespace('posts');

@Component({
  middleware: ['authenticated', 'workspaces'],
})
export default class extends Vue {
  @PostsModule.Getter('posts')
  posts!: Array<Post> | null;

  page: number = 1;

  async fetch({ store, params }: { store: Store<any>, params: Route['params'] }) {
    await store.dispatch('getWorkspaces');
    await store.dispatch('posts/listPosts', { page: 1 });
  }

  showPost(post: Post): void {
    this.$router.push(`/posts/${post.id}`);
  }
}
</script>
