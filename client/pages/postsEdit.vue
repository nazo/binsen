
<script lang="ts">
import marked from 'marked';
import { Store } from 'vuex';
import { Route } from 'vue-router';
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
    CFooter: CFooter,
  },
})
export default class extends Vue {
  @Getter('loggedUser')
  loggedUser: User | null = null;

  @Getter('currentWorkspace')
  currentWorkspace: Workspace | null = null;

  @PostsModule.Getter('currentPost')
  currentPost: Post | null = null;

  @PostsModule.Action('createPost')
  createPost: any;

  @PostsModule.Action('updatePost')
  updatePost: any;

  @PostsModule.Action('getPost')
  getPost: any;

  title: string = '';
  body: string = '';

  clipped = true;
  drawer = true;
  fixed = false;
  miniVariant = false;

  get markedBody() {
    return marked(this.$data.body);
  }

  get pageTitle() {
    if (this.currentPost === null) {
      return 'new post';
    }
    return 'edit post';
  }

  saveDraft() {}

  async submitPost() {
    if (this.currentPost === null) {
      if (this.currentWorkspace !== null) {
        const response = await this.createPost({
          workspaceId: this.currentWorkspace.id,
          title: this.$data.title,
          body: this.$data.body,
        });
        this.$router.push('/posts/' + response.post.id);
      }
    } else {
      const response = await this.updatePost({
        id: this.currentPost.id,
        title: this.$data.title,
        body: this.$data.body,
      });
      this.$router.push('/posts/' + response.post.id);
    }
  }

  mounted() {
    if (this.currentPost !== null) {
      this.title = this.currentPost.title;
      this.body = this.currentPost.body;
    }
  }

  async fetch({ store, query, redirect }: { store: Store<any>, query: Route['query'], redirect(path: string, query?: Route['query']): void }) {
    if (!('id' in query)) {
      return;
    }
    try {
      await store.dispatch('posts/getPost', { id: query.id });
    } catch (e) {
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
