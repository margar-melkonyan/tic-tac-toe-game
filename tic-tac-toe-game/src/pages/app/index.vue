<template>
  <v-container>
    <v-row class="pa-0 mt-6">
      <v-col cols="6">
        <div class="text-h5">
          Комнаты
        </div>
      </v-col>
      <v-col
        class="d-flex justify-end"
        cols="6"
      >
        <v-btn @click="enterRoom">
          Создать комнату
        </v-btn>
      </v-col>
    </v-row>
    <v-divider class="my-6" />
    <div>
      <v-row
        v-for="(currentGroup, key) in groupedRooms"
        :key="`group-${key}`"
        class="mt-2"
      >
        <v-col
          v-for="(room, room_key) in currentGroup"
          :key="`room-${key}-${room_key}`"
        >
          <RoomCard />
        </v-col>
      </v-row>
    </div>
    <div />
  </v-container>
  <v-dialog
    v-model="dialogCreateRoom"
    origin="left center"
    max-width="500"
    persistent
  >
    <v-card>
      <v-card-title class="mx-2 my-2">
        Создать комнату
      </v-card-title>
      <v-divider />
      <v-card-text>
        <v-text-field variant="outlined">
          <template #label>
            Название комнаты
          </template>
        </v-text-field>
        <v-row class="mb-2">
          <v-col class="py-0 d-flex align-center">
            Приватная комната
          </v-col>
          <v-col class="py-0 d-flex justify-end">
            <v-switch
              v-model="isPrivateRoom"
              inset
              base-color="red"
              color="green"
              hide-details
              @click="changeRoomPrivacy"
            />
          </v-col>
        </v-row>
        <v-text-field variant="outlined" v-if="isPrivateRoom">
          <template #label>
            Пароль
          </template>
        </v-text-field>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-col class="d-flex justify-start py-0">
          <v-btn @click="closeRoomCreationDialog">
            Закрыть
          </v-btn>
        </v-col>
        <v-col class="d-flex justify-end py-0">
          <v-btn>Создать</v-btn>
        </v-col>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import {ref} from "vue"

const dialogCreateRoom: boolean = ref(false)
const isPrivateRoom: boolean = ref(false)
const rooms = ref([
  {
    id: "13123",
    player_in: 1,
    max_player: 2,
    title: 'Комната №1'
  },
  {
    id: "13123",
    player_in: 2,
    max_player: 2,
    title: 'Комната №2'
  },
  {
    id: "13123",
    player_in: 2,
    max_player: 2,
    title: 'Комната №3'
  },
  {
    id: "13123",
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

const groupChunkSize = 3;
const groupedRooms = computed(() => {
  const result = [];
  for (let i = 0; i < rooms.value.length; i += groupChunkSize) {
    result.push(rooms.value.slice(i, i + groupChunkSize));
  }
  return result;
});

function enterRoom() {
  dialogCreateRoom.value = true
}

function closeRoomCreationDialog() {
  isPrivateRoom.value = false
  dialogCreateRoom.value = false
}

function changeRoomPrivacy() {
  isPrivateRoom.value = !isPrivateRoom.value
}
</script>
