<template>
  <v-card-text>
    <v-col>
      <v-text-field v-model="form.email" :error-messages="form.errors.get('email')" hide-details="auto"
                    :label="$t('enterForm.email')" variant="outlined"
      />
    </v-col>
    <v-col>
      <v-text-field v-model="form.password" :error-messages="form.errors.get('password')" hide-details="auto"
                    :label="$t('enterForm.password')" variant="outlined" type="password"
      />
    </v-col>
  </v-card-text>
  <v-card-actions>
    <v-btn @click="signIn" variant="tonal" block>
      {{ $t('enterForm.sign_in') }}
    </v-btn>
  </v-card-actions>
</template>
<script lang="ts" setup>
import {ref} from 'vue';
import {Form} from 'vform';

import API from '@/api';


const form = ref(new Form({
  email: '',
  password: ''
}));

const auth = (new API()).api.auth;
const signIn = () => {
  form.value.post(auth.urls.signIn(), {
    headers: {
      "Content-Type": "application/json",
    }
  })
    .then(response => {
      console.log(response);
    })
    .catch(({response}) => {
      if (response.status === 422) {
      } else {
      }
    });
}
</script>
