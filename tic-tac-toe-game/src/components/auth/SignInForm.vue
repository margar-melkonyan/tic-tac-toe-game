<template>
  <v-card-text>
    <v-col>
      <v-text-field
        hide-details="auto"
        :label="$t('enterForm.email')"
        variant="outlined"
      />
    </v-col>
    <v-col>
      <v-text-field
        hide-details="auto"
        :label="$t('enterForm.password')"
        variant="outlined"
        type="password"
      />
    </v-col>
  </v-card-text>
  <v-card-actions>
    <v-btn
      @click="signIn"
      variant="tonal"
      block
    >
      {{ $t('enterForm.sign_in') }}
    </v-btn>
  </v-card-actions>
</template>
<script lang="ts" setup>
import {defineEmits} from 'vue';
import axios from "axios";
import {inject} from 'vue'

const $api = inject('$api')

const emits = defineEmits([
  'close'
])

const signIn = () => {
  axios.post($api.api.auth.signIn())
}
</script>
