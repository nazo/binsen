<template>
  <v-container
    page>
    <v-layout>
      <v-flex xs12>
        <v-app-bar flat color="white">
          <v-toolbar-title>Users</v-toolbar-title>
          <v-divider class="mx-2" inset vertical />
          <v-spacer/>
          <v-dialog v-model="dialog" max-width="500px">
            <v-btn slot="activator" color="primary" dark class="mb-2">New User</v-btn>
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
                    <v-flex xs12 sm6 md4>
                      <v-text-field v-model="editedItem.email" :error-messages="errors.collect('email')" required v-validate="'required|email'" data-vv-name="email" label="email"/>
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
        <v-data-table :headers="headers" :items="users" hide-actions class="elevation-1" >
          <template slot="items" slot-scope="props">
            <td class="text-xs-right">{{ props.item.id }}</td>
            <td>{{ props.item.name }}</td>
            <td>{{ props.item.email }}</td>
            <td>{{ props.item.users }}</td>
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
import { reactive, computed, Ref, UnwrapRef, defineComponent, useFetch, useContext, watch } from '@nuxtjs/composition-api';
import { User } from '~/api/types/user';
import { actionType as UsersAction, namespace as UsersNamespace } from '~/store/users';

export default defineComponent({
  middleware: ['authenticated'],

  setup(props, { root }) {
    const { store, redirect, route, params } = useContext();

    type State = {
      users: User[],
      formTitle: UnwrapRef<string>,
      dialog: boolean,
      deleteDialog: boolean,
      deleteDialogItem: User | null,
      headers: { text: string, value: string, sortable?: boolean }[],
      editedIndex: number,
      editedItem: { id: number, name: string },
      defaultItem: { id: number, name: string },
      dictionary: { custom: { name: { required: () => string, max: string } } },
    };
    const state: State = reactive({
      users: [],
      formTitle: computed(() => state.editedIndex === -1 ? 'New Item' : 'Edit Item'),
      dialog: false,
      deleteDialog: false,
      deleteDialogItem: null,
      headers: [
        { text: 'ID', value: 'id' },
        { text: 'Name', value: 'name' },
        { text: 'Actions', value: 'name', sortable: false },
      ],
      editedIndex: -1,
      editedItem: {
        id: 0,
        name: '',
      },
      defaultItem: {
        id: 0,
        name: '',
      },
      dictionary: {
        custom: {
          name: {
            required: () => 'Name can not be empty',
            max: 'The name field may not be greater than 10 characters',
          },
        },
      },
    });

    useFetch(async () => {
      await store.dispatch(`${UsersNamespace}/${UsersAction.LIST_USERS}`);
    });

    watch(
      () => state.dialog,
      (newValue, oldValue) => {
        newValue || close();
      }
    );

    function editItem(item: User) {
      state.editedIndex = state.users.indexOf(item);
      state.editedItem = Object.assign({}, item);
      state.dialog = true;
    };

    function deleteItem() {
      const item = state.deleteDialogItem;
      if (item != null) {
        store.dispatch(`${UsersNamespace}/${UsersAction.DESTROY_USER}`, { id: item.id });
        store.dispatch(`${UsersNamespace}/${UsersAction.LIST_USERS}`);
      }
    }

    function showDeleteDialog(item: User) {
      state.deleteDialog = true;
      state.deleteDialogItem = item;
    }

    function close() {
      state.dialog = false;
      setTimeout(() => {
        state.editedItem = Object.assign({}, state.defaultItem);
        state.editedIndex = -1;
      }, 0);
    }

    async function save() {
      if (state.editedIndex > -1) {
        store.dispatch(`${UsersNamespace}/${UsersAction.UPDATE_USER}`, { user: state.editedItem });
      } else {
        store.dispatch(`${UsersNamespace}/${UsersAction.CREATE_USER}`, { user: state.editedItem });
      }
      close();
    }

    return {
      state,
      editItem,
      deleteItem,
      showDeleteDialog,
      close,
      save,
    };
  }
});
</script>
