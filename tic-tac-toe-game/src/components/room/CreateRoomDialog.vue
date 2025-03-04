<template>
  <v-dialog v-model="dialogCreateRoom" origin="left center" max-width="500" persistent>
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
        <v-text-field v-if="isPrivateRoom" variant="outlined" type="password">
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
import { ref } from 'vue';

const dialogCreateRoom = ref(false)
const isPrivateRoom = ref(false)

function closeRoomCreationDialog() {
  isPrivateRoom.value = false
  dialogCreateRoom.value = false
}

function changeRoomPrivacy() {
  isPrivateRoom.value = !isPrivateRoom.value
}

function createRoom() {
  dialogCreateRoom.value = true
}

defineExpose({ createRoom })
</script>
