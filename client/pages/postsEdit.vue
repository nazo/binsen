<template>
  <v-app>
    <v-toolbar
      fixed
      app
      :clipped-left="clipped">
      <v-btn
        icon
        class="hidden-xs-only"
        @click="goBack">
        <v-icon>arrow_back</v-icon>
      </v-btn>
      <v-toolbar-title v-text="pageTitle"/>
    </v-toolbar>
    <v-content>
      <v-container
        fluid
        fill-height>
        <v-container
          fluid
          fill-height>
          <v-layout column>
            <v-flex xs12>
              <v-layout
                fill-height
                justify-space-between
                column>
                <div>
                  <v-text-field
                    label="Title"
                    placeholder="Title"
                    v-model="title"
                  />
                </div>
                <div class="editor-area">
                  <v-layout
                    align-space-between
                    justify-space-between
                    row
                    fill-height>
                    <div>
                      <textarea
                        class="editor-margined"
                        v-model="body" />
                    </div>
                    <div>
                      <v-card
                        dark
                        class="editor-margined editor-preview"
                        color="secondary">
                        <v-card-text
                          class="px-0"
                          v-html="markedBody"/>
                      </v-card>
                    </div>
                  </v-layout>
                </div>
                <div>
                  <v-btn @click="saveDraft">Draft</v-btn>
                  <v-btn
                    color="success"
                    @click="submitPost">Post</v-btn>
                </div>
              </v-layout>
            </v-flex>
          </v-layout>
        </v-container>
      </v-container>
    </v-content>
    <c-footer/>
  </v-app>
</template>

<script lang="ts">
import marked from 'marked';
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';
import { State, Getter, Action, Mutation, namespace } from 'vuex-class';
import CFooter from '../components/footer.vue';
import { User } from '../api/types/user';
import { Post } from '../api/types/post';
import { Workspace } from '../api/types/workspace';

const PostsModule = namespace('posts');

@Component({
  layout: 'clean',
  middleware: ['authenticated', 'workspaces'],
  components: {
    CFooter: CFooter
  }
})
export default class extends Vue {
  @Getter('loggedUser') loggedUser: User | null = null;
  @Getter('currentWorkspace') currentWorkspace: Workspace | null = null;
  @PostsModule.Getter('currentPost') currentPost: Post | null = null;
  @PostsModule.Action('createPost') createPost: any;
  @PostsModule.Action('updatePost') updatePost: any;
  @PostsModule.Action('getPost') getPost: any;

  title: string = ''
  body: string = ''

  clipped = true
  drawer = true
  fixed = false
  miniVariant = false

  get markedBody() {
    return marked(this.$data.body);
  }

  get pageTitle() {
    if (this.currentPost === null) {
      return 'new post';
    }
    return 'edit post';
  }

  saveDraft() {
  }

  async submitPost() {
    if (this.currentPost === null) {
      if (this.currentWorkspace !== null) {
        const response = await this.createPost({ workspaceId: this.currentWorkspace.id, title: this.$data.title, body: this.$data.body });
        this.$router.push('/posts/' + response.post.id);
      }
    } else {
      const response = await this.updatePost({ id: this.currentPost.id, title: this.$data.title, body: this.$data.body });
      this.$router.push('/posts/' + response.post.id);
    }
  }

  mounted() {
    if (this.currentPost !== null) {
      this.title = this.currentPost.title;
      this.body = this.currentPost.body;
    }
  }

  async fetch({ store, query, redirect }) {
    if (!('id' in query)) {
      return;
    }
    try {
      await store.dispatch('posts/getPost', { id: query.id });
    } catch(e) {
      redirect('/');
    }
  }

  goBack() {
    this.$router.replace('/');
  }
}
</script>

<style scoped>
.editor-area {
  height: 100%;
  width: 100%;
}
.editor-area * {
  height: 100%;
  width: 100%;
}
.editor-area textarea {
  height: 100%;
  width: 100%;
  border: 1px solid #000;
}
.editor-margined {
  margin: 5px;
  word-break: break-all;
}
.editor-preview {
  overflow: scroll;
}
</style>
