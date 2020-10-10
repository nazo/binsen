<template>
  <v-app>
    <v-app-bar
      fixed
      app
      :clipped-left="clipped">
      <v-btn
        icon
        class="hidden-xs-only"
        @click="goBack">
        <v-icon>{{ icons.mdiArrowLeftCircle }}</v-icon>
      </v-btn>
      <v-toolbar-title v-text="pageTitle"/>
    </v-app-bar>
    <v-main>
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
    </v-main>
    <c-footer/>
  </v-app>
</template>

<script lang="ts">
import marked from 'marked';
import { Store } from 'vuex';
import { Route } from 'vue-router';
import { shallowReadonly, reactive, computed, ref, defineComponent, useFetch, useContext, onMounted } from '@nuxtjs/composition-api';
import CFooter from '~/components/footer.vue';
import { User } from '~/api/types/user';
import { Post } from '~/api/types/post';
import { Workspace } from '~/api/types/workspace';
import { mdiArrowLeftCircle } from '@mdi/js';
import { actionType as PostsAction, namespace as PostsNamespace } from '~/store/posts';

export default defineComponent({
  layout: 'clean',
  middleware: ['authenticated', 'workspaces'],
  components: {
    CFooter: CFooter,
  },
  setup(props, { root }) {
    const { store, redirect, query } = useContext();

    const title = ref('');
    const body = ref('');
    const clipped = ref(true);
    const markedBody = computed(() => marked(body.value));
    const pageTitle = computed(() => store.getters[`${PostsNamespace}/currentPost`] ? 'edit post' : 'new post');
    const icons = shallowReadonly({
      mdiArrowLeftCircle
    });

    function saveDraft() {}

    async function submitPost() {
      const currentPost = store.getters[`${PostsNamespace}/currentPost`];
      const currentWorkspace = store.getters[`${PostsNamespace}/currentWorkspace`];
      if (currentPost === null) {
        if (currentWorkspace !== null) {
          const response = await store.dispatch(`${PostsNamespace}/${PostsAction.CREATE_POST}`, {
            workspaceId: currentWorkspace.id,
            title: title.value,
            body: body.value,
          });
          root.$router.push('/posts/' + response.post.id);
        }
      } else {
        const response = await store.dispatch(`${PostsNamespace}/${PostsAction.UPDATE_POST}`, {
          id: currentPost.id,
          title: title.value,
          body: body.value,
        });
        root.$router.push('/posts/' + response.post.id);
      }
    }

    onMounted(() => {
      const currentPost = store.getters[`${PostsNamespace}/currentPost`];
      if (currentPost !== null) {
        title.value = currentPost.title;
        body.value = currentPost.body;
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
      title,
      body,
      clipped,
      markedBody,
      pageTitle,
      saveDraft,
      submitPost,
      goBack,
      icons,
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
