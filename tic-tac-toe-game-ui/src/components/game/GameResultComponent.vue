<template>
  <v-row
    v-if="props.wonFlag !== 0"
    class="d-flex justify-center mt-6"
  >
    <v-col
      cols="12"
      class="text-center"
      style="max-width: 500px"
    >
      <v-alert
        v-if="props.wonFlag === -1"
        type="success"
        colored-border
        elevation="2"
      >
        🟢 Победил первый игрок (О)
      </v-alert>
      <v-alert
        v-if="props.wonFlag === 1"
        type="info"
        colored-border
        elevation="2"
      >
        🔵 Победил второй игрок (X)
      </v-alert>
      <v-alert
        v-if="props.wonFlag === -2"
        type="warning"
        colored-border
        elevation="2"
      >
        ⚪ Ничья
      </v-alert>
      <v-btn
        class="mt-4"
        color="primary"
        @click="props.doResetGame"
      >
        🔄 Начать заново
      </v-btn>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
const props = defineProps({
  wonFlag: Number,
  doResetGame: Function,
});
</script>
