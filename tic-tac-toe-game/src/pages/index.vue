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
          {{ $t('titles.rooms') }}
        </div>
      </v-col>
      <v-col
        v-if="auth.user !== null"
        class="d-flex justify-end"
        cols="6"
      >
        <v-btn @click="openCreateRoomDialog">
          Создать комнату
        </v-btn>
      </v-col>
    </v-row>
    <v-divider class="my-6" />
    <RoomList
      :rooms="rooms"
      @open-login-dialog="openLoginDialog"
    />
    <v-divider class="my-6" />
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
