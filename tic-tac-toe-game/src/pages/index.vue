<template>
  <v-container>
    <v-row>
      <v-col
        v-if="auth.user == null"
        class="d-flex justify-end"
      >
        <v-btn
          @click="openLoginDialog"
        >
          {{ $t('menu.sign_in') }}
        </v-btn>
      </v-col>
      <v-col
        v-else
        class="d-flex justify-end"
      >
        <div>
          <span>
            {{ auth.user?.name }} / {{ auth.user?.email }}
          </span>
          <v-btn
            class="ml-4"
            @click="auth.signOut()"
          >
            {{ $t('menu.sign_out') }}
          </v-btn>
        </div>
      </v-col>
    </v-row>
    <v-row class="pa-0 mt-6">
      <v-col cols="6">
        <div class="text-h5">
          {{ $t('rooms.title') }}
        </div>
      </v-col>
      <v-col
        v-if="auth.user !== null"
        class="d-flex justify-end"
        cols="6"
      >
        <v-btn @click="openCreateRoomDialog">
          {{ $t('rooms.create') }}
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-tabs
          color="red"
          fixed-tabs
        >
          <v-tab fixed>
            {{ $t('rooms.all') }}
          </v-tab>
          <v-tab fixed>
            {{ $t('rooms.my') }}
          </v-tab>
        </v-tabs>
      </v-col>
    </v-row>
    <v-divider class="my-8" />
    <RoomList
      class="my-8"
      :rooms="rooms"
      @open-login-dialog="openLoginDialog"
    />
    <v-divider class="my-4" />
    <CreateRoomDialog
      ref="newRoomDialog"
      @close-room-create-dialog="fetchRooms"
    />
    <FormDialog
      :login-dialog="loginDialog"
      @close="loginDialog = false"
    />
  </v-container>
</template>

<script lang="ts" setup>
import type CreateRoomDialog from "@/components/room/CreateRoomDialog.vue";
import { ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import axios from "axios";
const auth = useAuthStore();
const { proxy } = getCurrentInstance();
const newRoomDialog = ref<InstanceType<typeof CreateRoomDialog> | null>(null);
const loginDialog = ref(false)
const rooms = ref([]);
const apiRooms = proxy.$api.rooms;
auth.currentUser()
const openLoginDialog = () => {
  loginDialog.value = true;
}
const openCreateRoomDialog = () => {
  if (newRoomDialog.value) {
    newRoomDialog.value.createRoom();
  }
};
const fetchRooms = () => {
  axios.get(apiRooms.urls.rooms())
    .then(({data}) => {
      rooms.value = data.data ?? []
    })
}
fetchRooms()
</script>
