<template>
  <v-col
    class="d-flex justify-space-between"
    cols="12"
  >
    <span
      v-if="props.versus"
      class="justify-start"
      style="color:red; font-size: 18px;"
    >
      {{ props.versus }}
    </span>
    <span v-else>Ждем противника!</span>
    <img
      src="../../assets/vs.png"
      style="height: 45px;"
    >
    <span
      class="justify-end"
      style="color:greenyellow; font-size: 18px;"
    >
      {{ authStore.user?.name }}
    </span>
  </v-col>
</template>

<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();
const props = defineProps({
  versus: String,
});
</script>
