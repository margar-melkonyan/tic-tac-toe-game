<template>
  <v-card>
    <v-card-title class="mx-2 my-2">
      {{ $t('room.title', [props.roomName]) }}
    </v-card-title>
    <v-divider />
    <v-card-text>
      <v-text-field
        v-model="password"
        variant="outlined"
        density="compact"
        :append-inner-icon="isHiddePassword ? 'mdi-eye' : 'mdi-eye-off'"
        :type="isHiddePassword ? 'password' : 'text'"
        @click:append-inner="showPassword"
      >
        <template #label>
          {{ $t('room.fields.password') }}
        </template>
      </v-text-field>
    </v-card-text>
    <v-divider />
    <v-card-actions>
      <v-col class="d-flex justify-start py-0">
        <v-btn @click="exitRoom">
          {{ $t('rooms.exit') }}
        </v-btn>
      </v-col>
      <v-col
        class="d-flex justify-end py-0"
      >
        <v-btn
          :loading="loading"
          @click="enterRoom"
        >
          {{ $t('rooms.enter') }}
        </v-btn>
      </v-col>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from "vue";
const password = ref<string>("");
const loading = ref<boolean>(false);
const router = useRouter();
const props = defineProps<{
  roomName: String,
  roomId: Number,
  connectToRoom: Function,
}>()
const isHiddePassword = ref(true);
const showPassword = () => {
  isHiddePassword.value = !isHiddePassword.value
}
const enterRoom = () => {
  props.connectToRoom(props.roomId, password.value)
}
const exitRoom = () => {
  router.push({ name: "index" })
}
onMounted(() => {
  loading.value = true;
  setTimeout(() => {
    loading.value = false;
  }, 2000);
})
</script>
