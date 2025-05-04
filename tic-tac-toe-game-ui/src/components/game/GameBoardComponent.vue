<template>
  <div style="display: flex; flex-direction: column; width: 100%; max-width: 600px;">
    <div
      v-for="i in props.rowsAndColumns"
      :key="i"
      style="display: flex; width: 100%;"
    >
      <div
        v-for="j in props.rowsAndColumns"
        :key="j"
        :class="`grid-index-${i}-${j}`"
        :style="props.getCellStyle()"
        @click="props.makeStep(i, j)"
      >
        <span
          class="no-select"
          :style="props.getFontStyle()"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps({
  rowsAndColumns: Number,
  makeStep: Function,
  getCellStyle: Function,
  getFontStyle: Function,
});
</script>
