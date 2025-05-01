<template>
  <v-card-text>
    <v-col>
      <v-text-field
        v-model="form.name"
        :error-messages="form.errors.get('name')"
        hide-details="auto"
        :label="$t('enterForm.name')"
        variant="outlined"
      />
    </v-col>
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
    <v-col>
      <v-text-field
        v-model="form.password_confirmation"
        :error-messages="form.errors.get('password_confirmation')"
        hide-details="auto"
        :label="$t('enterForm.password_confirmation')"
        variant="outlined"
      />
    </v-col>
  </v-card-text>
  <v-card-actions>
    <v-btn
      variant="tonal"
      block
      @click="signUp"
    >
      {{ $t('enterForm.sign_up') }}
    </v-btn>
  </v-card-actions>
</template>

<script lang="ts" setup>
import { useAuthStore } from '@/stores/auth';
import Form from 'vform';
import { defineEmits } from 'vue';
import { toast } from 'vue3-toastify';
const authStore = useAuthStore();
const emit = defineEmits([
  "close",
])
const form = ref(new Form({
  name: '',
  email: '',
  password: '',
  password_confirmation: null,
}));
const signUp = () => {
  authStore.signUp(form.value)
    .then((data) => {
      if (data.status === 200) {
        toast.success(data.data?.message)
        closeFormDialog()
      }
    })
}
const closeFormDialog = () => {
  emit("close");
}
</script>
