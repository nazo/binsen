
<script lang="ts">
import marked from 'marked';
import { Store } from 'vuex';
import { Route } from 'vue-router';
import { reactive, computed, Ref, UnwrapRef, defineComponent, useFetch, useContext, onMounted } from '@nuxtjs/composition-api';
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
    const { store, redirect, query } = useContext();

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
      loggedUser: null,
      currentWorkspace: null,
      title: '',
      body: '',
      clipped: true,
      drawer: true,
      fixed: false,
      miniVariant: false,
      markedBody: computed(() => marked(state.body)),
      pageTitle: computed(() => store.getters[`${PostsNamespace}/currentPost`] ? 'edit post' : 'new post'),
    });

    function saveDraft() {}

    async function submitPost() {
      const currentPost = store.getters[`${PostsNamespace}/currentPost`];
      const currentWorkspace = store.getters[`${PostsNamespace}/currentWorkspace`];
      if (currentPost === null) {
        if (currentWorkspace !== null) {
          const response = await store.dispatch(`${PostsNamespace}/${PostsAction.CREATE_POST}`, {
            workspaceId: currentWorkspace.id,
            title: state.title,
            body: state.body,
          });
          root.$router.push('/posts/' + response.post.id);
        }
      } else {
        const response = await store.dispatch(`${PostsNamespace}/${PostsAction.UPDATE_POST}`, {
          id: currentPost.id,
          title: state.title,
          body: state.body,
        });
        root.$router.push('/posts/' + response.post.id);
      }
    }

    onMounted(() => {
      const currentPost = store.getters[`${PostsNamespace}/currentPost`];
      if (currentPost !== null) {
        state.title = currentPost.title;
        state.body = currentPost.body;
      }
    });

    function goBack() {
      root.$router.replace('/');
    }

    useFetch(async () => {
      if (!('id' in query.value)) {
        return;
      }
      try {
        await store.dispatch(`${PostsNamespace}/${PostsAction.GET_POST}`, { id: query.value.id });
      } catch (e) {
        redirect('/');
      }
    });

    return {
      state,
      saveDraft,
      submitPost,
      goBack,
    }
  }
});
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
