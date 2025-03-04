<template>
  <div>
    <v-row v-for="(currentGroup, key) in groupedRooms" :key="`group-${key}`" class="mt-2">
      <v-col v-for="(room, room_key) in currentGroup" :key="`room-${key}-${room_key}`">
        <RoomCard :room="room" />
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts" setup>
import { computed, defineProps } from 'vue';

const props = defineProps<{
  rooms: Array<{
    id: string;
    title: string;
    player_in: number;
    max_player: number;
  }>;
}>();

const groupChunkSize = 3;
const groupedRooms = computed(() => {
  const result = [];
  for (let i = 0; i < props.rooms.length; i += groupChunkSize) {
    result.push(props.rooms.slice(i, i + groupChunkSize));
  }
  return result;
});
</script>
