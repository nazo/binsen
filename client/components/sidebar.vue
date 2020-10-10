<template>
  <div>
    <v-navigation-drawer permanent fixed clipped app >
      <v-app-bar flat class="transparent">
        <v-list class="pa-0" v-if="loggedUser">
          <v-list-item avatar>
            <v-list-item-avatar>
              <img :src="loggedUser.image_url">
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title>{{ loggedUser.email }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-app-bar>
      <v-app-bar flat class="transparent">
        <v-toolbar-title slot="activator">
          <span v-text="currentWorkspaceName"/>
          <v-icon dark>icons.mdiMenuDown</v-icon>
        </v-toolbar-title>
        <v-list>
          <v-list-item v-for="workspace in workspaces" :key="workspace.id" @click="changeWorkspace(workspace.id)" >
            <v-list-item-title v-text="workspace.name"/>
          </v-list-item>
        </v-list>
      </v-app-bar>
      <v-list>
        <v-list-item router :to="item.to" :key="i" v-for="(item, i) in items" exact >
          <v-list-item-action>
            <v-icon v-html="item.icon"/>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title"/>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script lang="ts">
import { reactive, computed, ref, shallowReadonly, defineComponent } from '@nuxtjs/composition-api';
import { TranslateResult } from 'vue-i18n';
import { User } from '~/api/types/user';
import { Workspace } from '~/api/types/workspace';
import { actionType as RootAction } from '~/store';
import {
  mdiHome,
  mdiPencil,
  mdiOfficeBuilding,
  mdiAccountMultiple,
  mdiAccountCircleOutline,
  mdiExitToApp,
  mdiMenuDown,
} from '@mdi/js';

export default defineComponent({
  setup(props, { root }) {
    const loggedUser = ref<User | null>(null);
    const workspaces = reactive<Workspace[]>([]);
    const currentWorkspace = ref<Workspace | null>(null);
    const items = shallowReadonly([
      { icon: mdiHome, title: root.$t('menu.home'), to: '/' },
      { icon: mdiPencil, title: root.$t('menu.newpost'), to: '/postsEdit' },
      {
        icon: mdiOfficeBuilding,
        title: root.$t('menu.workspacesadmin'),
        to: '/admin/workspaces',
      },
      {
        icon: mdiAccountMultiple,
        title: root.$t('menu.usersadmin'),
        to: '/admin/users',
      },
      {
        icon: mdiAccountCircleOutline,
        title: root.$t('menu.groupsadmin'),
        to: '/admin/groups',
      },
      {
        icon: mdiExitToApp,
        title: root.$t('menu.signout'),
        to: '/signout',
      },
    ]);
    const icons = shallowReadonly({
      mdiMenuDown
    });
    const currentWorkspaceName = computed(() => currentWorkspace.value?.name ?? '');

    function changeWorkspace(workspaceId: number) {
      root.$store.dispatch(RootAction.SET_WORKSPACE, { workspaceId });
    }

    return {
      loggedUser,
      workspaces,
      currentWorkspace,
      items,
      currentWorkspaceName,
      changeWorkspace,
      icons,
    }
  }
})
</script>
