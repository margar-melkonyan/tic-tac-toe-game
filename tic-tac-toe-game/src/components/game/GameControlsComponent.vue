<template>
  <v-row
    v-if="props.gameStarted === 0 && authStore.user?.id === roomInfo?.creator_id"
    class="d-flex justify-center"
  >
    <v-col
      cols="2"
      class="d-flex justify-space-between"
    >
      <v-btn
        color="red"
        block
        :disabled="props.rowsAndColumns === 3"
        @click="props.resizeBoard(-1)"
      >
        <v-icon>mdi-minus</v-icon>
      </v-btn>
    </v-col>
    <v-col cols="2">
      <v-btn
        color="green"
        block
        :disabled="props.rowsAndColumns === 10"
        @click="props.resizeBoard(1)"
      >
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();
const props = defineProps({
  gameStarted: Number,
  rowsAndColumns: Number,
  resizeBoard: Function,
  roomInfo: Object,
});
</script>
