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
            <v-list-tile
              :key="post.id"
              avatar
              @click="showPost(post)"
            >
              <v-list-tile-avatar>
                <img src="">
              </v-list-tile-avatar>

              <v-list-tile-content>
                <v-list-tile-title v-text="post.title"/>
                <v-list-tile-sub-title v-text="post.body"/>
              </v-list-tile-content>
            </v-list-tile>
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
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';
import { State, Getter, Action, Mutation, namespace } from 'vuex-class';

const PostsModule = namespace('posts');

@Component({
  middleware: ['authenticated', 'workspaces']
})
export default class extends Vue {
  @PostsModule.Getter('posts') posts: any;
  @Getter('loggedUser') loggedUser!: object;

  page: number = 1

  async fetch({ store, params }) {
    await store.dispatch('getWorkspaces');
    await store.dispatch('posts/listPosts', { page: 1 });
  }

  showPost(post: any): void {
    this.$router.push(`/posts/${post.id}`);
  }
}
</script>
