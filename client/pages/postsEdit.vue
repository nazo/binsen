
<script lang="ts">
import marked from 'marked';
import { Store } from 'vuex';
import { Route } from 'vue-router';
import { reactive, computed, Ref, UnwrapRef, defineComponent } from '@nuxtjs/composition-api';
import CFooter from '~/components/footer.vue';
import { User } from '~/api/types/user';
import { Post } from '~/api/types/post';
import { Workspace } from '~/api/types/workspace';
import { actionType as PostsAction, namespace as PostsNamespace } from '~/store/posts';

export default defineComponent({
  layout: 'clean',
  middleware: ['authenticated', 'workspaces'],
  components: {
    CFooter: CFooter,
  },
  setup(props, { root }) {
    type State = {
      loggedUser: User | null,
      currentWorkspace: Workspace | null,
      title: string,
      body: string,
      clipped: boolean,
      drawer: boolean,
      fixed: boolean,
      miniVariant: boolean,
      markedBody: UnwrapRef<String>,
      pageTitle: UnwrapRef<String>,
    };
    const state: State = reactive({
      loggedUser = null,
      currentWorkspace = null,
      title: '',
      body: '',
      clipped: true,
      drawer: true,
      fixed: false,
      miniVariant: false,
      markedBody: computed(() => marked(state.body)),
      pageTitle: computed(() => root.$store.getters[`${PostsNamespace}/currentPost`] ? 'edit post' : 'new post'),
    });

    function saveDraft() {}

    async function submitPost() {
      const currentPost = root.$store.getters[`${PostsNamespace}/currentPost`];
      const currentWorkspace = root.$store.getters[`${PostsNamespace}/currentWorkspace`];
      if (currentPost === null) {
        if (currentWorkspace !== null) {
          const response = await this.createPost({
            workspaceId: currentWorkspace.id,
            title: state.title,
            body: state.body,
          });
          root.$router.push('/posts/' + response.post.id);
        }
      } else {
        const response = await this.updatePost({
          id: currentPost.id,
          title: state.title,
          body: state.body,
        });
        root.$router.push('/posts/' + response.post.id);
      }
    }

    const currentPost = root.$store.getters[`${PostsNamespace}/currentPost`];
    if (currentPost !== null) {
      state.title = currentPost.title;
      state.body = currentPost.body;
    }

    async function fetch({ store, query, redirect }: { store: Store<any>, query: Route['query'], redirect(path: string, query?: Route['query']): void }) {
      if (!('id' in query)) {
        return;
      }
      try {
        await store.dispatch('posts/getPost', { id: query.id });
      } catch (e) {
        redirect('/');
      }
    }

    function goBack() {
      root.$router.replace('/');
    }

    return {
      state,
      saveDraft,
      submitPost,
      goBack,
    }
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
