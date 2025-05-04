<template>
  <v-card class="pa-4">
    <v-card-title>
      Статистика за последние  {{ scores.length }} / 50 игр
    </v-card-title>
    <v-card
      v-for="(score, key) in scores"
      :key="key"
      class="my-2 mx-4"
      :color="labelWithColor(score).color"
    >
      <v-card-text style="color: white">
        <span>
          {{ labelWithColor(score).title }}
        </span> {{ score.nickname }} {{ moment(score.created_at).utc().format('DD.MM.YYYY HH:mm:ss') }}
      </v-card-text>
    </v-card>
    <v-card-actions>
      <v-btn
        variant="tonal"
        block
        @click="close"
      >
        {{ $t('close') }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts" setup>
import { ref, getCurrentInstance, defineEmits } from 'vue';
import moment from "moment";
import axios from "axios";
const { proxy } = getCurrentInstance();
const apiScores = proxy.$api.scores;
const scores = ref([])
const emit = defineEmits([
  "closeDialog",
])
const labelWithColor = (score) => {
  if (score.is_won === 1) {
    return {title:"Выиграл у игрока", color: "green"}
  }
  if (score.is_won === 0) {
    return {title:"Проиграл игроку", color: "red"}
  }
  if (score.is_won === -1) {
    return {title:"Ничья c", color: "orange"}
  }
}
const close = () => {
  emit("closeDialog")
}
axios.get(apiScores.urls.scores())
  .then(({data}) => {
    scores.value = data.data
  })
</script>
