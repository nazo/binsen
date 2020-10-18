<template>
  <v-container page>
    <v-layout>
      <v-flex xs12>
        <v-app-bar flat color="white">
          <v-toolbar-title>Workspaces</v-toolbar-title>
          <v-divider class="mx-2" inset vertical />
          <v-spacer />
          <v-btn color="primary" dark class="mb-2" @click.stop="newItem">
            New Workspace
          </v-btn>
          <v-dialog v-model="dialog" max-width="500px">
            <v-card>
              <ValidationObserver v-slot="{ invalid }">
                <form @submit.prevent="save">
                  <v-card-title>
                    <span class="headline">{{ formTitle }}</span>
                  </v-card-title>
                  <v-card-text>
                    <v-container grid-list-md>
                      <v-layout wrap>
                        <v-flex xs12 sm6 md4>
                          <ValidationProvider v-slot="{ errors }" name="name" rules="required|max:10">
                            <v-text-field v-model="editedItem.name" :counter="10" required data-vv-name="name" label="Name" />
                            <div>{{ errors[0] }}</div>
                          </ValidationProvider>
                        </v-flex>
                      </v-layout>
                    </v-container>
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer />
                    <v-btn color="blue darken-1" text @click.native="close">
                      Cancel
                    </v-btn>
                    <v-btn type="submit" color="blue darken-1" text :disabled="invalid">
                      Save
                    </v-btn>
                  </v-card-actions>
                </form>
              </ValidationObserver>
            </v-card>
          </v-dialog>
        </v-app-bar>
        <v-data-table :headers="headers" :items="workspaces" hide-default-footer class="elevation-1">
          <template slot="item" slot-scope="props">
            <tr>
              <td class="text-xs-right">
                {{ props.item.id }}
              </td>
              <td>{{ props.item.name }}</td>
              <td class="justify-center layout px-0">
                <v-icon small class="mr-2" @click="editItem(props.item)">
                  edit
                </v-icon>
                <v-icon small class="mr-2" @click.stop="showDeleteDialog(props.item)">
                  delete
                </v-icon>
              </td>
            </tr>
          </template>
          <template slot="no-data">
            <v-btn color="primary" @click="initialize">
              Reset
            </v-btn>
          </template>
        </v-data-table>
        <v-dialog v-model="deleteDialog" persistent max-width="290">
          <v-card>
            <v-card-title class="headline">
              Delete
            </v-card-title>
            <v-card-text>Are you sure?</v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn color="green darken-1" flat @click.native="deleteDialog = false">
                Cancel
              </v-btn>
              <v-btn color="green darken-1" flat @click.native="deleteDialog = false" @click.stop="deleteItem()">
                Delete
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { reactive, computed, ref, defineComponent, useFetch, useContext, watch, shallowReadonly } from '@nuxtjs/composition-api';
import { DataTableHeader } from 'vuetify';
import { Workspace } from '~/api/types/workspace';
import { actionType as WorkspacesAction, namespace as WorkspacesNamespace } from '~/store/workspaces';

export default defineComponent({
  middleware: ['authenticated'],

  setup (props, { root }) {
    const { store, redirect, route, params } = useContext();

    const workspaces = ref<Workspace[]>([]);
    const dialog = ref(false);
    const deleteDialog = ref(false);
    const deleteDialogItem = ref<Workspace | null>(null);
    const headers = reactive<DataTableHeader[]>([
      { text: 'ID', value: 'id' },
      { text: 'Name', value: 'name' },
      { text: 'Actions', value: 'name', sortable: false }
    ]);
    const editedIndex = ref(-1);
    const editedItem = ref({
      id: 0,
      name: ''
    });
    const defaultItem = ref({
      id: 0,
      name: ''
    });
    const dictionary = shallowReadonly({
      custom: {
        name: {
          required: () => 'Name can not be empty',
          max: 'The name field may not be greater than 10 characters'
        }
      }
    });
    const formTitle = computed(() => (editedIndex.value === -1 ? 'New Item' : 'Edit Item'));

    const { fetch } = useFetch(async () => {
      await store.dispatch(`${WorkspacesNamespace}/${WorkspacesAction.LIST_WORKSPACES}`);
      workspaces.value = store.getters[`${WorkspacesNamespace}/workspaces`] as Workspace[];
    });

    watch(
      () => dialog,
      (newValue, oldValue) => {
        newValue || close();
      }
    );

    function editItem (item: Workspace) {
      console.log(item);
      editedIndex.value = workspaces.value.indexOf(item);
      editedItem.value = Object.assign({}, item);
      dialog.value = true;
    }

    async function deleteItem () {
      const item = deleteDialogItem.value;
      if (item != null) {
        await store.dispatch(`${WorkspacesNamespace}/${WorkspacesAction.DESTROY_WORKSPACE}`, { id: item.id });
        fetch();
      }
    }

    function showDeleteDialog (item: Workspace) {
      deleteDialog.value = true;
      deleteDialogItem.value = item;
    }

    async function close () {
      dialog.value = false;
      await root.$nextTick();
      editedItem.value = Object.assign({}, defaultItem.value);
      editedIndex.value = -1;
    }

    async function save () {
      if (editedIndex.value > -1) {
        await store.dispatch(`${WorkspacesNamespace}/${WorkspacesAction.UPDATE_WORKSPACE}`, { workspace: editedItem.value });
      } else {
        await store.dispatch(`${WorkspacesNamespace}/${WorkspacesAction.CREATE_WORKSPACE}`, { workspace: editedItem.value });
      }
      fetch();
      close();
    }

    async function initialize () {
      await store.dispatch(`${WorkspacesNamespace}/${WorkspacesAction.LIST_WORKSPACES}`);
    }

    function newItem () {
      editedIndex.value = -1;
      dialog.value = true;
    }

    return {
      workspaces,
      dialog,
      deleteDialog,
      deleteDialogItem,
      headers,
      editedIndex,
      editedItem,
      defaultItem,
      dictionary,
      formTitle,
      newItem,
      editItem,
      deleteItem,
      showDeleteDialog,
      close,
      save,
      initialize
    };
  }
});
</script>
