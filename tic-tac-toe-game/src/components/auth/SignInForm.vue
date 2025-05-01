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
      variant="tonal"
      block
      @click="signIn"
    >
      {{ $t('enterForm.sign_in') }}
    </v-btn>
  </v-card-actions>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import { Form } from 'vform';
import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';
import { getCurrentInstance } from 'vue';
import type { AxiosError, AxiosResponse } from 'axios';
const { proxy } = getCurrentInstance();
const form = ref(new Form({
  email: '',
  password: ''
}));
const auth = proxy.$api.auth;
const signIn = () => {
  form.value.post(auth.urls.signIn(), {
    headers: {
      "Content-Type": "application/json",
    }
  })
    .then((data: AxiosResponse) => {
      console.log(data);
    })
    .catch((error:AxiosError ) => {
      if (error.response?.status !== 422) {
        toast.error(error.response?.data.message)
      }
    });
}
</script>
