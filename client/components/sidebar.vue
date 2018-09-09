<template>
  <div>
    <v-navigation-drawer
      permanent
      fixed
      clipped
      app
    >
      <v-toolbar
        flat
        class="transparent">
        <v-list
          class="pa-0"
          v-if="loggedUser">
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <img :src="loggedUser.image_url">
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>{{ loggedUser.email }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-toolbar>
      <v-toolbar
        flat
        class="transparent">
        <v-toolbar-title slot="activator">
          <span v-text="currentWorkspaceName"/>
          <v-icon dark>arrow_drop_down</v-icon>
        </v-toolbar-title>
        <v-list>
          <v-list-tile
            v-for="workspace in workspaces"
            :key="workspace.id"
            @click="changeWorkspace(workspace.id)"
          >
            <v-list-tile-title v-text="workspace.name"/>
          </v-list-tile>
        </v-list>
      </v-toolbar>
      <v-list>
        <v-list-tile
          router
          :to="item.to"
          :key="i"
          v-for="(item, i) in items"
          exact
        >
          <v-list-tile-action>
            <v-icon v-html="item.icon"/>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title v-text="item.title"/>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';
import { State, Getter, Action, Mutation, namespace } from 'vuex-class';
import { User } from '../api/types/user';
import { Workspace } from '../api/types/workspace';

@Component({})
export default class extends Vue {
  @Getter('loggedUser') loggedUser: User | null
  @Getter('workspaces') workspaces: Array<Workspace>
  @Getter('currentWorkspace') currentWorkspace: Workspace | null
  @Action('setWorkspace') setWorkspace: any

  get currentWorkspaceName() {
    if (this.currentWorkspace === null) {
      return null;
    }
    return this.currentWorkspace.name;
  }

  items = [
    { icon: 'apps', title: this.$root.$t('menu.home'), to: '/' },
    { icon: 'edit', title: this.$root.$t('menu.newpost'), to: '/postsEdit' },
    { icon: 'edit', title: this.$root.$t('menu.workspacesadmin'), to: '/admin/workspaces' },
    { icon: 'people', title: this.$root.$t('menu.usersadmin'), to: '/admin/users' },
    { icon: 'group', title: this.$root.$t('menu.groupsadmin'), to: '/admin/groups' },
    { icon: 'exit_to_app', title: this.$root.$t('menu.signout'), to: '/signout' },
  ]

  changeWorkspace(workspaceId) {
    this.setWorkspace({ workspaceId });
  }
}
</script>
