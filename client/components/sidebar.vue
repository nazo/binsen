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
          <v-icon dark>arrow_drop_down</v-icon>
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
import { computed, defineComponent, ref, useAsync, useContext, useFetch, useMeta } from '@nuxtjs/composition-api';
import { User } from '~/api/types/user';
import { Workspace } from '~/api/types/workspace';

export default defineComponent({
  @Getter('loggedUser')
  let loggedUser!: User | null;

  @Getter('workspaces')
  workspaces!: Array<Workspace>;

  @Getter('currentWorkspace')
  currentWorkspace!: Workspace | null;

  @Action('setWorkspace')
  setWorkspace: any;

  get currentWorkspaceName() {
    if (this.currentWorkspace === null) {
      return null;
    }
    return this.currentWorkspace.name;
  }

  items = [
    { icon: 'apps', title: this.$root.$t('menu.home'), to: '/' },
    { icon: 'edit', title: this.$root.$t('menu.newpost'), to: '/postsEdit' },
    {
      icon: 'edit',
      title: this.$root.$t('menu.workspacesadmin'),
      to: '/admin/workspaces',
    },
    {
      icon: 'people',
      title: this.$root.$t('menu.usersadmin'),
      to: '/admin/users',
    },
    {
      icon: 'group',
      title: this.$root.$t('menu.groupsadmin'),
      to: '/admin/groups',
    },
    {
      icon: 'exit_to_app',
      title: this.$root.$t('menu.signout'),
      to: '/signout',
    },
  ];

  changeWorkspace(workspaceId: number) {
    this.setWorkspace({ workspaceId });
  }
}
</script>
