<template>
  <v-container>
    <v-row class="pa-0 mt-6">
      <v-col cols="6">
        <div class="text-h5">
          {{ $t('titles.rooms') }}
        </div>
      </v-col>
      <v-col class="d-flex justify-end" cols="6">
        <v-btn @click="openCreateRoomDialog">
          Создать комнату
        </v-btn>
      </v-col>
    </v-row>
    <v-divider class="my-6" />
    <RoomList :rooms="rooms" />
    <v-divider class="my-6" />
    <CreateRoomDialog ref="newRoomDialog" />
  </v-container>
</template>

<script lang="ts" setup>
import type CreateRoomDialog from "@/components/room/CreateRoomDialog.vue";
import { ref } from "vue"
const rooms = ref([
  {
    id: "1",
    player_in: 1,
    max_player: 2,
    title: 'Комната №1'
  },
  {
    id: "2",
    player_in: 2,
    max_player: 2,
    title: 'Комната №2'
  },
  {
    id: "3",
    player_in: 2,
    max_player: 2,
    title: 'Комната №3'
  },
  {
    id: "4",
    player_in: 2,
    max_player: 2,
    title: 'Комната №4'
  },
  {
    id: "13123",
    player_in: 1,
    max_player: 2,
    title: 'Комната №5'
  },
  {
    id: "13123",
    player_in: 2,
    max_player: 2,
    title: 'Комната №6'
  },
  {
    id: "13123",
    player_in: 2,
    max_player: 2,
    title: 'Комната №7'
  },
  {
    id: "13123",
    player_in: 2,
    max_player: 2,
    title: 'Комната №8'
  },
  {
    id: "13123",
    player_in: 1,
    max_player: 2,
    title: 'Комната №9'
  },
  {
    id: "13123",
    player_in: 2,
    max_player: 2,
    title: 'Комната №10'
  },
  {
    id: "13123",
    player_in: 2,
    max_player: 2,
    title: 'Комната №11'
  },
]);

const newRoomDialog = ref<InstanceType<typeof CreateRoomDialog> | null>(null);

const openCreateRoomDialog = () => {
  if (newRoomDialog.value) {
    newRoomDialog.value.createRoom();
  }
};
</script>
