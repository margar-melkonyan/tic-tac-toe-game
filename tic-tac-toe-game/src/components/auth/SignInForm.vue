<template>
  <v-card-text>
    <v-col>
      <v-text-field
        v-model="form.email"
        :error-messages="form.errors.get('email')"
        hide-details="auto"
        :label="$t('enterForm.email')"
        variant="outlined"
      />
    </v-col>
    <v-col>
      <v-text-field
        v-model="form.password"
        :error-messages="form.errors.get('password')"
        hide-details="auto"
        :label="$t('enterForm.password')"
        variant="outlined"
        type="password"
      />
    </v-col>
  </v-card-text>
  <v-card-actions>
    <v-btn
      :loading="form.busy"
      variant="tonal"
      block
      @click="signIn"
    >
      {{ $t('enterForm.sign_in') }}
    </v-btn>
  </v-card-actions>
</template>
<script lang="ts" setup>
import { ref, defineEmits } from 'vue';
import { Form } from 'vform';
import 'vue3-toastify/dist/index.css';
import { useAuthStore } from '@/stores/auth';
const authStore = useAuthStore();
const form = ref(new Form({
  email: '',
  password: ''
}));
const emit = defineEmits([
  "close",
])
const signIn = () => {
  authStore.signIn(form.value)
    .then((value) => {
      if (value === 200) {
        closeFormDialog()
      }
    })
}
const closeFormDialog = () => {
  emit("close");
}
</script>
