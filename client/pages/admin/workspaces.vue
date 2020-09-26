<template>
  <v-container page>
    <v-layout>
      <v-flex xs12>
        <v-app-bar flat color="white">
          <ValidationObserver v-slot="{ handleSubmit }">
            <form @submit.prevent="handleSubmit(save)">
              <v-toolbar-title>Workspaces</v-toolbar-title>
              <v-divider class="mx-2" inset vertical />
              <v-spacer/>
              <v-dialog v-model="dialog" max-width="500px">
                <v-btn slot="activator" color="primary" dark class="mb-2">New Workspace</v-btn>
                <v-card>
                  <v-card-title>
                    <span class="headline">{{ formTitle }}</span>
                  </v-card-title>
                  <v-card-text>
                    <v-container grid-list-md>
                      <v-layout wrap>
                        <v-flex xs12 sm6 md4>
                          <ValidationProvider v-slot="{ errors }">
                            <v-text-field v-model="editedItem.name" :error-messages="errors.collect('name')" :counter="10" required v-validate="'required|max:10'" data-vv-name="name" label="Name"/>
                            <div>{{ errors[0] }}</div>
                          </ValidationProvider>
                        </v-flex>
                      </v-layout>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer/>
                    <v-btn color="blue darken-1" flat @click.native="close">Cancel</v-btn>
                    <v-btn color="blue darken-1" flat>Save</v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </form>
          </ValidationObserver>
        </v-app-bar>
        <v-data-table :headers="headers" :items="workspaces" hide-actions class="elevation-1" >
          <template slot="items" slot-scope="props">
            <td class="text-xs-right">{{ props.item.id }}</td>
            <td>{{ props.item.name }}</td>
            <td class="justify-center layout px-0">
              <v-icon small class="mr-2" @click="editItem(props.item)">edit</v-icon>
              <v-icon small class="mr-2" @click.stop="showDeleteDialog(props.item)">delete</v-icon>
            </td>
          </template>
          <template slot="no-data">
            <v-btn color="primary" @click="initialize">Reset</v-btn>
          </template>
        </v-data-table>
        <v-dialog v-model="deleteDialog" persistent max-width="290">
          <v-card>
            <v-card-title class="headline">Delete</v-card-title>
            <v-card-text>Are you sure?</v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn color="green darken-1" flat @click.native="deleteDialog = false">Cancel</v-btn>
              <v-btn color="green darken-1" flat @click.native="deleteDialog = false" @click.stop="deleteItem()">Delete</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Store } from 'vuex';
import { Component, Prop, Emit, Watch, Vue } from 'nuxt-property-decorator';
import { State, Getter, Action, Mutation, namespace } from 'vuex-class';
import { Workspace } from '../../api/types/workspace';

const WorkspacesModule = namespace('workspaces');

@Component({
  middleware: ['authenticated'],
})
export default class extends Vue {
  @WorkspacesModule.Getter('workspaces')
  workspaces: any;

  @WorkspacesModule.Action('listWorkspaces')
  listWorkspaces: any;

  @WorkspacesModule.Action('destroyWorkspace')
  destroyWorkspace: any;

  @WorkspacesModule.Action('createWorkspace')
  createWorkspace: any;

  @WorkspacesModule.Action('updateWorkspace')
  updateWorkspace: any;

  get formTitle() {
    return this.editedIndex === -1 ? 'New Item' : 'Edit Item';
  }

  dialog: Boolean = false;
  deleteDialog: Boolean = false;
  deleteDialogItem: Workspace | null = null;
  headers = [
    { text: 'ID', value: 'id' },
    { text: 'Name', value: 'name' },
    { text: 'Actions', value: 'name', sortable: false },
  ];
  editedIndex = -1;
  editedItem = {
    id: 0,
    name: '',
  };
  defaultItem = {
    id: 0,
    name: '',
  };
  dictionary = {
    custom: {
      name: {
        required: () => 'Name can not be empty',
        max: 'The name field may not be greater than 10 characters',
      },
    },
  };

  @Watch('dialog')
  onDialogChange(newValue: string, oldValue: string): void {
    newValue || this.close();
  }

  async fetch({ store }: { store: Store<any> }) {
    await store.dispatch('workspaces/listWorkspaces');
  }

  initialize() {
    this.listWorkspaces();
  }

  editItem(item: Workspace) {
    this.editedIndex = this.workspaces.indexOf(item);
    this.editedItem = Object.assign({}, item);
    this.dialog = true;
  }

  deleteItem() {
    const item = this.deleteDialogItem;
    if (item != null) {
      this.destroyWorkspace({ id: item.id });
    }
  }

  showDeleteDialog(item: Workspace) {
    this.deleteDialog = true;
    this.deleteDialogItem = item;
  }

  close() {
    this.dialog = false;
    setTimeout(() => {
      this.editedItem = Object.assign({}, this.defaultItem);
      this.editedIndex = -1;
    }, 0);
  }

  async save() {
    if (this.editedIndex > -1) {
      this.updateWorkspace({ workspace: this.editedItem });
    } else {
      this.createWorkspace({ workspace: this.editedItem });
    }
    this.close();
  }
}
</script>
