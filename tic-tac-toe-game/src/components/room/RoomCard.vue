<template>
  <v-card color="#7fff94" variant="outlined">
    <v-card-title>
      Комната #{{ props.room.id }}
    </v-card-title>
    <v-card-text>
      Игроков: {{ props.room.player_in }}/{{ props.room.max_player }}
    </v-card-text>
    <v-card-actions>
      <v-row>
        <v-col cols="12" class="d-flex justify-end">
          <v-btn variant="outlined" color="#ff7fea" density="comfortable" @click="openRoom">
            Войти
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>
  </v-card>
  <v-dialog v-model="enterRoom" width="500" persistent>
    <v-card>
      <v-card-title class="mx-2 my-2">
        Комната #{{ props.room.id }}
      </v-card-title>
      <v-divider />
      <v-card-text>
        <v-text-field variant="outlined" density="compact" type="password">
          <template #label>
            Пароль
          </template>
        </v-text-field>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-col class="d-flex justify-start py-0">
          <v-btn @click="closeRoom">
            Закрыть
          </v-btn>
        </v-col>
        <v-col class="d-flex justify-end py-0">
          <v-btn>
            Войти
          </v-btn>
        </v-col>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const props = defineProps<{
  room: {
    id: string;
    title: string;
    player_in: number;
    max_player: number;
  },
}>()

const enterRoom = ref(false);

function openRoom() {
  enterRoom.value = true
}

function closeRoom() {
  enterRoom.value = false
}
</script>
