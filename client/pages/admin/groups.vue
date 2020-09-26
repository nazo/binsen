<template>
  <v-container
    page>
    <v-layout>
      <v-flex xs12>
        <v-app-bar flat color="white">
          <v-toolbar-title>Groups</v-toolbar-title>
          <v-divider class="mx-2" inset vertical />
          <v-spacer/>
          <v-dialog v-model="dialog" max-width="500px">
            <v-btn slot="activator" color="primary" dark class="mb-2">New Group</v-btn>
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>
              <v-card-text>
                <v-container grid-list-md>
                  <v-layout wrap>
                    <v-flex xs12 sm6 md4>
                      <v-text-field v-model="editedItem.name" :error-messages="errors.collect('name')" :counter="10" required v-validate="'required|max:10'" data-vv-name="name" label="Name"/>
                    </v-flex>
                  </v-layout>
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer/>
                <v-btn color="blue darken-1" flat @click.native="close">Cancel</v-btn>
                <v-btn color="blue darken-1" flat @click.native="save">Save</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-app-bar>
        <v-data-table :headers="headers" :items="groups" hide-actions class="elevation-1" >
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
import { Group } from '../../api/types/group';

const GroupsModule = namespace('groups');

@Component({
  middleware: ['authenticated'],
})
export default class extends Vue {
  @GroupsModule.Getter('groups')
  groups: any;

  @GroupsModule.Action('listGroups')
  listGroups: any;

  @GroupsModule.Action('destroyGroup')
  destroyGroup: any;

  @GroupsModule.Action('createGroup')
  createGroup: any;

  @GroupsModule.Action('updateGroup')
  updateGroup: any;

  get formTitle() {
    return this.editedIndex === -1 ? 'New Item' : 'Edit Item';
  }

  dialog: Boolean = false;
  deleteDialog: Boolean = false;
  deleteDialogItem: Group | null = null;
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
    await store.dispatch('groups/listGroups');
  }

  initialize() {
    this.listGroups();
  }

  editItem(item: Group) {
    this.editedIndex = this.groups.indexOf(item);
    this.editedItem = Object.assign({}, item);
    this.dialog = true;
  }

  deleteItem() {
    const item = this.deleteDialogItem;
    if (item != null) {
      this.destroyGroup({ id: item.id });
    }
  }

  showDeleteDialog(item: Group) {
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
      this.updateGroup({ group: this.editedItem });
    } else {
      this.createGroup({ group: this.editedItem });
    }
    this.close();
  }
}
</script>
