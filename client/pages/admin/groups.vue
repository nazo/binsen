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
import { reactive, computed, ref, defineComponent, useFetch, useContext, watch, shallowReadonly } from '@nuxtjs/composition-api';
import { Group } from '~/api/types/group';
import { actionType as GroupsAction, namespace as GroupsNamespace } from '~/store/groups';

export default defineComponent({
  middleware: ['authenticated'],

  setup(props, { root }) {
    const { store, redirect, route, params } = useContext();

    const groups = reactive<Group[]>([]);
    const dialog = ref(false);
    const deleteDialog = ref(false);
    const deleteDialogItem = ref<Group | null>(null);
    const headers = shallowReadonly([
      { text: 'ID', value: 'id' },
      { text: 'Name', value: 'name' },
      { text: 'Actions', value: 'name', sortable: false },
    ]);
    const editedIndex = ref(-1);
    const editedItem = ref({
      id: 0,
      name: '',
    });
    const defaultItem = ref({
      id: 0,
      name: '',
    });
    const dictionary = shallowReadonly({
      custom: {
        name: {
          required: () => 'Name can not be empty',
          max: 'The name field may not be greater than 10 characters',
        },
      },
    });
    const formTitle = computed(() => editedIndex.value === -1 ? 'New Item' : 'Edit Item');

    useFetch(async () => {
      await store.dispatch(`${GroupsNamespace}/${GroupsAction.LIST_GROUPS}`);
    });

    watch(
      () => dialog,
      (newValue, oldValue) => {
        newValue || close();
      }
    );

    function editItem(item: Group) {
      editedIndex.value = groups.indexOf(item);
      editedItem.value = Object.assign({}, item);
      dialog.value = true;
    };

    function deleteItem() {
      const item = deleteDialogItem.value;
      if (item != null) {
        store.dispatch(`${GroupsNamespace}/${GroupsAction.DESTROY_GROUP}`, { id: item.id });
        store.dispatch(`${GroupsNamespace}/${GroupsAction.LIST_GROUPS}`);
      }
    }

    function showDeleteDialog(item: Group) {
      deleteDialog.value = true;
      deleteDialogItem.value = item;
    }

    async function close() {
      dialog.value = false;
      await root.$nextTick();
      editedItem.value = Object.assign({}, defaultItem.value);
      editedIndex.value = -1;
    }

    async function save() {
      if (editedIndex.value > -1) {
        store.dispatch(`${GroupsNamespace}/${GroupsAction.UPDATE_GROUP}`, { group: editedItem.value });
      } else {
        store.dispatch(`${GroupsNamespace}/${GroupsAction.CREATE_GROUP}`, { group: editedItem.value });
      }
      close();
    }
    
    async function initialize() {
      await store.dispatch(`${GroupsNamespace}/${GroupsAction.LIST_GROUPS}`);
    }

    return {
      groups,
      dialog,
      deleteDialog,
      deleteDialogItem,
      headers,
      editedIndex,
      editedItem,
      defaultItem,
      dictionary,
      formTitle,
      editItem,
      deleteItem,
      showDeleteDialog,
      close,
      save,
      initialize,
    };
  }
});
</script>
