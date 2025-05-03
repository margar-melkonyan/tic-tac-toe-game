<template>
  <v-container>
    <GameRoom :room-id="route.params.id" />
  </v-container>
</template>

<script setup lang="ts">

import { useRoute } from 'vue-router'
const route = useRoute()
</script>
