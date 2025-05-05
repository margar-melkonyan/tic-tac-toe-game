<template>
  <v-container
    max-width="800px"
    class="pa-0"
  >
    <!-- Используйте v-container для ограничения ширины -->
    <v-row class="d-flex justify-center">
      <v-col
        cols="12"
        class="d-flex justify-space-between align-center"
      >
        <div
          v-if="roomInfo"
          class="d-flex align-center"
        >
          Комната: {{ roomInfo?.name.toUpperCase() }}
        </div>
        <div v-else>
          Идет загрузка комнаты...
        </div>
        <v-btn @click="exitFromRoom">
          {{ $t('rooms.exit') }}
        </v-btn>
      </v-col>
    </v-row>

    <v-divider class="mb-4 mt-8" />

    <v-row class="d-flex justify-center">
      <v-col
        cols="12"
        class="d-flex align-center"
      >
        <v-col
          v-if="props.versus"
          cols="4"
          class="d-flex justify-start"
          style="color:red; font-size: 18px;"
        >
          {{ props.versus }}
        </v-col>
        <v-col
          v-else
          cols="4"
        >
          Ждем противника!
        </v-col>
        <v-col class="d-flex justify-center" cols="4">
          <img
            src="../../assets/vs.png"
            style="height: 45px;"
          >
        </v-col>
        <v-col
          cols="4"
          class="d-flex justify-end"
          style="color:greenyellow; font-size: 18px;"
        >
          {{ authStore.user?.name }}
        </v-col>
      </v-col>
    </v-row>
    <v-divider class="my-4" />
  </v-container>
</template>

<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();
const props = defineProps({
  versus: String,
  roomInfo: Object,
  exitFromRoom: Function,
});
</script>
