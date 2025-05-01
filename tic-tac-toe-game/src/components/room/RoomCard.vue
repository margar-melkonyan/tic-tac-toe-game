<template>
  <v-card
    color="#7fff94"
    variant="outlined"
  >
    <v-card-title>
      {{ $t('room.title', [props.room.name ]) }}
    </v-card-title>
    <v-card-text>
      {{ $t('room.players', [props.room.player_in, props.room.capacity]) }}
    </v-card-text>
    <v-card-actions>
      <v-row>
        <v-col
          cols="12"
          class="d-flex
          justify-end"
        >
          <v-btn
            variant="outlined"
            :color="props.room.player_in === props.room.capacity ? 'white' : '#ff7fea'"
            density="comfortable"
            :disabled="props.room.player_in === props.room.capacity"
            @click="openRoom"
          >
            {{ $t('enter') }}
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>
  </v-card>
  <v-dialog
    v-model="enterRoom"
    width="500"
    persistent
  >
    <v-card>
      <v-card-title class="mx-2 my-2">
        {{ $t('room.title', [props.room.name]) }}
      </v-card-title>
      <v-divider />
      <v-card-text>
        <v-text-field
          variant="outlined"
          density="compact"
          type="password"
        >
          <template #label>
            {{ $t('room.fields.password') }}
          </template>
        </v-text-field>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-col class="d-flex justify-start py-0">
          <v-btn @click="closeRoom">
            {{ $t('close') }}
          </v-btn>
        </v-col>
        <v-col class="d-flex justify-end py-0">
          <v-btn>
            {{ $t('enter') }}
          </v-btn>
        </v-col>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { ref, defineEmits } from 'vue';
import { useAuthStore } from '@/stores/auth';
const emit = defineEmits([
  "openLoginDialog"
])
const authStore = useAuthStore()
const props = defineProps<{
  room: {
    id: string;
    name: string;
    player_in: number;
    capacity: number;
    is_private: boolean;
    created_at: string;
  },
}>()
const enterRoom = ref(false);
function openRoom() {
  if (authStore.user == null) {
    emit("openLoginDialog")
    return
  }
  enterRoom.value = true
}

function closeRoom() {
  enterRoom.value = false
}
</script>
